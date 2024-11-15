package main

import (
	"fmt"
	"log"
	"os"
  "github.com/joho/godotenv"
)

func main() {
  fmt.Println("Hello World!")

  godotenv.Load()

  portString := os.Getenv("PORT")
  
  if portString == "" {
    log.Fatal("PORT is not found in the environment")
  }

  fmt.Println("Port :", portString)
}
