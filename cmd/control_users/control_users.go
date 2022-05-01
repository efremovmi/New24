package main

import (
	"News24/internal/app/control_users/server"
	"News24/internal/common/helpers_function"
)

func main() {
	config := helpers_function.GetEnvParams()
	app := server.NewApp(&config)
	app.Run()
}
