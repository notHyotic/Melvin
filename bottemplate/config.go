package bottemplate

import (
	"log/slog"
	"os"

	"github.com/disgoorg/snowflake/v2"
	"github.com/pelletier/go-toml/v2"
)

func DefaultConfig() *Config {
	return &Config{
		Log: LogConfig{
			Level:     slog.LevelInfo, // Default log level
			Format:    "text",         // Default format
			AddSource: true,          // Default value
		},
		Bot: BotConfig{
			DevGuilds: []snowflake.ID{},
			Token:     os.Getenv("DISCORD_TOKEN"),
		},
	}
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return DefaultConfig(), nil
	}

	var cfg Config
	if err = toml.NewDecoder(file).Decode(&cfg); err != nil {
		return DefaultConfig(), nil
	}
	return &cfg, nil
}

type Config struct {
	Log LogConfig `toml:"log"`
	Bot BotConfig `toml:"bot"`
}

type BotConfig struct {
	DevGuilds []snowflake.ID `toml:"dev_guilds"`
	Token     string         `toml:"token"`
}

type LogConfig struct {
	Level     slog.Level `toml:"level"`
	Format    string     `toml:"format"`
	AddSource bool       `toml:"add_source"`
}