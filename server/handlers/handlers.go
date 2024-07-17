package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pnkj-kmr/looking-glass/controllers"
)

func getProtocol(c *gin.Context) {
	protocols := controllers.Protocol
	err := controllers.SetAccessData(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to update the access count",
			"error":   err.Error(),
		})
		return
	}

	var data controllers.ProtoSorter
	for id, name := range protocols {
		data = append(data, controllers.Proto{ID: id, Name: name})
	}
	// sorting the protocols
	sort.Sort(controllers.ProtoSorter(data))
	c.JSON(http.StatusOK, data)
}

func getSrcHost(c *gin.Context) {
	type host struct {
		IP   string `json:"ip"`
		Host string `json:"host"`
	}
	dataSet, err := controllers.ListIPConfig("")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}
	var data []host
	for _, ip := range dataSet {
		data = append(data, host{IP: ip.IP, Host: ip.Host})
	}
	c.JSON(http.StatusOK, data)
}

// func getResult(c *gin.Context) {
// 	var payload LGPayload
// 	out := make(chan []byte)
// 	// err := c.ShouldBind(&payload)
// 	err := c.BindJSON(&payload)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid input parameters",
// 		})
// 		return
// 	}
// 	// TODO - nee to implement the websocket here...
// 	// TODO - need to add mode check for src, proto, dst value
// 	err = controllers.Execute(payload.Src, payload.Dst, payload.Proto, out)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Unable to execute the command",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Succeed",
// 		"input":   "", // TODO - need to prepare the output
// 		"data":    "",
// 	})
// }

func getVendor(c *gin.Context) {
	// Vendor holds the display type
	type make struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	vendors := controllers.Vendor
	var data []make
	for id, name := range vendors {
		data = append(data, make{ID: id, Name: name})
	}
	c.JSON(http.StatusOK, data)
}

func getIPInfo(c *gin.Context) {
	ip := c.Param("ip")
	data, err := controllers.ListIPConfig(ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Succeed",
	})
}

func addIPInfo(c *gin.Context) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	err = controllers.AddIPConfig(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// "data":    string(payload),
		"message": "Succeed",
	})
}

func updateIPInfo(c *gin.Context) {
	ip := c.Param("ip")
	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   fmt.Sprintf("Invalid url"),
		})
		return
	}
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	err = controllers.EditIPConfig(ip, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// "data":    string(payload),
		"message": "Succeed",
	})
}

func deleteIPInfo(c *gin.Context) {
	ip := c.Param("ip")
	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   fmt.Sprintf("Invalid url"),
		})
		return
	}

	err := controllers.DeleteIPConfig(ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    ip,
		"message": "Succeed",
	})
}

func getAccessData(c *gin.Context) {
	data, err := controllers.GetAccessData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Succeed",
	})
}

func setAccessData(c *gin.Context) {
	t := c.Param("t")
	if t == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   fmt.Sprintf("Invalid url"),
		})
		return
	}

	tt, err := strconv.Atoi(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Use either 1 or 2 as valid parameter",
			"error":   err.Error(),
		})
		return
	}
	if !(tt == 1 || tt == 2) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Use either 1 or 2 as valid parameter",
			"error":   "Use Option 1 | 2",
		})
		return
	}

	err = controllers.SetAccessData(tt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input parameters",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    t,
		"message": "Succeed",
	})
}
