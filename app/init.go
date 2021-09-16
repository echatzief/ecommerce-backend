package app

import "github.com/gin-gonic/gin"
import "services/laiki-eu-backend/controllers"

func InitializeApp() App{
	c := App {
		process: Process {
			name : "Laiki-EU-Backend",
			addr: "0.0.0.0",
			port: "8080",
			secret: "yVVQ6cOC0HdoISIb",
			env: "development",
		},
	}
	return c;
}

func (app App) Start(){

	// Set the gin mode
	if app.GetEnv() == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize the controllers and start the app 
	router := gin.Default()
	controllers.Setup(router)	
	router.Run(app.GetAddr()+":"+app.GetPort())
}