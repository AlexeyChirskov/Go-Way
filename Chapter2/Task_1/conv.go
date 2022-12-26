package tempconv

// CToF преобразует температуру по Цельсию в температуру по Фаренгейту
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
