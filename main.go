/*
Copyright © 2024 rowandev
*/
package main

import (
	"fmt"
	"os"

	"github.com/danieloliveira085/autostarter"
)

func main() {
	//detect if this is the first time the app has been ran
	if _, err := os.Stat(os.Getenv("APPDATA") + "/scape.exists"); os.IsNotExist(err) {
		// add the app to startup apps
		fmt.Println("First-time install detected, adding to startup apps...")
		exec, err := os.Executable()
		if err != nil {
			fmt.Println("Error getting executable path, Tip: Make sure it's not symlinked. Error:", err)
		}
		app := autostarter.NewAutostart(
			autostarter.Shortcut{
				Name: "Scape",
				Exec: exec,
				Args: []string{},
			},
			autostarter.DefaultIcon,
		)
		app.Enable()

		os.Create(os.Getenv("APPDATA") + "/scape.exists")
	}
}
