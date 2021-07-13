package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Option struct {
	key   string
	value string
}

var options = []Option{
	{
		key:   "status",
		value: "Server Status",
	},
	{
		key:   "configure",
		value: "Configure Server",
	},
	{
		key:   "toggle",
		value: "Start / Stop Server",
	},
}

var reader = bufio.NewReader(os.Stdin)

func (config *AppConfig) printStatus() {
	fmt.Println("-----Status-----")
	fmt.Println("Running : ", config.isRunning)
	fmt.Println("----------------")
}

func printOptions() {
	for i, option := range options {
		fmt.Println(fmt.Sprintf("%d : %s", i, option.value))
	}
}

func (config *AppConfig) configure(){
	fmt.Println("Configuring Thing")
}

func (app *App) toggleServer(){
	fmt.Println("Toggling Server")
}

func handleOption(app *App,option int) {
	if option > -1 && option < len(options) {
		selected := options[option]
		switch selected.key{
		case "status":
			app.config.printStatus()
			break
		case "configure":
			app.config.configure()
			break
		case "toggle":
			app.toggleServer()
			break
		}
	} else {
		fmt.Println("Invalid Option Selected")
	}
}

func (app *App) runCli() {

	fmt.Println("Welcome To App")
	app.config.printStatus()

	for {
		fmt.Println()
		printOptions()

		fmt.Print("Select An Option : ")
		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		} else {
			option, err := strconv.Atoi(strings.TrimSpace(readString))
			if err != nil {
				log.Println(err)
			} else {
				handleOption(app,option)
			}
		}
	}
}
