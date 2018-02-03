package goja

import (
	"encoding/binary"
	"bytes"
	"io"
	"errors"
)

const prgMagic uint32 = 0xFFEEDD00

var ErrNoCacheFile = errors.New("magic number is not a goja cache file")
var ErrVersionNotMatching = errors.New("requested version of the cache file is not matching")

type opcode int

const (
	op_add              opcode = iota
	op_and
	op_bnot
	op_boxThis
	op_dec
	op_deleteElem
	op_deleteElemStrict
	op_div
	op_dup
	op_enterWith
	op_enumGet
	op_enumPop
	op_enumerate
	op_getElem
	op_getElemCallee
	op_getValue
	op_halt
	op_inc
	op_leaveWith
	op_loadCallee
	op_loadGlobalObject
	op_loadNil
	op_loadUndef
	op_mod
	op_mul
	op_neg
	op_new
	op_newObject
	op_newStash
	op_noop
	op_not
	op_op_eq
	op_op_gt
	op_op_gte
	op_op_in
	op_op_instanceof
	op_op_lt
	op_op_lte
	op_op_neq
	op_op_strict_eq
	op_op_strict_neq
	op_or
	op_plus
	op_pop
	op_putValue
	op_ret
	op_retFinally
	op_retStashless
	op_sal
	op_sar
	op_setElem
	op_setElemStrict
	op_setProto
	op_shr
	op_sub
	op_swap
	op_throw
	op_toNumber
	op_typeof
	op_xor

	opBindName
	opCall
	opCallEval
	opCallEvalStrict
	opCreateArgs
	opCreateArgsStrict
	opDeleteGlobal
	opDeleteProp
	opDeletePropStrict
	opDeleteVar
	opDupN
	opEnterCatch
	opEnterFunc
	opEnterFuncStashless
	opEnumNext
	opGetLocal
	opGetProp
	opGetPropCallee
	opGetVar
	opGetVar1
	opGetVar1Callee
	opJeq
	opJeq1
	opJne
	opJneq1
	opJump
	opLoadStack
	opLoadVal
	opLoadVal1
	opNewArray
	opNewFunc
	opNewRegexp
	opRdupN
	opResolveVar
	opResolveVar1
	opResolveVar1Strict
	opSetGlobal
	opSetGlobalStrict
	opSetLocal
	opSetLocalP
	opSetProp
	opSetProp1
	opSetPropGetter
	opSetPropSetter
	opSetPropStrict
	opSetVar
	opSetVar1Strict
	opSetVarStrict
	opStoreStack
	opStoreStackP
	opTry
)

func instruction2Opcode(ins instruction) (opcode, error) {
	switch ins.(type) {
	case _add:
		return op_add, nil
	case _and:
		return op_and, nil
	case _bnot:
		return op_bnot, nil
	case _boxThis:
		return op_boxThis, nil
	case _dec:
		return op_dec, nil
	case _deleteElem:
		return op_deleteElem, nil
	case _deleteElemStrict:
		return op_deleteElemStrict, nil
	case _div:
		return op_div, nil
	case _dup:
		return op_dup, nil
	case _enterWith:
		return op_enterWith, nil
	case _enumGet:
		return op_enumGet, nil
	case _enumPop:
		return op_enumPop, nil
	case _enumerate:
		return op_enumerate, nil
	case _getElem:
		return op_getElem, nil
	case _getElemCallee:
		return op_getElemCallee, nil
	case _getValue:
		return op_getValue, nil
	case _halt:
		return op_halt, nil
	case _inc:
		return op_inc, nil
	case _leaveWith:
		return op_leaveWith, nil
	case _loadCallee:
		return op_loadCallee, nil
	case _loadGlobalObject:
		return op_loadGlobalObject, nil
	case _loadNil:
		return op_loadNil, nil
	case _loadUndef:
		return op_loadUndef, nil
	case _mod:
		return op_mod, nil
	case _mul:
		return op_mul, nil
	case _neg:
		return op_neg, nil
	case _new:
		return op_new, nil
	case _newObject:
		return op_newObject, nil
	case _newStash:
		return op_newStash, nil
	case _noop:
		return op_noop, nil
	case _not:
		return op_not, nil
	case _op_eq:
		return op_op_eq, nil
	case _op_gt:
		return op_op_gt, nil
	case _op_gte:
		return op_op_gte, nil
	case _op_in:
		return op_op_in, nil
	case _op_instanceof:
		return op_op_instanceof, nil
	}
	panic("illegal opcode found in bytecode")
}

type progReader struct {
	io.Reader
}

func (r *progReader) readUint16() (uint16, error) {
	var val uint16
	if err := binary.Read(r, binary.BigEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func (r *progReader) readUint32() (uint32, error) {
	var val uint32
	if err := binary.Read(r, binary.BigEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func (r *progReader) readString() (string, error) {
	size, err := r.readUint32()
	if err != nil {
		return "", err
	}
	data := make([]byte, size)
	if _, err := r.Read(data); err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *progReader) readIntArray() ([]int, error) {
	size, err := r.readUint32()
	if err != nil {
		return nil, err
	}
	data := make([]int, size)
	for i := 0; i < int(size); i++ {
		val, err := r.readUint32()
		if err != nil {
			return nil, err
		}
		data[i] = int(val)
	}
	return data, nil
}

type progWriter struct {
	io.Writer
}

func (w *progWriter) writeUint16(val uint16) error {
	return binary.Write(w, binary.BigEndian, val)
}

func (w *progWriter) writeUint32(val uint32) error {
	return binary.Write(w, binary.BigEndian, val)
}

func (w *progWriter) writeString(val string) error {
	data := []byte(val)
	if err := w.writeUint32(uint32(len(data))); err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

func (w *progWriter) writeIntArray(val []int) error {
	if err := w.writeUint32(uint32(len(val))); err != nil {
		return err
	}
	for i := 0; i < len(val); i++ {
		if err := w.writeUint32(uint32(val[i])); err != nil {
			return err
		}
	}
	return nil
}

func ExportProgram(prg *Program, version uint16) ([]byte, error) {
	buffer := new(bytes.Buffer)
	writer := &progWriter{buffer}

	if err := writer.writeUint32(prgMagic); err != nil {
		return nil, err
	}
	if err := writer.writeUint16(version); err != nil {
		return nil, err
	}

	if err := writer.writeString(prg.funcName); err != nil {
		return nil, err
	}

	if err := writer.writeString(prg.src.name); err != nil {
		return nil, err
	}

	if err := writer.writeString(prg.src.src); err != nil {
		return nil, err
	}

	if err := writer.writeUint32(uint32(prg.src.lastScannedOffset)); err != nil {
		return nil, err
	}
	if err := writer.writeIntArray(prg.src.lineOffsets); err != nil {
		return nil, err
	}

	//fmt.Println(buffer.Bytes())

	return buffer.Bytes(), nil
}

func ReadProgram(r io.Reader, version uint16) (*Program, error) {
	reader := &progReader{r}

	prg := &Program{
		code:   make([]instruction, 0),
		values: make([]Value, 0),

		src:    &SrcFile{},
		srcMap: make([]srcMapItem, 0),
	}

	magic, err := reader.readUint32()
	if err != nil {
		return nil, err
	}
	if magic != prgMagic {
		return nil, ErrNoCacheFile
	}

	v, err := reader.readUint16()
	if err != nil {
		return nil, err
	}
	if v != version {
		return nil, ErrVersionNotMatching
	}

	funcName, err := reader.readString()
	if err != nil {
		return nil, err
	}
	prg.funcName = funcName

	name, err := reader.readString()
	if err != nil {
		return nil, err
	}
	prg.src.name = name

	src, err := reader.readString()
	if err != nil {
		return nil, err
	}
	prg.src.src = src

	lastScannedOffset, err := reader.readUint32()
	if err != nil {
		return nil, err
	}
	prg.src.lastScannedOffset = int(lastScannedOffset)

	lineOffsets, err := reader.readIntArray()
	if err != nil {
		return nil, err
	}
	prg.src.lineOffsets = lineOffsets

	return prg, nil
}
