package config

type Config struct {
	App struct {
		Name string
		Port int
	}
	Db struct {
		Addr     string
		User     string
		Password string
	}
}

var AppConfig *Config

func InitConfig() {

}
