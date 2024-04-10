package config

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App     App     `mapstructure:"app"`
	Adaptor Adaptor `mapstructure:"adaptor"`
}

type App struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type Adaptor struct {
	BeefAPI URL `mapstructure:"beef_api"`
}

type URL struct {
	URL string `mapstructure:"url"`
}

var config *Config
var configOnce sync.Once

func InitConfig(ctx context.Context) *Config {
	configOnce.Do(func() {
		// Change in .vscode -> launch.json -> RUN_DEBUG_PATH
		configPath, ok := os.LookupEnv("RUN_DEBUG_PATH")
		if !ok {
			slog.Info("RUN_DEBUG_PATH not found, using default config", ctx)
			configPath = "./config"
		}

		// Change in .vscode -> launch.json -> RUN_DEBUG_NAME
		configName, ok := os.LookupEnv("RUN_DEBUG_NAME")
		if !ok {
			slog.Info("RUN_DEBUG_NAME not found, using default config", ctx)
			configName = "config"
		}

		viper.AddConfigPath(configPath)
		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			slog.Warn("failed to read config file: %v (context: %v)", err, ctx)
			return
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			slog.Warn("failed to unmarshal config data: %v (context: %v)", err, ctx)
			return
		}
	})
	return config
}
