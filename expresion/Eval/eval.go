package eval

import (
	"fmt"
	"math"
	"strings"
)

// Var string interface
// for saving variables
type Var string

// Eval method
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Check method
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

// String method
func (v Var) String() string {
	return string(v)
}

// Number float interface
// for saving constants
type Number float64

// Eval method
func (n Number) Eval(_ Env) float64 {
	return float64(n)
}

// Check method
func (n Number) Check(vars map[Var]bool) error {
	return nil
}

// String method
func (n Number) String() string {
	return fmt.Sprintf("%f", n)
}

// Unary operation
type Unary struct {
	op  rune
	val Piper
}

// Eval method
func (u Unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.val.Eval(env)
	case '-':
		return -u.val.Eval(env)

	}
	panic(fmt.Sprintf("unsupported operation %q", u.op))
}

// Check method
func (u Unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("%q", u.op)
	}
	return nil
}

// String method
func (u Unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.val)
}

// Binary operation
type Binary struct {
	op         rune
	val1, val2 Piper
}

// Eval method
func (b Binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.val1.Eval(env) + b.val2.Eval(env)
	case '-':
		return b.val1.Eval(env) - b.val2.Eval(env)
	case '*':
		return b.val1.Eval(env) * b.val2.Eval(env)
	case '/':
		return b.val1.Eval(env) / b.val2.Eval(env)
	}
	panic(fmt.Sprintf("unsupported argument %q", b.op))
}

// Check method
func (b Binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unsupported operantor %q", b.op)
	}
	if err := b.val1.Check(vars); err != nil {
		return err
	}
	return b.val2.Check(vars)
}

// String method
func (b Binary) String() string {
	return fmt.Sprintf("%s %c %s", b.val1, b.op, b.val2)
}

// Function operations
type Function struct {
	ident string
	vals  []Piper
}

// Eval expresion
func (f Function) Eval(env Env) float64 {
	switch strings.ToLower(f.ident) {
	case "pow":
		return math.Pow(f.vals[0].Eval(env), f.vals[1].Eval(env))
	case "sqrt":
		return math.Sqrt(f.vals[0].Eval(env))
	case "sin":
		return math.Sin(f.vals[0].Eval(env))
	case "cos":
		return math.Cos(f.vals[0].Eval(env))
	case "tan":
		return math.Tan(f.vals[0].Eval(env))
	case "log":
		return math.Log10(f.vals[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported opration %q", f.ident))
}

// Check method
func (f Function) Check(vars map[Var]bool) error {
	funcArgs, ok := functions[f.ident]
	if !ok {
		return fmt.Errorf("unsupported function %s", f.ident)
	}
	if len(f.vals) != funcArgs {
		return fmt.Errorf("function %s require %d arguments, gave %d",f.ident,funcArgs,len(f.vals))
	}
	
	for _,el := range f.vals {
		if err := el.Check(vars); err != nil {
			return fmt.Errorf("Variable check error %s",err)
		}
	}

	return nil
}

func (f Function) String() string {
	if len(f.vals) > 1 {
		return fmt.Sprintf("%s(%s,%s)",f.ident,f.vals[0],f.vals[1])
	}
	return fmt.Sprintf("%s(%s)",f.ident,f.vals[0])
}

// Env interface
// for saving variables
type Env map[Var]float64

// Piper interface
type Piper interface {
	Eval(Env) float64
	Check(vars map[Var]bool) error
	String() string
}

var functions = map[string]int{
	"sin":  1,
	"cos":  1,
	"tan":  1,
	"log":  1,
	"pow":  2,
	"sqrt": 1,
}


type lexicalpanic string