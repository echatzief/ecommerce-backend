package app

type Process struct {
	name string
	addr string
	port string
	secret string
	env string
}

type Database struct {
	url string
	name string
}
type App struct {
	process Process
	database Database
}


func (c App) GetApp() Process{
	return c.process
}

func (a App) GetName() string{
	return a.process.name
}

func (a App) GetAddr() string{
	return a.process.addr
}

func (a App) GetPort() string{
	return a.process.port
}

func (a App) GetSecret() string{
	return a.process.secret
}

func (a App) GetEnv() string{
	return a.process.env
}

func (a App) GetDatabaseURL() string{
	return a.database.url
}

func (a App) GetDatabaseName() string{
	return a.database.name
}