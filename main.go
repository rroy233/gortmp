package main

import (
	"gortmp/config"
	"gortmp/rtmp"
	"log"
)

func main() {
	var err error

	if err = InitAppConfig(); err != nil {
		return
	}

	l := ":1935"
	log.Printf("listening on :1935\n")
	err = rtmp.ListenAndServe(l)
	if err != nil {
		panic(err)
	}

	select {}

}

func InitAppConfig() (err error) {
	cfg := new(config.Config)
	err = cfg.Init("app.conf")
	if err != nil {
		return
	}

	return
}
