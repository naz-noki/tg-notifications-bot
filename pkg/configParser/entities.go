package configParser

type ConfigApp struct {
	Mode        string `yaml:"mode" env-required:"true" env-default:"dev"`
	LogMode     string `yaml:"logMode" env-required:"true" env-default:"debug"`
	LogFilePath string `yaml:"logFilePath" env-required:"true" env-default:"./logs/app.log"`
}

type ConfigServer struct {
	Host string `yaml:"logFilePath" env-required:"true" env-default:""`
	Port int    `yaml:"port" env-required:"true" env-default:"6565"`
}

type ConfigBot struct {
	Token  string `yaml:"token" env-required:"true"`
	Key    string `yaml:"key" env-required:"true"`
	ChatId int64  `yaml:"chatId" env-required:"true"`
}

type config struct {
	App    *ConfigApp    `yaml:"app" env-required:"true"`
	Server *ConfigServer `yaml:"server" env-required:"true"`
	Bot    *ConfigBot    `yaml:"bot" env-required:"true"`
}
