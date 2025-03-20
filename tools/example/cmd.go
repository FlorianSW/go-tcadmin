package main

import (
	"github.com/floriansw/go-tcadmin/tcadmin"
	"net/http"
	"net/http/cookiejar"
	"os"
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
	}

	c := tcadmin.NewClient(hc, "qp.qonzer.com", hllGameId, hllModId, hllFileId, tcadmin.Credentials{Username: os.Getenv("USERNAME"), Password: os.Getenv("PASSWORD")})

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
}
