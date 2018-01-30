package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	paths := strings.Split(os.Getenv("GOPATH"), ":")
	for _, path := range paths {
		fmt.Println(path)
		viper.AddConfigPath(path + "/src/golang-learn/53_viper/config")
	}
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// Setting Defaults
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Indexes", map[string]string{"tag": "tags", "category": "categories"})

	// Setting Overrides
	LogFile := "teste.log"
	viper.Set("Verbose", true)
	viper.Set("LogFile", LogFile)

	// Registering And Using Aliases
	viper.RegisterAlias("loud", "Verbose")

	viper.Set("verbose", true) // same result as next line
	viper.Set("loud", true)    // same result as prior line

	viper.GetBool("loud")    // true
	viper.GetBool("verbose") // true

	// Getting Values
	viper.GetString("logfile") // case insensitive Setting & Getting
	if viper.GetBool("verbose") {
		fmt.Println("verbose enabled")
	}

	// Read properties from .json file
	var str interface{}
	str = viper.Get("teste_1")
	fmt.Println("teste_1:", str)

	str = viper.Sub("teste_2_1")
	fmt.Println("teste_2_1:", str)

	fmt.Println("CONFS: ", viper.AllSettings())

	// Unmarshal
	type Config struct {
		Teste_1 string
		Teste_2 struct {
			Teste_2_1 string
			Teste_2_2 int
		}
		Teste_3 struct {
			Teste_3_1 struct {
				Teste_3_1_1 string
				Teste_3_1_2 int
			}
		}
	}

	var C Config
	err = viper.Unmarshal(&C)
	if err != nil {
		log.Error("unable to decode into struct, %v", err)
	} else {
		fmt.Println("\n\nSUCESSO: ", C)
	}
}

func main() {
	// Initialization
}
