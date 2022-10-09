package main

import (
	"fmt"
	"log"

	"github.com/alresave/to-do-list/config"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(config)

}
