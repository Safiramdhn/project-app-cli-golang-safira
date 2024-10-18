package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"project-app-cli-golang-safira/converter"
	"project-app-cli-golang-safira/suhu"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func printMenu() {
	currentTemp := converter.GetSuhu()

	fmt.Println("--------- Suhu Converter ---------")
	fmt.Printf("Current Temperature %.2f Celcius\n", currentTemp)
	fmt.Println("1. Convert Current Temperature")
	fmt.Println("2. Custom Convert Temperature")
	fmt.Println("3. Exit")
	fmt.Println("Choose Option: ")
}

func main() {
	// use loop label for back to main menu
mainMenu:
	for {
		var option int

		ClearScreen()
		printMenu()
		fmt.Scan(&option)

		if option == 1 {
			for {
				ClearScreen()
				var celciusOption int
				var back, backToMain string
				currentTemp := converter.GetSuhu()
				celcius := suhu.Celcius{Value: currentTemp}

				fmt.Println("--------- Celcius Converter ---------")
				fmt.Println("1. Convert to Fahrenheit")
				fmt.Println("2. Convert to Kelvin")
				fmt.Println("3. Convert to Reamur")
				fmt.Println("Choose Option: ")
				fmt.Scan(&celciusOption)

				switch celciusOption {
				case 1:
					fahrenheit, err := converter.CelciusConverter(celcius, "fahrenheit")
					if err != nil {
						log.Fatalf("Conversion error: %v", err)
					}
					fmt.Printf("Current Temperature %.2f Celcius is equal to %.2f Fahrenheit\n", currentTemp, fahrenheit)

				case 2:
					kelvin, err := converter.CelciusConverter(celcius, "kelvin")
					if err != nil {
						log.Fatalf("Conversion error: %v", err)
					}
					fmt.Printf("Current Temperature %.2f Celcius is equal to %.2f Kelvin\n", currentTemp, kelvin)

				case 3:
					reamur, err := converter.CelciusConverter(celcius, "reamur")
					if err != nil {
						log.Fatalf("Conversion error: %v", err)
					}
					fmt.Printf("Current Temperature %.2f Celcius is equal to %.2f Reamur\n", currentTemp, reamur)

				default:
					fmt.Println("Invalid option, please try again.")
					continue
				}

				fmt.Println("Back to converter menu? (y/n)")
				fmt.Scan(&back)
				if back != "y" {
					fmt.Println("Back to main menu? (y/n)")
					fmt.Scan(&backToMain)
					if backToMain == "y" {
						break
					}
					ClearScreen()
					fmt.Println("Thank You!")
					break mainMenu
				}

			}
		} else if option == 2 {

		}
	}
}
