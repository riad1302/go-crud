package configs

type Connection struct {
	Username string
	Password string
	Hostname string
	Dbname   string
}

func GetConnectionString() Connection {
	return Connection{
		Username: Config("DB_USER"),
		Password: Config("DB_PASSWORD"),
		Hostname: Config("DB_HOST_PORT"),
		Dbname:   Config("DB_NAME"),
	}
}
