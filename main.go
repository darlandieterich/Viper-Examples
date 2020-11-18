package main

import (
  "fmt"
  "github.com/spf13/viper"
)

func main() {
  fmt.Println("init")
  err := viper.ReadInConfig() // Find and read the config file
  if err != nil { // Handle errors reading the config file
	  fmt.Errorf("Error config file: %s \n", err)
  } else {
    debug := viper.Get("log")
    fmt.Println("Debug", debug)
  } 
  fmt.Println("end") 
}