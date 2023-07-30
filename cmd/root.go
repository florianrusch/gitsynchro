package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/florianrusch/gitsynchro/internal/config"
	"github.com/florianrusch/gitsynchro/internal/git"
	"github.com/florianrusch/gitsynchro/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFilePath string
	cfg         config.Config

	rootCmd = &cobra.Command{
		Use:   "gitsynchro",
		Short: "gitsynchro is a tool to synchronize git repos.",
		Long:  "gitsynchro is a tool to synchronize git repos.",
		Run: func(cmd *cobra.Command, args []string) {
			errors := 0

			for _, repo := range cfg.Repos {
				err := git.HandleRepo(repo)
				if err != nil {
					log.Error(err)
					errors++
				}
			}

			if errors == 0 {
				log.Infof("Successfully processed all repositories")
				os.Exit(0)
			} else {
				log.Warningf("%d error(s) occurred while processing the repositories. Please check the logs!", errors)
				os.Exit(1)
			}
		},
	}
)

// Execute executes the root command.
func Execute(_ []string) error {
	return rootCmd.Execute()
}

func init() {
	viper.Set("Verbose", true)

	rootCmd.PersistentFlags().StringVar(
		&cfgFilePath,
		"config",
		"",
		"config file (default: $HOME/gitsynchro.yaml or ./gitsynchro.yaml)",
	)
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFilePath != "" {
		// Use config file from the flag.
		if strings.Contains(cfgFilePath, "$HOME") {
			userHomeDir, err := os.UserHomeDir()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			cfgFilePath = strings.ReplaceAll(cfgFilePath, "$HOME", userHomeDir)
		}

		log.Debugf("Use config: %s", cfgFilePath)
		viper.SetConfigFile(cfgFilePath)
	} else {
		log.Debugf("Check default config locations")
		viper.SetConfigName("gitsynchro")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/")
		viper.AddConfigPath("$HOME/gitsynchro/")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("Can't parse config:", err)
		os.Exit(1)
	}
}
