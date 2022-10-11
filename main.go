package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				r1 := rand.Intn(100) + 1
				r2 := rand.Intn(100) + 1
				values := map[string]int{"water": r1, "wind": r2}

				jsonValue, _ := json.Marshal(values)

				_, err := http.Post("http://127.0.0.1:8080", "application/json", bytes.NewBuffer(jsonValue))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done
	close(quit)
}
