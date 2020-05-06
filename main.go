package main

import (
	"fmt"
	"strings"
)

var buf [1024]byte

// Program represents the series of instructions to execute
type Program struct {
	instructions string
	pointer      int
}

func (p *Program) nextInstruction(tape Tape) rune {
	if p.pointer > len(p.instructions)-1 {
		return -1
	}

	ins := p.instructions[p.pointer]
	p.pointer++

	return rune(ins)
}

// Tape represents the memory model.
type Tape struct {
	cells   map[int]int
	pointer int
}

func (t *Tape) shiftRight() {
	t.pointer++
}

func (t *Tape) shiftLeft() {
	t.pointer--
}

func (t *Tape) increment() {
	val := t.cells[t.pointer]
	t.cells[t.pointer] = val + 1
}

func (t *Tape) decrement() {
	val := t.cells[t.pointer]
	t.cells[t.pointer] = val - 1
}

func (t *Tape) getChar() string {
	return string(t.cells[t.pointer])
}

func (t *Tape) getValue() int {
	return t.cells[t.pointer]
}

// Functions to be implemented by the JS wrapper
func log(message string)
func output(message string)

//go:export echo
func echo(message string) {
	log(message)
	output(parseInput(message))
}

//go:export runBf
func runBf(input string) {
	instructions := parseInput(input)
	log(fmt.Sprintf("Executing program: '%s'", instructions))

	program := Program{pointer: 0, instructions: instructions}
	tape := Tape{pointer: 0, cells: make(map[int]int)}

	executeProgram(program, tape)
}

//go:export getBuffer
func getBuffer() *byte {
	return &buf[0]
}

func parseInput(input string) string {
	program := ""
	openBrackets := 0

	for _, char := range input {

		if strings.ContainsRune("<>+-.,[]", char) {
			if char == '[' {
				openBrackets++
			}

			if char == ']' {

			}

			program += string(char)
		}
	}

	return program
}

func executeProgram(program Program, tape Tape) {

	ins := program.nextInstruction(tape)

	for ins > 0 {

		switch ins {
		case '>':
			tape.shiftRight()
			break

		case '<':
			tape.shiftLeft()
			break

		case '+':
			tape.increment()
			break

		case '-':
			tape.decrement()
			break

		case '.':
			output(tape.getChar())
			break

		case ',':
		case ']':
		case '[':
			log("Not yet implemented")
			break

		default:
			panic("Invalid instruction")
		}

		ins = program.nextInstruction(tape)
	}
}

func main() {

}
