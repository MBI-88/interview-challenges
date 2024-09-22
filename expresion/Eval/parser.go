package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// TokenScanner struct
// read token from input string
type TokenScanner struct {
	token  rune
	reader scanner.Scanner
}

func (t *TokenScanner) next()               { t.token = t.reader.Scan() }
func (t *TokenScanner) stringToken() string { return t.reader.TokenText() }
func (t *TokenScanner) describe() string {
	switch t.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("%s", t.stringToken())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", t.stringToken())
	}
	return fmt.Sprintf("%c", t.token)
}

// Expression function
// Initial scanner
func Expression(t *TokenScanner) Piper { return BinaryOp(t, 1) }

// BinaryOp function
// Operations binary type
func BinaryOp(t *TokenScanner, preInit int) Piper {
	leftOp := UnaryOp(t)
	for pre := precedence(t.token); pre >= preInit; pre-- {
		for precedence(t.token) == pre {
			op := t.token
			t.next()
			rightOp := BinaryOp(t, pre+1)
			leftOp = Binary{op, leftOp, rightOp}
		}
	}
	return leftOp
}

// UnaryOp function
// Operations unary type
func UnaryOp(t *TokenScanner) Piper {
	if t.token == '+' || t.token == '-' {
		op := t.token
		t.next()
		return Unary{op, UnaryOp(t)}
	}
	return GetOp(t)
}

// GetOp function
// get operation type
func GetOp(t *TokenScanner) Piper {
	switch t.token {
	case scanner.Ident:
		id := t.reader.TokenText()
		t.next()
		if t.token != '(' {
			return Var(id)
		}
		t.next()
		var args []Piper
		for {
			args = append(args, Expression(t))
			if t.token != ',' {
				break
			}
			t.next()
		}
		if t.token != ')' {
			msg := fmt.Sprintf("Expected ) but was gave %s", t.describe())
			panic(lexicalpanic(msg))
		}

		t.next()
		return Function{id, args}

	case scanner.Int, scanner.Float:
		number, err := strconv.ParseFloat(t.stringToken(), 64)
		if err != nil {
			panic(lexicalpanic(err.Error()))
		}
		t.next()
		return Number(number)

	case '(':
		t.next()
		e := Expression(t)
		if t.token != ')' {
			msg := fmt.Sprintf("Expected ) but was gave %s", t.describe())
			panic(lexicalpanic(msg))
		}
		t.next()
		return e

	}
	msg := fmt.Sprintln("Unsupported token")
	panic(lexicalpanic(msg))
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1

	}
	return 0
}

// Parse expretion
func Parse(s string) (_ Piper, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// not error
		case lexicalpanic:
			err = fmt.Errorf("Error %s", x)
		default:
			panic(x)

		}
	}()
	token := new(TokenScanner)
	token.reader.Init(strings.NewReader(s))
	token.reader.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	token.next()
	e := Expression(token)
	if token.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %q", token.describe())
	}
	return e, nil
}
