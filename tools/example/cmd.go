package main

import (
	"github.com/floriansw/go-tcadmin"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

const (
	hllGameId = "1098726659"
	hllModId  = "0"
	hllFileId = "1"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	hc := http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		// Choose a relatively high timeout. Requests to restart a service are synchronous and might take a while depending on the game.
		Timeout: 60 * time.Second,
	}

	c := go_tcadmin.NewClient(hc, "qp.qonzer.com", hllGameId, hllModId, hllFileId, go_tcadmin.Credentials{Username: os.Getenv("USERNAME"), Password: os.Getenv("PASSWORD")})

	si, err := c.ServerInfo(os.Getenv("SERVICE_ID"))
	if err != nil {
		panic(err)
	}
	println(si.Name)
	println(si.Password)
	err = c.SetServerInfo(os.Getenv("SERVICE_ID"), "NewFancyName", "NewFanyPassword")
	if err != nil {
		panic(err)
	}
	pid, err := c.Restart(os.Getenv("SERVICE_ID"))
	if err != nil {
		panic(err)
	}
	println("new PID", pid)
}
