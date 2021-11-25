package main

import (
	"geekTime/excel/excel"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 服务1
	go startServer1(&wg)

	// 服务2
	go startServer2(&wg)

	wg.Wait()
}

func startServer1(wg *sync.WaitGroup)  {
	defer wg.Done()
	mux := http.NewServeMux()
	pathPrefix := "/excel/"
	mux.HandleFunc(pathPrefix + "create", excel.CreateExcel)
	mux.HandleFunc(pathPrefix + "download", excel.DownloadExcel)

	server := http.Server{
		Addr:              "localhost:8080",
		Handler:           mux,
	}

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Panicln("HTTP server closed.")
		} else {
			log.Panicf("HTTP server error: %v\n", err)
		}
	}
}

func startServer2(wg *sync.WaitGroup)  {
	defer wg.Done()
	var httpServer http.Server
	httpServer.Addr = "127.0.0.1:8081"
	if err := httpServer.ListenAndServe();err != nil {
		if err == http.ErrServerClosed {
			log.Println("HTTP server 1 closed.")
		} else {
			log.Printf("HTTP server 1 error: %v\n", err)
		}
	}
}