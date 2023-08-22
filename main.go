package main

import "fmt"

func main() {
	app := InitializeApp()

	app.fiberInstance.Listen(fmt.Sprintf(":%d", app.port))
}
