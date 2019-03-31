package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	t := in.Text()

	var algo = flag.String("algo", "sha256", "hash algorithm name. sha256/sha384/sha512")
	flag.Parse()

	fmt.Printf("algo: %s\n", *algo)

	switch *algo {
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(t)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(t)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(t)))
	}
}
