package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	CartSvcUrl    string `mapstructure:"CART_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigFile("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
