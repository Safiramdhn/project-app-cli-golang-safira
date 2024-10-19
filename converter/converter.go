package converter

import (
	"errors"
	"fmt"
	"log"
	"project-app-cli-golang-safira/suhu"
)

// default temperature
func GetCurrentTemperature() float64 {
	return 27.9
}

// convert calculation
func celciusConverterCalculation(suhu suhu.Suhu, convert_to string) (float64, error) {
	var result float64
	switch convert_to {
	case "fahrenheit":
		result = suhu.ConvertCToF()
	case "reamur":
		result = suhu.ConvertCToR()
	case "kelvin":
		result = suhu.ConvertCToK()
	default:
		return 0.0, errors.New("Invalid Converter")
	}
	return result, nil
}

func fahrenheitConverterCalculation(suhu suhu.Suhu, convert_to string) (float64, error) {
	var result float64
	switch convert_to {
	case "celcius":
		result = suhu.ConvertFtoC()
	case "kelvin":
		result = suhu.ConvertFtoK()
	case "reamur":
		result = suhu.ConvertFToR()
	default:
		return 0.0, errors.New("Invalid Converter")
	}
	return result, nil
}

func kelvinConverterCalculation(suhu suhu.Suhu, convert_to string) (float64, error) {
	var result float64
	switch convert_to {
	case "celcius":
		result = suhu.ConvertKToC()
	case "fahrenheit":
		result = suhu.ConvertKtoF()
	case "reamur":
		result = suhu.ConvertKToR()
	default:
		return 0.0, errors.New("Invalid Converter")
	}
	return result, nil
}

func reamurConverterCalculation(suhu suhu.Suhu, convert_to string) (float64, error) {
	var result float64
	switch convert_to {
	case "celcius":
		result = suhu.ConvertRToC()
	case "fahrenheit":
		result = suhu.ConvertRToF()
	case "kelvin":
		result = suhu.ConvertRToK()
	default:
		return 0.0, errors.New("Invalid Converter")
	}
	return result, nil
}

// converter based on option from input convert menu
func CelciusConverter(option int, celcius suhu.Celcius) bool {
	switch option {
	case 1:
		fahrenheit, err := celciusConverterCalculation(celcius, "fahrenheit")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Celcius is equal to %.2f Fahrenheit\n", celcius.Value, fahrenheit)

	case 2:
		kelvin, err := celciusConverterCalculation(celcius, "kelvin")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Celcius is equal to %.2f Kelvin\n", celcius.Value, kelvin)

	case 3:
		reamur, err := celciusConverterCalculation(celcius, "reamur")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Celcius is equal to %.2f Reamur\n", celcius.Value, reamur)

	default:
		fmt.Println("Invalid option, please try again.")
		return false
	}
	return true
}

func FahrenheitConverter(option int, fahrenheit suhu.Fahrenheit) bool {
	switch option {
	case 4:
		celcius, err := fahrenheitConverterCalculation(fahrenheit, "celcius")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Fahrenheit is equal to %.2f Celcius\n", fahrenheit.Value, celcius)

	case 5:
		kelvin, err := fahrenheitConverterCalculation(fahrenheit, "kelvin")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Fahrenheit is equal to %.2f Kelvin\n", fahrenheit.Value, kelvin)

	case 6:
		reamur, err := fahrenheitConverterCalculation(fahrenheit, "reamur")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Fahrenheit is equal to %.2f Reamur\n", fahrenheit.Value, reamur)

	default:
		fmt.Println("Invalid option, please try again.")
		return false
	}
	return true
}

func KelvinConverter(option int, kelvin suhu.Kelvin) bool {
	switch option {
	case 7:
		celcius, err := kelvinConverterCalculation(kelvin, "celcius")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Kelvin is equal to %.2f Celcius\n", kelvin.Value, celcius)

	case 8:
		fahrenheit, err := kelvinConverterCalculation(kelvin, "fahrenheit")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Kelvin is equal to %.2f fahrenheit\n", kelvin.Value, fahrenheit)

	case 9:
		reamur, err := kelvinConverterCalculation(kelvin, "reamur")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Kelvin is equal to %.2f Reamur\n", kelvin.Value, reamur)

	default:
		fmt.Println("Invalid option, please try again.")
		return false
	}
	return true
}

func ReamurConverter(option int, reamur suhu.Reamur) bool {
	switch option {
	case 10:
		celcius, err := reamurConverterCalculation(reamur, "celcius")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Reamur is equal to %.2f Celcius\n", reamur.Value, celcius)

	case 11:
		fahrenheit, err := reamurConverterCalculation(reamur, "fahrenheit")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Reamur is equal to %.2f fahrenheit\n", reamur.Value, fahrenheit)

	case 12:
		kelvin, err := reamurConverterCalculation(reamur, "kelvin")
		if err != nil {
			log.Fatalf("Conversion error: %v", err)
		}
		fmt.Printf("Temperature %.2f Reamur is equal to %.2f Kelvin\n", reamur.Value, kelvin)

	default:
		fmt.Println("Invalid option, please try again.")
		return false
	}
	return true
}
