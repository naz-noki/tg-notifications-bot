package configParser

import (
	"github.com/spf13/viper"
)

var (
	Cfg *config
)

func New(
	fileName, filePath, fileType string,
) error {
	Cfg = new(config)

	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(Cfg); err != nil {
		return err
	}

	return nil
}
