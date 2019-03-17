package converter

import "fmt"

type Celsius float64
type Fahrenheit float64

type Meter float64
type Feet float64

type KiloGram float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

func (k KiloGram) String() string { return fmt.Sprintf("%gKg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%g£", p) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func MToF(m Meter) Feet { return Feet(m * 3.2808) }
func FToM(f Feet) Meter { return Meter(f / 3.2808) }

func KToP(k KiloGram) Pound { return Pound(k * 2.20462) }
func PToK(p Pound) KiloGram { return KiloGram(p / 2.20462) }
