package dependencies

import (
	"github.com/jessevdk/go-flags"
)

type NovaPoshta struct {
	URL    string `long:"NP_API_URL" description:"nova poshta API URL" required:"true" default:"https://api.novaposhta.ua" env:"NP_API_URL"`
	ApiKey string `long:"NP_API_Key" description:"nova poshta API key"  default:"" env:"NP_API_KEY"`
}

type MeestExpress struct {
	ID       string `long:"ME_ID" description:"meest express ID" required:"true" default:"0xA79E003048D2B47311E26B7D4A430FFC" env:"ME_ID"`
	Login    string `long:"ME_Login" description:"meest express login" required:"true" default:"public" env:"ME_LOGIN"`
	Password string `long:"ME_Password" description:"meest express password" required:"true" default:"PUBLIC" env:"ME_PASSWORD"`
	URL      string `long:"ME_API_URL" description:"meest express API URL" required:"true" default:"https://apii.meest-group.com/T/1C_Query.php" env:"ME_API_URL"`
}

type Config struct {
	Port         string `short:"p" long:"port" description:"Port" required:"true" default:"8080" env:"PORT"`
	NovaPoshta   NovaPoshta
	MeestExpress MeestExpress
}

func initConfig() (*Config, error) {
	cfg := &Config{}
	if _, err := flags.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
