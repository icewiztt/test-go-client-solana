package util

import "github.com/spf13/viper"

type Config struct {
	WALLET                    string
	KYBERSWAP_POOL_PROGRAM    string
	KYBERSWAP_FACTORY_PROGRAM string
	KYBERSWAP_ROUTER_PROGRAM  string
	FACTORY_STATE_ADDRESS     string
	ROUTER_STATE_ADDRESS      string
	POOL_STATE_ADDRESS        string
	KNC_MINT                  string
	ETH_MINT                  string
	BNB_MINT                  string
	BTC_MINT                  string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
