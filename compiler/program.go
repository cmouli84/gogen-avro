package compiler

import (
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
)

// Build an intermediate representation of the program where
// methods, loops, switches, etc. are represented logically.
// Then concatenate everything together and replace the flow
// control with jumps to absolute offsets.

type IRProgram struct {
	main     *IRMethod
	methods  map[string]*IRMethod
	blocks   []*IRBlock
	switches []*IRSwitch
	errors   []string
}

type IRBlock struct {
	start int
	end   int
}

func (b *IRBlock) String() string {
	return fmt.Sprintf("%v - %v", b.start, b.end)
}

type IRSwitch struct {
	start int
	cases map[int]int
	end   int
}

func (b *IRSwitch) String() string {
	return fmt.Sprintf("%v - %v - %v", b.start, b.cases, b.end)
}

func (p *IRProgram) createMethod(name string) *IRMethod {
	method := NewIRMethod(name, p)
	p.methods[name] = method
	return method
}

// Concatenate all the IR instructions and assign them absolute offsets.
// An IR instruction maps to a fixed number of VM instructions,
// So we track the length of the finished output to get the real offsets.
// Main ends with a halt(0), everything else ends with a return.
func (p *IRProgram) CompileToVM() (*vm.Program, error) {
	irProgram := make([]IRInstruction, 0)
	vmLength := 0

	p.main.addLiteral(vm.Halt, 0)
	vmLength += p.main.VMLength()
	irProgram = append(irProgram, p.main.body...)

	for _, method := range p.methods {
		method.offset = vmLength
		method.addLiteral(vm.Return, vm.NoopField)
		vmLength += method.VMLength()
		irProgram = append(irProgram, method.body...)
	}

	p.findOffsets(irProgram)
	log("Found blocks: %v", p.blocks)

	vmProgram := make([]vm.Instruction, 0)
	for _, instruction := range irProgram {
		compiled, err := instruction.CompileToVM(p)
		if err != nil {
			return nil, err
		}
		vmProgram = append(vmProgram, compiled...)
	}
	return &vm.Program{
		Instructions: vmProgram,
		Errors:       make([]string, 0),
	}, nil
}

func (p *IRProgram) findOffsets(inst []IRInstruction) {
	offset := 0
	for _, instruction := range inst {
		switch v := instruction.(type) {
		case *BlockStartIRInstruction:
			log("findOffsets() block %v - start %v", v.blockId, offset)
			p.blocks[v.blockId].start = offset
		case *BlockEndIRInstruction:
			log("findOffsets() block %v - end %v", v.blockId, offset)
			p.blocks[v.blockId].end = offset
		case *SwitchStartIRInstruction:
			log("findOffsets() block %v - start %v", v.switchId, offset)
			p.switches[v.switchId].start = offset
		case *SwitchCaseIRInstruction:
			log("findOffsets() block %v - start %v", v.switchId, offset)
			p.switches[v.switchId].cases[v.value] = offset
		case *SwitchEndIRInstruction:
			log("findOffsets() block %v - end %v", v.switchId, offset)
			p.switches[v.switchId].end = offset
		}
		offset += instruction.VMLength()
	}
}