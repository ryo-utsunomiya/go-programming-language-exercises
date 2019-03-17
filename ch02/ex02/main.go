package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch02/ex02/converter"
)

func main() {
	var (
		u = flag.String("u", "t", "unit")
	)
	flag.Parse()

	if flag.NFlag() > 0 {
		arg := flag.Arg(0)
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "converter: %v\n", err)
			os.Exit(1)
		}
		printUnit(f, *u)
	} else {
		var unit string
		var value string
		scanner := bufio.NewScanner(os.Stdin)
		if ok := scanner.Scan(); ok {
			unit = scanner.Text()
		}

		if ok := scanner.Scan(); ok {
			value = scanner.Text()
		}

		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "converter: %v\n", err)
			os.Exit(1)
		}
		printUnit(f, unit)
	}
}

func printUnit(f float64, unit string) {
	switch unit {
	case "t":
		c := converter.Celsius(f)
		fmt.Printf("%s = %s\n", c, converter.CToF(c))
	case "l":
		c := converter.Meter(f)
		fmt.Printf("%s = %s\n", c, converter.MToF(c))
	case "w":
		c := converter.KiloGram(f)
		fmt.Printf("%s = %s\n", c, converter.KToP(c))
	default:
		fmt.Println("unit is required")
	}
}
