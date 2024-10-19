package suhu

type celciusConverter interface {
	ConvertCToF() float64
	ConvertCToR() float64
	ConvertCToK() float64
}

type fahrenheitConverter interface {
	ConvertFToR() float64
	ConvertFtoK() float64
	ConvertFtoC() float64
}

type kelvinConverter interface {
	ConvertKToR() float64
	ConvertKtoF() float64
	ConvertKToC() float64
}

type reamurConverter interface {
	ConvertRToF() float64
	ConvertRToK() float64
	ConvertRToC() float64
}

type Suhu interface {
	celciusConverter
	fahrenheitConverter
	kelvinConverter
	reamurConverter
}

// celcius
func (c Celcius) ConvertCToF() float64 {
	return 9.0/5.0*c.Value + 32.0
}

func (c Celcius) ConvertCToK() float64 {
	return c.Value + 273.0
}

func (c Celcius) ConvertCToR() float64 {
	return 4.0 / 5.0 * c.Value
}

// fahrenheit
func (f Fahrenheit) ConvertFtoK() float64 {
	return (5.0 / 9.0 * (f.Value - 32.0)) + 273.0
}

func (f Fahrenheit) ConvertFToR() float64 {
	return 4.0 / 9.0 * (f.Value - 32.0)
}

func (f Fahrenheit) ConvertFtoC() float64 {
	return 5.0 / 9.0 * (f.Value - 32.0)
}

// reamur
func (r Reamur) ConvertRToC() float64 {
	return 5.0 / 4.0 * r.Value
}

func (r Reamur) ConvertRToF() float64 {
	return 9.0/4.0*r.Value + 32.0
}

func (r Reamur) ConvertRToK() float64 {
	return (5.0 / 4.0 * r.Value) + 273.0
}

// kelvin
func (k Kelvin) ConvertKtoF() float64 {
	return 9.0/5.0*(k.Value-273.0) + 32.0
}

func (k Kelvin) ConvertKToC() float64 {
	return k.Value - 273.0
}

func (k Kelvin) ConvertKToR() float64 {
	return 4.0 / 5.0 * (k.Value - 273.0)
}
