package middleware

type dbconfig struct {
	Addr        string
	Maxopencons int
	Maxidlecons int
	Maxidletime string
}

type config struct {
	Addr     string
	DBconfig dbconfig
}

type Application struct {
	Config config
}
