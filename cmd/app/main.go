package main

import (
	"fmt"
	"net/http"
	game "soho/web/app"
	task "soho/web/app"
	user "soho/web/app"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	user.User_main()
	game.Game_main()
	task.Task_main()
}
