package main

import (
	"fmt"
	"os"
	"os/exec"
	"project-app-cli-golang-safira/converter"
	"project-app-cli-golang-safira/suhu"
	"runtime"
)

// clear terminal screen
func clearScreen() {
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

// print main menu
func printMainMenu() {
	currentTemp := converter.GetCurrentTemperature()

	fmt.Println("--------- Suhu Converter ---------")
	fmt.Printf("Current Temperature %.2f Celcius\n", currentTemp)
	fmt.Println("1. Convert Current Temperature")
	fmt.Println("2. Custom Convert Temperature")
	fmt.Println("3. Exit")
	fmt.Println("Choose Option: ")
}

func convertCurrentTemp() {
	for {
		clearScreen()
		var celciusOption int
		var back string

		// get default or current temperature
		currentTemp := converter.GetCurrentTemperature()
		celcius := suhu.Celcius{Value: currentTemp}

		fmt.Println("--------- Celcius Converter ---------")
		fmt.Println("1. Convert to Fahrenheit")
		fmt.Println("2. Convert to Kelvin")
		fmt.Println("3. Convert to Reamur")
		fmt.Println("Choose Option: ")
		// save input from terminal
		fmt.Scan(&celciusOption)

		// add loop validation to set break or continue the infinite loop
		loopValidation := converter.CelciusConverter(celciusOption, celcius)
		if loopValidation {
			fmt.Println("Back to converter menu? (y/n)")
			fmt.Scan(&back)
			if back != "y" {
				break
			}
		} else {
			continue
		}
	}
}

// }

// print convert menu
func printCustomConvertMenu() int {
	var input int
	fmt.Println("--------- Custom Convert Menu ---------")
	fmt.Println("--------- Celcius ---------")
	fmt.Println("1. Convert from Celcius to Fahrenheit")
	fmt.Println("2. Convert from Celcius to Kelvin")
	fmt.Println("3. Convert from Celcius to Reamur")
	fmt.Println("--------- Fahrenheit ---------")
	fmt.Println("4. Convert from Fahrenheit to Celcius")
	fmt.Println("5. Convert from Fahrenheit to Kelvin")
	fmt.Println("6. Convert from Fahrenheit to Reamur")
	fmt.Println("--------- Kelvin ---------")
	fmt.Println("7. Convert from Kelvin to Celcius")
	fmt.Println("8. Convert from Kelvin to Fahrenheit")
	fmt.Println("9. Convert from Kelvin to Reamur")
	fmt.Println("--------- Reamur ---------")
	fmt.Println("10. Convert from Reamur to Celcius")
	fmt.Println("11. Convert from Reamur to Fahrenheit")
	fmt.Println("12. Convert from Reamur to Kelvin")
	fmt.Println("Choose Option: ")

	// save input from terminal
	fmt.Scan(&input)
	return input
}

// function to check if variable value has included in slice
func includes(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	// use loop label for back to main menu
mainMenu:
	for {
		var option int

		clearScreen()
		printMainMenu()
		fmt.Scan(&option)

		// convert current temperature
		if option == 1 {
			for {
				clearScreen()
				var backToMain string

				convertCurrentTemp()

				// condition to break or continue main menu infinite loop
				fmt.Println("Back to main menu? (y/n)")
				fmt.Scan(&backToMain)
				if backToMain == "y" {
					continue mainMenu
				}
				clearScreen()
				fmt.Println("Thank You!")
				break mainMenu

			}
		} else if option == 2 {
			for {
				clearScreen()
				var back, backToMain string
				var temperature float64
				var loopValidation bool
				// choosing convert menu
				option := printCustomConvertMenu()
				celciusOption := []int{1, 2, 3}
				fahrenheitOption := []int{4, 5, 6}
				kelvinOption := []int{7, 8, 9}
				reamurOption := []int{10, 11, 12}

				// save custom temperature from terminal
				fmt.Println("Temperature : ")
				fmt.Scan(&temperature)

				// use includes function to navigate converter type
				if includes(celciusOption, option) {
					celcius := suhu.Celcius{Value: temperature}
					loopValidation = converter.CelciusConverter(option, celcius)
				} else if includes(fahrenheitOption, option) {
					fahrenheit := suhu.Fahrenheit{Value: temperature}
					loopValidation = converter.FahrenheitConverter(option, fahrenheit)
				} else if includes(kelvinOption, option) {
					kelvin := suhu.Kelvin{Value: temperature}
					loopValidation = converter.KelvinConverter(option, kelvin)
				} else if includes(reamurOption, option) {
					reamur := suhu.Reamur{Value: temperature}
					loopValidation = converter.ReamurConverter(option, reamur)
				}

				// add loop validation to set break or continue current infinite loop or main menu infinite loop
				if loopValidation {
					fmt.Println("Back to converter menu? (y/n)")
					fmt.Scan(&back)
					if back != "y" {
						fmt.Println("Back to main menu? (y/n)")
						fmt.Scan(&backToMain)
						if backToMain == "y" {
							continue mainMenu
						}
						clearScreen()
						fmt.Println("Thank You!")
						break mainMenu
					}
				} else {
					fmt.Println("Back to main menu? (y/n)")
					fmt.Scan(&backToMain)
					if backToMain == "y" {
						continue mainMenu
					}
					clearScreen()
					fmt.Println("Thank You!")
					break mainMenu
				}
			}
		}
	}
}
