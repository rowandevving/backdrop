/*
Copyright © 2024 rowandev
*/
package main

import (
	"os"
	"runtime"

	"github.com/danieloliveira085/autostarter"
	"github.com/pelletier/go-toml/v2"
	"github.com/trevex/tray"
)

type Config struct {
	runOnStartup bool
}

func main() {
	var configDir string
	var autostart *autostarter.Autostart

	if runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA") + "/scape.toml"
		//detect if this is the first time the app has been ran
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			// add the app to startup apps
			exec, err := os.Executable()
			if err != nil {
				panic(err)
			}
			autostart = autostarter.NewAutostart(
				autostarter.Shortcut{
					Name: "Scape",
					Exec: exec,
					Args: []string{},
				},
				autostarter.DefaultIcon,
			)
			autostart.Enable()

			os.Create(configDir)
		}
	} else if runtime.GOOS == "linux" {
		configDir = os.Getenv("HOME") + "/scape.toml"
		//detect if this is the first time the app has been ran
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			// add the app to startup apps
			exec, err := os.Executable()
			if err != nil {
				os.Exit(1)
			}
			autostart = autostarter.NewAutostart(
				autostarter.Shortcut{
					Name: "Scape",
					Exec: exec,
					Args: []string{},
				},
				autostarter.DefaultIcon,
			)
			autostart.Enable()

			os.Create(configDir)
		}
	} else if runtime.GOOS == "darwin" {
		configDir = os.Getenv("HOME") + "/Library/Application Support/scape.toml"
		// add the app to startup apps
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			// add the app to startup apps
			exec, err := os.Executable()
			if err != nil {
				os.Exit(1)
			}
			autostart = autostarter.NewAutostart(
				autostarter.Shortcut{
					Name: "Scape",
					Exec: exec,
					Args: []string{},
				},
				autostarter.DefaultIcon,
			)
			autostart.Enable()

			os.Create(configDir)
		}
	}

	appTray := tray.Tray{
		Menu: []*tray.Menu{
			tray.Seperator,
			{
				Text: "Run on startup",
				Callback: func(appTray *tray.Tray, menu *tray.Menu) {
					cfgFile, err := os.ReadFile(configDir)
					if err != nil {
						panic(err)
					}
					var cfg Config
					err = toml.Unmarshal([]byte(cfgFile), &cfg)
					if err != nil {
						panic(err)
					}
					if cfg.runOnStartup {
						cfg.runOnStartup = false
						autostart.Disable()
						menu.Checked = false
					} else {
						cfg.runOnStartup = true
						autostart.Enable()
						menu.Checked = true
					}
				},
			},
		},
	}
	appTray.Run()

}
