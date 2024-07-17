package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pnkj-kmr/looking-glass/controllers"
	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
)

var dstRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9:._-]+$`)
var ipRegex *regexp.Regexp = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

type wsResponse struct {
	Completed bool   `json:"completed"`
	IsError   bool   `json:"is_error"`
	Message   string `json:"message"`
}

func wsHandler(c *gin.Context) {
	utils.L.Debug("Websocket connection request")
	//Upgrade get request to webSocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.L.Error("Unable to upgrade connection", zap.String("err", err.Error()))
		return
	}
	defer ws.Close()

	var payload LGPayload
	for {
		//read data from ws
		mt, message, err := ws.ReadMessage()
		if err != nil {
			errMsg := fmt.Sprintf(" %s", err.Error())
			utils.L.Warn("Unable to complete the read request", zap.String("err", err.Error()))
			_ = writeToWS(ws, mt, errMsg, true, true)
			break
		}
		utils.L.Debug("New request received", zap.ByteString("data", message))

		// forming ws payload
		err = json.Unmarshal(message, &payload)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to complete the request: %s", err.Error())
			utils.L.Error("Unable to unmarshal request", zap.String("err", errMsg))
			_ = writeToWS(ws, mt, errMsg, true, true)
			break
		}

		// checking dst input valid or not
		if dstRegex.FindString(payload.Dst) == "" {
			err := fmt.Errorf("invalid destination input: pattern supported [a-zA-Z0-9._-]")
			utils.L.Error("Input Dst Error", zap.String("err", err.Error()))
			_ = writeToWS(ws, mt, err.Error(), true, true)
			continue
		}

		// increasing the query count
		err = controllers.SetAccessData(2)
		if err != nil {
			utils.L.Error("Unable to audit the query count", zap.String("err", err.Error()))
		}

		// calling main execution function
		outputChannel := make(chan []byte)
		err = controllers.Execute(payload.Src, payload.Dst, payload.Proto, outputChannel)
		if err != nil {
			errMsg := fmt.Sprintf("Error: %s", err.Error())
			utils.L.Error("Error", zap.String("err", errMsg))
			// skiping error if occurs such cases
			skipError := "remote command exited without exit status or exit signal"
			if !strings.Contains(errMsg, skipError) {
				_ = writeToWS(ws, mt, errMsg, false, true)
			}
		}
		// writing into ws
		var str string
		for x := range outputChannel {
			str = string(x)
			ok, str := checkUnwantedData(str)
			if ok {
				_ = writeToWS(ws, mt, maskingIP(str, payload.Dst, payload.Proto), false, false)
			}
		}
		_ = writeToWS(ws, mt, "", true, false)
	}
	utils.L.Debug("Websocket connection disconnected")
}

func writeToWS(ws *websocket.Conn, mt int, message string, completed bool, isErr bool) (err error) {
	output, err := json.Marshal(&wsResponse{
		Completed: completed,
		Message:   message,
		IsError:   isErr,
	})
	if err != nil {
		utils.L.Error("Error while marshal response", zap.String("err", err.Error()))
		return
	}
	if completed {
		utils.L.Debug("writing into ws: COMPLETED")
	}
	err = ws.WriteMessage(mt, output)
	if err != nil {
		utils.L.Warn("Error while writing response", zap.String("err", err.Error()))
		return
	}
	return nil
}

func maskingIP(str, dstSkip string, proto int) string {
	// masking only in bgp cases
	if proto == 1 || proto == 2 || proto == 3 || proto == 4 {
		return str
	}
	if ipRegex.MatchString(str) {
		var newStr string = str
		submatchall := ipRegex.FindAllString(str, -1)
		for _, element := range submatchall {
			if dstSkip == element {
				continue
			}
			_masked := strings.Split(element, ".")[0] + ".X.X.X"
			newStr = strings.Replace(newStr, element, _masked, -1)
		}
		return newStr
	}
	return str
}

func checkUnwantedData(str string) (bool, string) {
	if strings.HasPrefix(str, "Info: ") {
		return false, ""
	}
	if strings.HasPrefix(str, "The current login time") {
		return false, ""
	}
	if strings.HasPrefix(str, "<") && strings.HasSuffix(str, ">") {
		return false, ""
	}
	// if strings.Contains(str, "exited without exit status or exit signal") {
	// 	return false, ""
	// }
	if strings.Contains(str, "Community") {
		return true, str[:strings.Index(str, "Community")+10] + " XXXX"
	}
	if strings.Contains(str, "Direct Out-interface") {
		return true, str[:strings.Index(str, "Direct Out-interface")+21] + " XXXX"
	}
	if strings.Contains(str, "Relay IP Out-Interface") {
		return true, str[:strings.Index(str, "Relay IP Out-Interface")+23] + " XXXX"
	}
	return true, str
}
