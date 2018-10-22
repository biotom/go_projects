package tempconv

// Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c * 273.15) }

//Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k / 273.15) }

//Fahrenheit to Kelvin
func FtoK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

//Kelvin to Fahrenheit
func KtoF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
