# Calc
Calc is a prototype of a calculator with basic algebraic simplification written in Go.

# Screenshot
![Calc Screenshot](/screenshot.png)

# Features
* Basic algebraic simplification (e.g. x+x = 2*x)
* Base conversion (binary, octal and hexadecimal to decimal)
* Web interface for input and output (based on bootstrap)

# Usage
* go get github.com/sj14/calc
* cd $GOPATH/src/github.com/sj14/calc
* go run main.go
* open a web browser and enter http://localhost:9000/

# Basic architecture
1. Decide (regex) the input type (base conversion or mathematical term) (relay)

  Base Conversion
2. Do base conversion (relay)
3. Show result

  Algebraic simplification
2. Convert the input term into an Abstract Syntax Tree (shuntingyard)
3. Evaluate the AST (eval)
4. Show the result

# Problems / Limitations / TODO
* The traversing order of the AST is from right to left instead left to right
* No unary operations (e.g. -5 will result into 0-5)
* Usage of float64 instead of math/big
* Make usage of steps.go
* ...
