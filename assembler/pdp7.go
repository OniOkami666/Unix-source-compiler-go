package pdp7

import (
	'encoding/binary',
	'math/bits',
	'strings'
	'bufio'
	'unicode'
)

type TokenType int 

type Word18 uint32 // Go does not have a standard 18 bit integer type
type Word13 uint16
type Word4 uint8

type PDP7 struct { // The CPU registers
	AC Word18
	PC Word13
	Link bool
	MAR Word13
	IR Word4
}

const (
	OP_MASK = 0740000
	INDIRECT = 0020000
	ADDR_MASK = 0017777
	OPR_BASE = 0740000
	CLA = 0000200
	CMA = 0001000
	HLT = 0000002
)
instructions := map[string]Word18{ 
	'cal': 000000, // Call subroutine
	'dac': 040000, // Deposit accumulator
	'jms': 100000, // Jump to subroutine
	'dzm': 140000, // Deposit zero in memory
	'lac': 200000, // Load accumulator
	'xor': 240000, // Exclusive OR (Boolean)
	'add': 300000, // Add, 1's complement
	'tad': 340000, // Add, 2's complement
	'xct': 400000, // Execute
	'isz': 440000, // Increment and skip if zero
	'and': 500000, // Boolean and
	'sad': 540000, // Skip if AC different from memory
	'jmp': 600000, // Jump
}