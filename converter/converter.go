package converter

import (
	"errors"
	"project-app-cli-golang-safira/suhu"
)

func GetSuhu() float64 {
	return 27.9
}

func CelciusConverter(suhu suhu.Suhu, convert_to string) (float64, error) {
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

func FahrenheitConverter(suhu suhu.Suhu, convert_to string) (float64, error) {
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

func KelvinConverter(suhu suhu.Suhu, convert_to string) (float64, error) {
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

func ReamurConverter(suhu suhu.Suhu, convert_to string) (float64, error) {
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
