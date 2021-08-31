// Package tempconv performs Celsius, Fahrenheit, Kelvin,
// Rankine, Delisle, and Newton conversions.
package tempconv

import "fmt"

// custom type definitions
type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Réaumur float64
type Rankine float64
type Delisle float64
type Newton float64
type Rømer float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	AbsoluteZero = 0
)

func (c Celsius) String() string { return fmt.Sprintf("%g °C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g °F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g K", k) }
func (ré Réaumur) String() string { return fmt.Sprintf("%g °Ré", ré) }
func (ra Rankine) String() string { return fmt.Sprintf("%g °Ra", ra) }
func (de Delisle) String() string { return fmt.Sprintf("%g °De", de) }
func (n Newton) String() string { return fmt.Sprintf("%g °N", n) }
func (r Rømer) String() string { return fmt.Sprintf("%g °Rø", r) }
