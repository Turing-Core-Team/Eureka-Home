package api

import (
	"fmt"

	"EurekaHome/src/api/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		fmt.Println("error starting server: ", err)
	}
}
