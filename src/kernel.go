package src

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/lobocode/wzkmo/assist"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)


var cfgFile string

// clicmd represents the base command when called without any subcommands
var cliCmd = &cobra.Command{
	Use:   "wzkmo",
	Short: "Controle de recuo e autoping",
	Long:  ``,
}

func Execute() {
	if err := cliCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func toStringMapInt(m map[string]string) map[string]int64 {
	mp := make(map[string]int64)
	for k, v := range m {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			i = 0
		}
		mp[k] = i
	}
	return mp
}

func toMilliseconds(s interface{}) time.Duration {
	v, _ := s.(string)
	d, _ := time.ParseDuration(v)
	return d
}

func toInt(s interface{}) int {
	v, _ := s.(int)
	return v
}
func toUint16(s interface{}) uint16 {

	i, _ := s.(int)
	v := uint16(i)

	return v
}

func printMaps(maps ...map[string]interface{}) {
	for _, m := range maps {
		for k, v := range m {
			fmt.Println(k, "=", v)
		}
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".wzkmo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wzkmo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings in yaml file.
	dir, _ := os.Getwd()
	defaultConfig := filepath.Join(dir, ".wzkmo.yaml")
	cliCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultConfig, "config file")
}
