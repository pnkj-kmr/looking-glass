package gateways

import (
	"bufio"
	"io"

	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
)

func outputListener(stdout, stderr io.Reader, outChan, errChan chan []byte, out chan<- []byte) {
	go func(outChan chan<- []byte) {
		defer close(outChan)
		utils.L.Debug("StdOut started")
		scanner := bufio.NewScanner(stdout)
		for {
			if tkn := scanner.Scan(); tkn {
				rcv := scanner.Bytes()
				// raw := make([]byte, len(rcv))
				// copy(raw, rcv)
				utils.L.Debug("StdOut:", zap.ByteString("data", rcv))
				outChan <- rcv
			} else {
				if scanner.Err() != nil {
					outChan <- []byte(scanner.Err().Error())
					utils.L.Debug("StdOut:", zap.String("data", scanner.Err().Error()))
				} else {
					utils.L.Debug("StdOut: io.EOF")
				}
				break
			}
		}
		utils.L.Debug("StdOut exited")
	}(outChan)

	go func(errChan chan<- []byte) {
		defer close(errChan)
		utils.L.Debug("StdErr started")
		scanner := bufio.NewScanner(stderr)
		for {
			if tkn := scanner.Scan(); tkn {
				rcv := scanner.Bytes()
				// raw := make([]byte, len(rcv))
				// copy(raw, rcv)
				utils.L.Debug("StdErr:", zap.ByteString("data", rcv))
				errChan <- rcv
			} else {
				if scanner.Err() != nil {
					errChan <- []byte(scanner.Err().Error())
					utils.L.Debug("StdErr:", zap.String("data", scanner.Err().Error()))
				} else {
					errChan <- []byte(scanner.Text())
					utils.L.Debug("StdErr: io.EOF")
				}
				break
			}
		}
		utils.L.Debug("StdErr exited")
	}(errChan)

	go func(outChan, errChan <-chan []byte, out chan<- []byte) {
		defer close(out)
		utils.L.Debug("Listener input ...")
		var outOk, errOk bool
		var out1, err1 []byte
		for {
			select {
			case out1, outOk = <-outChan:
				if outOk {
					out <- out1
				}
			case err1, errOk = <-errChan:
				if errOk {
					out <- err1
				}
			}
			if (!outOk) && (!errOk) {
				break
			}
		}
		utils.L.Debug("Listener output ...")
	}(outChan, errChan, out)
}
