package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"router"
)

type LoginController struct {
	router.Controller
}

func (this *LoginController) Excute(message router.Msg) []byte {
	msg, err := json.Marshal(message)
	CheckError(err)
	return msg
}

func init() {
	log.Println("controller init start")
	var login LoginController
	//routers := make([][2]interface{}, 0, 20)
	router.Route(func(entry router.Msg) bool {
		if entry.Meta["msgtype"] == "login" {
			return true
		}
		return false
	}, &login)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
