package bot

import (
	"log"

	"github.com/spf13/viper"
)

type DiscordConfig struct {
	Bot_Token       string   `mapstructure:"bot_token"`
	Guild_ID        string   `mapstructure:"guild_id"`
	Forbidden_Roles []string `mapstructure:"forbidden_roles"`
}

type Config struct {
	Discord DiscordConfig `mapstructure:"discord"`
}

func LoadConfig() DiscordConfig {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("couldn't read config: %s", err)
	}

	return c.Discord
}
