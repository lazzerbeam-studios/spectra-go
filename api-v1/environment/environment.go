package environment

import "github.com/spf13/viper"

type EnvironmentConfiguration struct {
	Secret              string `mapstructure:"SECRET"`
	Database            string `mapstructure:"DATABASE"`
	VALKEY              string `mapstructure:"VALKEY"`
	Google_Bucket       string `mapstructure:"GOOGLE_BUCKET"`
	Google_Project      string `mapstructure:"GOOGLE_PROJECT"`
	Google_Credentials  string `mapstructure:"GOOGLE_CREDENTIALS"`
	Google_Maps_API_Key string `mapstructure:"GOOGLE_MAPS_API_KEY"`
}

func SetEnvironment(env string) (cfg EnvironmentConfiguration, err error) {
	if env == "docker" {
		viper.SetConfigName("docker")
	} else if env == "prod" {
		viper.SetConfigName("prod")
	} else if env == "stag" {
		viper.SetConfigName("stag")
	} else {
		viper.SetConfigName("dev")
	}

	viper.AddConfigPath("./environment")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
