package router

import "fmt"

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

type Controller interface {
	Excute(message Msg) []byte
}

var routers [][2]interface{}

func Route(judge interface{}, controller Controller) {
	switch judge.(type) {
	case func(entry Msg) bool:
		{
			var arr [2]interface{}
			arr[0] = judge
			arr[1] = controller
			routers = append(routers, arr)
			fmt.Println(routers)
		}
	case map[string]interface{}:
		{
			defaultJudge := func(entry Msg) bool {
				for keyjudge, valjudge := range judge.(map[string]interface{}) {
					val, ok := entry.Meta[keyjudge]
					if !ok {
						return false
					}
					if val != valjudge {
						return false
					}
				}
				return true
			}
			var arr [2]interface{}
			arr[0] = defaultJudge
			arr[1] = controller
			routers = append(routers, arr)
			fmt.Println(routers)
		}
	default:
		fmt.Println("Something is wrong in Router")
	}
}
