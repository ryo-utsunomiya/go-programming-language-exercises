package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ryo-utsunomiya/go-programming-language-exercises/ch07/ex16/eval"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exprStr := r.FormValue("expr")
		envStr := r.FormValue("env")

		log.Printf("expr: %s\n", exprStr)
		log.Printf("env: %s\n", envStr)

		env, err := parseEnv(envStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
			return
		}

		expr, err := eval.Parse(exprStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w,"parse expr error")
			fmt.Fprint(w, err.Error())
			return
		}

		result := expr.Eval(env)

		log.Printf("result:%f\n", result)
		fmt.Fprint(w, result)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
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
