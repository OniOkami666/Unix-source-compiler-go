package pdp7

type TokenType int

type Word18 uint32 // Go does not have a standard 18 bit integer type
type Word13 uint16
type Word4 uint8

type PDP7 struct { // The CPU registers
	AC   Word18
	PC   Word13
	Link bool
	MAR  Word13
	IR   Word18
}

const (
	OP_MASK   Word18 = 0o740000
	INDIRECT  Word18 = 0o020000
	ADDR_MASK Word18 = 0o017777

	// Operate Instructions
	OPR_BASE Word18 = 0o740000
	CLA      Word18 = 0o000200
	CMA      Word18 = 0o001000
	HLT      Word18 = 0o000002
)

var instructions = map[string]Word18{
	"cal": 0o000000, // Call subroutine
	"dac": 0o040000, // Deposit accumulator
	"jms": 0o100000, // Jump to subroutine
	"dzm": 0o140000, // Deposit zero in memory
	"lac": 0o200000, // Load accumulator
	"xor": 0o240000, // Exclusive OR (Boolean)
	"add": 0o300000, // Add, 1's complement
	"tad": 0o340000, // Add, 2's complement
	"xct": 0o400000, // Execute
	"isz": 0o440000, // Increment and skip if zero
	"and": 0o500000, // Boolean and
	"sad": 0o540000, // Skip if AC different from memory
	"jmp": 0o600000, // Jump
}
