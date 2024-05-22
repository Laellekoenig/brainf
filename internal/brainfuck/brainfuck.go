package brainfuck

import (
	"bytes"
	"encoding/binary"
)

type Program struct {
	code []byte
	pc   int
	tape []byte
	head int
	out  *bytes.Buffer
}

func NewProgram(code []byte) *Program {
	return &Program{
		code: code,
		pc:   0,
		tape: make([]byte, 100),
		head: 0,
		out:  new(bytes.Buffer),
	}
}

func (s *Program) createLoopChild() *Program {
	end := s.getClosingBracketPos()
	child := &Program{
		code: s.code[s.pc+1 : end],
		pc:   0,
		tape: s.tape,
		head: s.head,
		out:  s.out,
	}
	s.pc = end

	return child
}

func (s *Program) copyChildData(child *Program) {
	s.tape = child.tape
	s.head = child.head
}

func (s *Program) write() {
	binary.Write(s.out, binary.LittleEndian, s.tape[s.head])
}

func (s *Program) isTrue() bool {
	return s.tape[s.head] != 0
}

func (s *Program) isFinished() bool {
	return s.pc >= len(s.code)
}

func (s *Program) getCurr() byte {
	return s.code[s.pc]
}

func (s *Program) getClosingBracketPos() int {
	end := s.pc + 1
	nestedCount := 0

	for {
		if end == len(s.code) {
			panic("invalid loop in program")
		}

		if s.code[end] == ']' {
			if nestedCount == 0 {
				return end
			}
			nestedCount -= 1
		}

		if s.code[end] == '[' {
			nestedCount += 1
		}

		end += 1
	}
}

func (s *Program) Run() []byte {
	for !s.isFinished() {
		switch s.getCurr() {
		case '>':
			s.head += 1
		case '<':
			s.head -= 1
		case '+':
			s.tape[s.head] += 1
		case '-':
			s.tape[s.head] -= 1
		case '.':
			s.write()
		case '[':
			loopChild := s.createLoopChild()
			for loopChild.isTrue() {
				loopChild.Run()
				loopChild.pc = 0
			}
			s.copyChildData(loopChild)
		}

		s.pc += 1
	}

	return s.out.Bytes()
}
