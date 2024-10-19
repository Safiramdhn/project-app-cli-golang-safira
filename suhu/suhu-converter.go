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
	return 9/5*c.Value + 32
}

func (c Celcius) ConvertCToK() float64 {
	return c.Value + 273
}

func (c Celcius) ConvertCToR() float64 {
	return 4 / 5 * c.Value
}

// fahrenheit
func (f Fahrenheit) ConvertFtoK() float64 {
	return (5 / 9 * (f.Value - 32)) + 273
}

func (f Fahrenheit) ConvertFToR() float64 {
	return 4 / 9 * (f.Value - 32)
}

func (f Fahrenheit) ConvertFtoC() float64 {
	return 5 / 9 * (f.Value - 32)
}

// reamur
func (r Reamur) ConvertRToC() float64 {
	return 5 / 4 * r.Value
}

func (r Reamur) ConvertRToF() float64 {
	return 9/4*r.Value + 32
}

func (r Reamur) ConvertRToK() float64 {
	return (5 / 4 * r.Value) + 273
}

// kelvin
func (k Kelvin) ConvertKtoF() float64 {
	return 9/5*(k.Value-273) + 32
}

func (k Kelvin) ConvertKToC() float64 {
	return k.Value - 273
}

func (k Kelvin) ConvertKToR() float64 {
	return 4 / 5 * (k.Value - 273)
}
