package main

import (
	"txc/m/server"
)

func main(){
	app := server.App{};
	app.Routes();
	app.Run();
}
