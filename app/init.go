package app

import "github.com/gin-gonic/gin"
import "fmt"
import "services/laiki-eu-backend/controllers"
import (
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func InitializeApp() App{
	c := App {
		process: Process {
			name : "Laiki-EU-Backend",
			addr: "0.0.0.0",
			port: "8080",
			secret: "yVVQ6cOC0HdoISIb",
			env: "development",
		},
		database: Database {
			name: "LaikiEU",
			url: "root:rootroot@tcp(127.0.0.1:3306)/LaikiEU",
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

	// Initialize database
	conn, err := sql.Open("mysql", app.GetDatabaseURL())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
	// Initialize the database schema
	schemaInitialization, err := conn.Query(schema)
	if err != nil {
			panic(err.Error())
	}
	defer schemaInitialization.Close()

	// Add the database handler to app configurations
	fmt.Println("[LOG] Succesfully connected to the database and initialized the schema.")

	// Initialize the controllers and start the app 
	router := gin.Default()
	controllers.Setup(router, conn)
	router.Run(app.GetAddr()+":"+app.GetPort())
}