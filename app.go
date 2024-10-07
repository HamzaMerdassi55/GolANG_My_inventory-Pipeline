package main

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialise() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v",DbUser, DBPassword, DBName)
}
