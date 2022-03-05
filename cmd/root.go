package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{}

var cfgFile string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")

	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(transCmd)
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.AddConfigPath("configs") // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	readSection("Translate", &TranslateSetting)
}

func readSection(k string, v interface{}) error {
	err := viper.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
