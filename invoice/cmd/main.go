package main

import (
	"log"
	"os"
)

func main() {
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetPrefix("Bunda")
	log.Println("Hello World!")
}
