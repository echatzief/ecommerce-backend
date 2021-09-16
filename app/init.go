package app

import "fmt"

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

func (a App) Start(){
	fmt.Println(a.GetName())
}