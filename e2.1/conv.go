package tempconv

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func CToK(c Celcius) Kelvin { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celcius { return Celcius(k - 273.15) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
