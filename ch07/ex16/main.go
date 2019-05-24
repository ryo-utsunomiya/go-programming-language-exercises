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
		if exprStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		env, err := parseEnv(r.FormValue("env"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		expr, err := eval.Parse(exprStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, expr.Eval(env))
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
