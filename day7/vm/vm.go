package vm

import (
	"log"
)

type RAM []int
type Binop func(v1 int, v2 int) int
type InputChannel chan int
type OutputChannel chan int

type IntegerVM struct {
	Id      int
	Ram     RAM
	Inputs  InputChannel
	Outputs OutputChannel

	pc int
}

func (vm IntegerVM) currentInstruction() int {
	return vm.Ram[vm.pc] % 100
}

func (vm IntegerVM) decodeModes() (bool, bool, bool) {
	instr := vm.Ram[vm.pc]
	m1 := (instr/10000)%10 == 1
	m2 := (instr/1000)%10 == 1
	m3 := (instr/100)%10 == 1
	return m3, m2, m1
}

func (vm IntegerVM) executeBinaryOp(binop Binop) int {
	m1, m2, m3 := vm.decodeModes()

	// 2 input, 1 output
	value1 := retrieveValue(vm.Ram, vm.pc+1, m1)
	value2 := retrieveValue(vm.Ram, vm.pc+2, m2)

	value3 := binop(value1, value2)
	setValue(vm.Ram, vm.pc+3, m3, value3)
	return vm.pc + 4
}

func (vm IntegerVM) jumpIfTrue() int {
	m1, m2, _ := vm.decodeModes()
	value1 := retrieveValue(vm.Ram, vm.pc+1, m1)
	value2 := retrieveValue(vm.Ram, vm.pc+2, m2)
	if value1 > 0 {
		return value2
	} else {
		return vm.pc + 3
	}
}

func (vm IntegerVM) jumpIfFalse() int {
	m1, m2, _ := vm.decodeModes()
	value1 := retrieveValue(vm.Ram, vm.pc+1, m1)
	value2 := retrieveValue(vm.Ram, vm.pc+2, m2)
	if value1 == 0 {
		return value2
	} else {
		return vm.pc + 3
	}
}

func (vm IntegerVM) readFromInput() int {
	value := <-vm.Inputs
	m1, _, _ := vm.decodeModes()
	setValue(vm.Ram, vm.pc+1, m1, value)
	return vm.pc + 2
}

func (vm IntegerVM) writeToOutput() int {
	m1, _, _ := vm.decodeModes()
	vm.Outputs <- retrieveValue(vm.Ram, vm.pc+1, m1)
	return vm.pc + 2
}

func retrieveValue(ram RAM, locOrValue int, mode bool) int {
	if !mode {
		locOrValue = ram[locOrValue]
	}
	return ram[locOrValue]
}

func setValue(ram RAM, locOrValue int, mode bool, value int) {
	if !mode {
		locOrValue = ram[locOrValue]
	}
	ram[locOrValue] = value
}

// actual implementations

func add(v1 int, v2 int) int {
	return v1 + v2
}

func mul(v1 int, v2 int) int {
	return v1 * v2
}

func lessthan(v1 int, v2 int) int {
	if v1 < v2 {
		return 1
	} else {
		return 0
	}
}

func equals(v1 int, v2 int) int {
	if v1 == v2 {
		return 1
	} else {
		return 0
	}
}

func (vm IntegerVM) Run() int {
	for {
		switch vm.currentInstruction() {
		case 1:
			vm.pc = vm.executeBinaryOp(add)
		case 2:
			vm.pc = vm.executeBinaryOp(mul)
		case 3:
			vm.pc = vm.readFromInput()
		case 4:
			vm.pc = vm.writeToOutput()
		case 5:
			vm.pc = vm.jumpIfTrue()
		case 6:
			vm.pc = vm.jumpIfFalse()
		case 7:
			vm.pc = vm.executeBinaryOp(lessthan)
		case 8:
			vm.pc = vm.executeBinaryOp(equals)
		case 99:
			// halt
			close(vm.Outputs)
			return vm.Ram[0]
		default:
			log.Fatalf("Invalid opcode %d", vm.Ram[vm.pc])
		}
	}
	return vm.Ram[0]
}
