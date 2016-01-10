package main

import (
	"bytes"
	"os"
)

// A "method value" is what you get when you evaluate a 
// method as an expression.  The result is a function value.
//
// Below we evaluate the WriteString method of a *bytes.Buffer.

func main() {
	// Typical usage of bytes.Buffer
    b := new(bytes.Buffer)
	b.WriteString("hi!\n")
    b.WriteTo(os.Stdout)

    var buf bytes.Buffer

	// Evaluating a method from a type yields a function.
	var f func(*bytes.Buffer, string) (int, error)
    f = (*bytes.Buffer).WriteString
    f(&buf, "y u no buf.WriteString?\n")
    buf.WriteTo(os.Stdout)

	// Evaluating a method from a value creates a closure that 
	// holds that value.
	var g func(string) (int, error)
	g = buf.WriteString
	g("Hey... ")
	g("this *is* cute.\n")
	buf.WriteTo(os.Stdout)
}
