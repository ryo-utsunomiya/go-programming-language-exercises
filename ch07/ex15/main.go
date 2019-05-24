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

	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	env, err := parseEnv(envStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(expr.Eval(env))
}

func parseEnv(envStr string) (eval.Env, error) {
	env := eval.Env{}

	for _, a := range strings.Fields(envStr) {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid var %s", a)
		}
		val, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, err
		}
		env[eval.Var(fields[0])] = val
	}

	return env, nil
}
