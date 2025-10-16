package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	inferURL := "http://edgenodepanel:8080/infer"
	i := 1
	for {
		// Example payload
		payload := map[string]interface{}{
			"input": fmt.Sprintf("test-%d", i),
		}
		body, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			continue
		}
		resp, err := http.Post(inferURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Printf("Request error: %v\n", err)
		} else {
			respBody, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Loop %d - Response: %s\n", i, string(respBody))
			resp.Body.Close()
		}
		i++
		time.Sleep(2 * time.Second)
	}
}