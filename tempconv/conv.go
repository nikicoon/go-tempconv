package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF (c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK (c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// CToRa converts a Celsius temperature to Rankine.
func CToRa (c Celsius) Rankine {
	return Rankine(c * 1.8 + 491.67)
}

// CToRé converts a Celsius temperature to Réaumur.
func CToRé (c Celsius) Réaumur {
	return Réaumur(c * 0.8)
}

// CToDe converts a Celsius temperature to Delisle.
func CToDe (c Celsius) Delisle {
	return Delisle((100 - c) * 1.5)
}

// CToN converts a Celsius temperature to Newton.
func CToN (c Celsius) Newton {
	return  Newton(c * 0.33)
}

// CToRø converts a Celsius temperature to Rømer.
func CToRø (c Celsius) Rømer {
	return Rømer(c * 21/40 + 7.5)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC (f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK (f Fahrenheit) Kelvin {
	return Kelvin(((f - 32) * 5 / 9) + 273.15)
}

// FToRa converts a Fahrenheit temperature to Rankine.
func FToRa (f Fahrenheit) Rankine {
	return Rankine(f + 459.67)
}

// FToRé converts a Fahrenheit temperature to Réaumur.
func FToRé (f Fahrenheit) Réaumur {
	return Réaumur((f -32) * 4/9)
}

// FToDe converts a Fahrenheit temperature to Delisle.
func FToDe (f Fahrenheit) Delisle {
	return Delisle((212 - f) * 5/6)
}

// FToN converts a Fahrenheit temperature to Newton.
func FToN (f Fahrenheit) Newton {
	return Newton((f - 32) * 11/60)
}

// FToRø converts a Fahrenheit temperature to Rømer.
func FToRø (f Fahrenheit) Rømer {
	return Rømer((f - 32) * 7/24 + 7.5)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC (k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF (k Kelvin) Fahrenheit {
	return Fahrenheit((k - 273.15) * 1.8 + 32)
}

// KToRa converts a Kelvin temperature to Rankine.
func KToRa (k Kelvin) Rankine {
	return Rankine(k * 9/5)
}

// KToRé converts a Kelvin temperature to Réaumur.
func KToRé (k Kelvin) Réaumur {
	return Réaumur((k - 273.15) * 0.8)
}

// KToDe converts a Kelvin temperature to Delisle.
func KToDe (k Kelvin) Delisle {
	return Delisle((373.15 - k) * 1.5)
}

// KToN converts a Kelvin temperature to Newton.
func KToN (k Kelvin) Newton {
	return Newton((k - 273.15) * 0.33)
}

// KToRø converts a Kelvin temperature to Rømer.
func KToRø (k Kelvin) Rømer {
	return Rømer((k - 273.15) * 21/40 + 7.5)
}

// RéToC converts a Réaumur temperature to Celsius.
func RéToC (ré Réaumur) Celsius {
	return Celsius(ré * 1.25)
}

// RéToF converts a Réaumur temperature to Fahrenheit.
func RéToF (ré Réaumur) Fahrenheit {
	return Fahrenheit(ré * 2.25 + 32)
}

// RéToRa converts a Réaumur temperature to Rankine.
func RéToRa (ré Réaumur) Rankine {
	return Rankine(ré * 2.25 + 491.67)
}

// RéToK converts a Réaumur temperature to Kelvin.
func RéToF (ré Réaumur) Kelvin {
	return Kelvin(ré * 1.25 + 273.15)
}

// RéToDe converts a Réaumur temperature to Delisle.
func RéToDe (ré Réaumur) Delisle {
	return Delisle((80 - ré) * 1.875)
}

// RéToN converts a Réaumur temperature to Newton.
func RéToN (ré Réaumur) Newton {
	return Newton(ré * 33/80)
}

// RéToRø converts a Réaumur temperature to Rømer.
func RéToRø (ré Réaumur) Rømer {
	return Rømer(ré * 21/32 + 7.5)
}

// RaToC converts a Rankine temperature to Celsius.
func RaToC (ra Rankine) Celsius {
	return Rankine(ra * 5/9 - 273.15)
}

// RaToF converts a Rankine temperature to Fahrenheit.
func RaToF (ra Rankine) Fahrenheit {
	return Fahrenheit(ra - 459.67)
}

// RaToK converts a Rankine temperature to Kelvin.
func RaToK (ra Rankine) Kelvin {
	return Kelvin(ra * 5/9)
}

// RaToRé converts a Rankine temperature to Réaumur.
func RaToRé (ra Rankine) Réaumur {
	return Réaumur(ra * 4/9 - 218.52)
}

// RaToDe converts a Rankine temperature to Delisle.
func RaToDe (ra Rankine) Delisle {
	return Delisle((671.67 - ra) * 5/6)
}

// RaToN convertsa Rankine temperature to Newton.
func RaToN (ra Rankine) Newton {
	return Newton((ra - 491.67) * 11/60)
}

// RaToRø converts a Rankine temperature to Rømer.
func RaToRø (ra Rankine) Rømer {
	return Rømer((ra - 491.67) * 7/24 + 7.5)
}

// NToC convert a Newton temperature to Celsius.
func NToC (n Newton) Celsius {
	return Celsius(n * 100/33)
}

// NToF converts a Newton temperature to Fahrenheit.
func NToF (n Newton) Fahrenheit {
	return Fahrenheit(n * 60/11 + 32)
}

// NToK converts a Newton temperature to Kelvin.
func NToK (n Newton) Kelvin {
	return Kelvin(n * 100/33 + 273.15)
}

// NToRa converts a Newton temperature to Rankine.
func NToRa (n Newton) Rankine {
	return Rankine(n * 60/11 + 491.67)
}

// NToRé converts a Newton temperature to Réaumur.
func NToRé (n Newton) Réaumur {
	return Réaumur(n * 80/33)
}

// NToDe converts a Newton temperature to Delisle.
func NToDe (n Newton) Delisle {
	return Delisle((33 - n) * 50/11)
}

// NToRø converts a Newton temperature to Rømer.
func NToRø (n Newton) Rømer {
	return Rømer(n * 35/22 + 7.5)
}
