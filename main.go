package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func checkDisruptions() {
	for true {
		disruptions := GetDisruptions(GetTubesStatus)
		fmt.Println("Fetched status")

		if len(disruptions) > 0 {
			fmt.Println("there are disruptions")
		} else {
			fmt.Println("regular service")
		}

		time.Sleep(15 * time.Second)
	}
}

func main() {
	fmt.Println("starting tubes monitor")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	checkDisruptions()
}
