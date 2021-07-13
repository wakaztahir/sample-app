package main

func main() {
	app := &App{
		config: &AppConfig{
			isRunning: false,
		},
	}

	app.runCli()
}
