package main

import (
	"fmt"
	"project-app-cli-golang-safira/converter"
	"project-app-cli-golang-safira/suhu"
	"time"
)

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
				if option == 13 {
					continue mainMenu
				} else {
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
					} else {
						fmt.Println("Invalid option, please try again")
						// add delay before continue to display error message
						time.Sleep(2 * time.Second)
						continue
					}
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
					continue
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
		} else if option == 3 {
			fmt.Println("Thank You")
			break
		}
	}
}
