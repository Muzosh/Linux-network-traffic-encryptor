/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"cyber.ee/pq/vpn/utils"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	ifce, err := utils.EnsureTun("172.25.37.0/24")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1500)
	for {
		n, err := ifce.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("packet: % x\n", buf[:n])
	}
	// cmd.Execute()
}
