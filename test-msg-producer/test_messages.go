// Simple Program to send logs at Kafka Producer server
// listening at localhost:8080 and endpoint "/event"

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	done := make(chan struct{})
	go func() {
		for i := 0; i <= 100; i++ {
			sendMsg(i)
		}
		done <- struct{}{}
	}()
	<-done
}

func sendMsg(i int) {

	str := fmt.Sprintf(`{"msg":"User access to Document", "code":%d}`, i)
	buf := strings.NewReader(str)
	resp, err := http.Post("http://localhost:8080/event", "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("error sending msg")
	}
}
