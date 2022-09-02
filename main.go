package main

import (
	"github.com/JuanPabloGarciaMonzon/ping_service/hi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	urls := strings.Split(os.Getenv("PING_URLS"), ",") //split the enviroment variable ping_urls, and set them in a string array
	log.Println(urls)
	for _, url := range urls { //travel through the array
		go pingURL(url)

	}

	stop := make(chan os.Signal, 1) //we create a slice where the length is 1 and the value is the channel os.Signal
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	v := <-stop
	log.Println("V:", v)
	log.Println("stop channel:", stop)
	log.Println("Shutting down")

}

func pingURL(url string) {
	for {
		_, err := http.Get(url)

		if err != nil {
			log.Println("There was an error pinging: ", url)
		}
		log.Println("Pinging: ", url)
	}

	time.Sleep(5 * time.Second)
	hi.hello()
}
