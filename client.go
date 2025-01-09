package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v3/pkg/client"
)

// SessionData struct
type SessionData struct {
	Data map[string]interface{}
}

var counter int64

func main() {
	fmt.Println("client start ...")

	start := time.Now()

	c := client.New(logrus.InfoLevel, 100*time.Millisecond)
	err := c.ConnectTo("127.0.0.1:3250")
	if err != nil {
		fmt.Println(" conn server error :", err)
		return
	}

	// go func(c *client.Client) {
	// 	for {
	// 		select {
	// 		case data := <-c.MsgChannel():
	// 			if data.Err {
	// 				fmt.Println("error :", string(data.Data))
	// 				break
	// 			}
	// 			fmt.Println("data --------------", string(data.Data))
	// 		}
	// 	}

	// }(c)

	go func() {
		var t = time.Now().UnixNano() / 1e6
		for {
			select {
			case <-time.After(time.Second * 5):
				now := time.Now().UnixNano() / 1e6
				v := atomic.SwapInt64(&counter, 0)
				fmt.Println("count: ", float64(v)/float64((now-t)/1000), "/s")
				t = now
			}
		}
	}()

	message, _ := json.Marshal(SessionData{Data: map[string]interface{}{"test": "test"}})

	for {
		_, err = c.SendRequest("gate.connector.setsessiondata", message) // gate.connector.getsessiondata、game.room.getsessiondata
		if err != nil {
			fmt.Println("send request error")
			return
		}
		// time.Sleep(time.Millisecond * 100)

		select {
		case data := <-c.MsgChannel():
			if data.Err {
				fmt.Println("error :", string(data.Data))
				break
			}
			// fmt.Println("data --------------", string(data.Data))
		}

		atomic.AddInt64(&counter, 1)
	}

	fmt.Println("duration:", time.Since(start))
}
