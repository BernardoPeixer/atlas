package main

import (
	"fmt"
	"log"
)

func main() {
	log.Printf("Call to start server")

	cfgPath, action, terminal := loadFlags()

	cfg := readCFGFile(cfgPath)

	if !terminal {
		file, err := configureOutput(cfg.LogDir)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	service, err := newService(*cfg)
	if err != nil {
		fmt.Printf("Failed to install: %s\n", err)
		return
	}

	if action != "" {
		switch action {
		case "run":
			log.Println("Running service")
			err := service.Run()
			if err != nil {
				fmt.Printf("Failed to run: %s\n", err)
			}
		case "uninstall":
			log.Println("Uninstalling service")
			err := service.Uninstall()
			if err != nil {
				fmt.Printf("Failed to run: %s\n", err)
			}
		case "install":
			log.Println("Installing service")
			err = service.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
			}
		case "stop":
			log.Println("Stopping service")
			err = service.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
			}
		}
	}

	err = service.Run()
	if err != nil {
		log.Printf("Error in: the program: %v", err)
	}
}
