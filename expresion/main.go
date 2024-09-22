package main

import (
	"fmt"
	"log"
	"flag"
	"eval/Eval"
)



func main() {
	x := flag.Float64("n",0,"data")
	flag.Parse()
	args := flag.Arg(0)
	expre, err := eval.Parse(args)
	dict := make(map[eval.Var]float64)
	dict[eval.Var("x")] = *x
	if err != nil {
		log.Fatalf("Error in Parse -> %s",err)
	}
	fmt.Printf("Expression %s\n result %v\n",expre,expre.Eval(eval.Env(dict)))
}