package main

import (
  "fmt"
  "github.com/spf13/viper"
)

func main() {  
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig()
  if err != nil {
	  fmt.Println("Error config file: %s \n", err)
  } else {
    debug := viper.GetBool("log.debug")
    fmt.Println("Debug:", debug)
    
    var config map[string]interface{}
    err := viper.Unmarshal(&config)
    if err != nil {
      fmt.Println("Error Unmarshal toml")
    } else {
      fmt.Println("Interface", config["params"])
    }    
  }   
}