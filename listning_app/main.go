package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	tcpProtocol    = "tcp4"
	keySize        = 1024
	readWriterSize = keySize / 8
)

type Config struct {
	Miner  map[string]map[string]*Miner `json:"miner"`
	Listen string                       `json:"listen"`
}

type remoteConn struct {
	c *net.TCPConn
}

type MinerStatus struct {
	Version string `json:"version"`
	Time    int64  `json:"time"`

	EHashrate      int64   `json:"e_hashrate"`
	EValidShares   int64   `json:"e_valid_shares"`
	EInvalidShares int64   `json:"e_invalid_shares"`
	EGpuHashrate   []int64 `json:"e_gpu_hashrate"`

	DHashrate      int64   `json:"d_hashrate"`
	DValidShares   int64   `json:"d_valid_shares"`
	DInvalidShares int64   `json:"d_invalid_shares"`
	DGpuHashrate   []int64 `json:"d_gpu_hashrate"`

	Pools string    `json:"pools"`
	Gpus  [][]int64 `json:"gpus"`
}

type Miner struct {
	Name   string       `json:"name"`
	Addr   string       `json:"addr"`
	Port   int64        `json:"port"`
	Status *MinerStatus `json:"status"`
	Online bool         `json:"online"`
}

var listenAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8989}

//var miner *Miner

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
}

func (cfg *Config) MinersHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(cfg)
	if err != nil {
		log.Println("Error serializing /miner: ", err)
	}
}

func getMassage(c net.Conn) Config {

	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg := Config{}

	json.Unmarshal([]byte(line), &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(miner.Status)
	return cfg

}

func listen(l net.Listener, ln net.Listener) {

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfg := getMassage(conn)
		r := mux.NewRouter()

		r.HandleFunc("/miner", cfg.MinersHandler)

		go func() { // слушает порт 8089 с отправкой состояния майнера

			if err := http.Serve(l, r); err != nil { // отправка состояния майнера
				log.Panic(err)
			}

		}()
		a := cfg.Miner
		fmt.Println(a)
	}

}

func main() {
	stop := make(chan bool)
	cfg := Config{}

	intv := time.Duration(time.Second * 10)
	timer := time.NewTimer(intv)
	go func() {
		ln, err := net.Listen("tcp", ":8989")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		l, err := net.Listen("tcp", ":8099") // слушает порт 8089
		if err != nil {
			log.Panic(err)
		}
		for {
			select {
			case <-timer.C:

				listen(l, ln)
				fmt.Println(cfg.Miner)
				//log.Println(worker.Status)

				timer.Reset(intv)
			}
		}
	}()
	stop <- true
}
