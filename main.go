package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	uuid := uuid.New()
	fmt.Println("UUID: ", uuid)

	color.Cyan("Fill this world with a color")
	color.Yellow(uuid.String())

	company := os.Getenv("COMPANY_NAME")
	fmt.Println(company)

	color.Blue(company)

}
