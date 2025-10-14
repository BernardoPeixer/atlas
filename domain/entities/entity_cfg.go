package entities

type Config struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	LogDir       string `toml:"log_dir"`
	MySQL        MySQL  `toml:"MySQL"`
	JwtSecretKey string `toml:"secret_key"`
}

type MySQL struct {
	DBHost     string `toml:"db_host"`
	DBPort     int    `toml:"db_port"`
	DBName     string `toml:"db_name"`
	DBUser     string `toml:"db_user"`
	DBPassword string `toml:"db_password"`
}
