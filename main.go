package main

import (
	"fmt"
	"strings"
)

var buf [1024]byte

// Program represents the series of instructions to execute
type Program struct {
	instructions string      // The sequence of instructions to execute
	pointer      int         // Index to the current instruction
	brackets     map[int]int // Map to keep track of control points
}

func (p *Program) nextInstruction(tape Tape) rune {

	// Stop when we reach the end.
	if p.pointer > len(p.instructions)-1 {
		return -1
	}

	ins := p.instructions[p.pointer]
	val := tape.getValue()

	if ins == '[' {

		// Jump past the matching ']' if the value in the current cell is 0
		if val == 0 {
			p.pointer = p.brackets[p.pointer] + 1
			return p.nextInstruction(tape)
		}

		// Otherwise return the next instruction
		p.pointer++
		return p.nextInstruction(tape)
	}

	if ins == ']' {

		// Jump back to the matching '[' if the value in the current cell is NOT 0
		if val != 0 {
			p.pointer = p.brackets[p.pointer]
			return p.nextInstruction(tape)
		}

		// Otherwise return the next instruction
		p.pointer++
		return p.nextInstruction(tape)
	}

	// Advance the instruction pointer for next time.
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

//go:export runBf
func runBf(input string) {
	program, err := parseInput(input)

	if err != nil {
		log(fmt.Sprintf("Syntax Error: %s", err))
		return
	}

	log(fmt.Sprintf("Executing program: '%s'", program.instructions))

	tape := Tape{pointer: 0, cells: make(map[int]int)}
	executeProgram(program, tape)
}

//go:export getBuffer
func getBuffer() *byte {
	return &buf[0]
}

func parseInput(input string) (Program, error) {
	brackets := []int{}
	bracketMap := make(map[int]int)
	instructions := ""

	for idx, char := range input {
		var start int

		if strings.ContainsRune("<>+-.,[]", char) {
			if char == '[' {
				brackets = append(brackets, idx)
			}

			if char == ']' {

				if len(brackets) == 0 {
					return Program{}, fmt.Errorf("unmatched brackets")
				}

				// Pair the most recently pushed open bracket with the current index
				start, brackets = brackets[len(brackets)-1], brackets[:len(brackets)-1]
				bracketMap[start] = idx
				bracketMap[idx] = start
			}

			instructions += string(char)
		}
	}

	if len(brackets) != 0 {
		return Program{}, fmt.Errorf("unmatched brackets")
	}

	program := Program{instructions: instructions, pointer: 0, brackets: bracketMap}
	return program, nil
}

func executeProgram(program Program, tape Tape) {

	ins := program.nextInstruction(tape)
	log("Ins: '" + string(ins) + "'")

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
			log("Not yet implemented")
			break

		default:
			panic("Invalid instruction: " + string(ins))
		}

		ins = program.nextInstruction(tape)
	}
}

func main() {

}
