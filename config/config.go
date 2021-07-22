package config

type Config struct {

	//Configuration Variables
	Mode        RunMode
	Certificate string
	Key         string
	UseHTTPS    bool

	//Different Configurations
	Server *ServerConfig
	Db     *DatabaseConfig
	Smtp   *SMTPConfig
}

type RunMode string

const (
	DevelopmentMode RunMode = "development"
	ProductionMode  RunMode = "production"
)

type ServerConfig struct {
	IsRunning       bool
	Port            int
	AllowCORS       bool
	CORSAllowedFor  []string
	VerifyRecaptcha bool
	RecaptchaSecret string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

type SMTPConfig struct {
}
