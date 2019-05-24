package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch07/ex15/eval"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("expr:")
	in.Scan()
	exprStr := in.Text()
	fmt.Print("vars:")
	in.Scan()
	envStr := in.Text()

	if in.Err() != nil {
		fmt.Fprintln(os.Stderr, in.Err())
		os.Exit(1)
	}

	env := eval.Env{}
	for _, a := range strings.Fields(envStr) {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "invalid var %s\n", a)
			os.Exit(1)
		}
		val, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		env[eval.Var(fields[0])] = val
	}

	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(expr.Eval(env))
}
