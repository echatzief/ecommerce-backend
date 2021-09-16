package main

import "services/laiki-eu-backend/app"

func main(){

	// Initialize the application with the necessary configs
	app := app.InitializeApp()

	// Start the application
	app.Start()
}
