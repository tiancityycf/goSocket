package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"route"
)

type LoginController struct {
}

func (this *LoginController) Excute(message Msg) []byte {
	mirrormsg, err := json.Marshal(message)
	CheckError(err)
	return mirrormsg
}

func init() {
	var login LoginController
	routers = make([][2]interface{}, 0, 20)
	route.Route(func(entry Msg) bool {
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
