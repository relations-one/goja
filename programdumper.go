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

type opcode uint8

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

func (_newStash) opcode() opcode {
	return op_newStash
}

func (_noop) opcode() opcode {
	return op_noop
}

func (loadVal) opcode() opcode {
	return opLoadVal
}

func (*loadVal1) opcode() opcode {
	return opLoadVal1
}

func (_loadUndef) opcode() opcode {
	return op_loadUndef
}

func (_loadNil) opcode() opcode {
	return op_loadNil
}

func (_loadGlobalObject) opcode() opcode {
	return op_loadGlobalObject
}

func (loadStack) opcode() opcode {
	return opLoadStack
}

func (_loadCallee) opcode() opcode {
	return op_loadCallee
}

func (storeStack) opcode() opcode {
	return opStoreStack
}

func (storeStackP) opcode() opcode {
	return opStoreStackP
}

func (_toNumber) opcode() opcode {
	return op_toNumber
}

func (_add) opcode() opcode {
	return op_add
}

func (_sub) opcode() opcode {
	return op_sub
}

func (_mul) opcode() opcode {
	return op_mul
}

func (_div) opcode() opcode {
	return op_div
}

func (_mod) opcode() opcode {
	return op_mod
}

func (_neg) opcode() opcode {
	return op_neg
}

func (_plus) opcode() opcode {
	return op_plus
}

func (_inc) opcode() opcode {
	return op_inc
}

func (_dec) opcode() opcode {
	return op_dec
}

func (_and) opcode() opcode {
	return op_and
}

func (_or) opcode() opcode {
	return op_or
}

func (_xor) opcode() opcode {
	return op_xor
}

func (_bnot) opcode() opcode {
	return op_bnot
}

func (_sal) opcode() opcode {
	return op_sal
}

func (_sar) opcode() opcode {
	return op_sar
}

func (_shr) opcode() opcode {
	return op_shr
}

func (_halt) opcode() opcode {
	return op_halt
}

func (jump) opcode() opcode {
	return opJump
}

func (_setElem) opcode() opcode {
	return op_setElem
}

func (_setElemStrict) opcode() opcode {
	return op_setElemStrict
}

func (_deleteElem) opcode() opcode {
	return op_deleteElem
}

func (_deleteElemStrict) opcode() opcode {
	return op_deleteElemStrict
}

func (deleteProp) opcode() opcode {
	return opDeleteProp
}

func (deletePropStrict) opcode() opcode {
	return opDeletePropStrict
}

func (setProp) opcode() opcode {
	return opSetProp
}

func (setPropStrict) opcode() opcode {
	return opSetPropStrict
}

func (setProp1) opcode() opcode {
	return opSetProp1
}

func (_setProto) opcode() opcode {
	return op_setProto
}

func (setPropGetter) opcode() opcode {
	return opSetPropGetter
}

func (setPropSetter) opcode() opcode {
	return opSetPropSetter
}

func (getProp) opcode() opcode {
	return opGetProp
}

func (getPropCallee) opcode() opcode {
	return opGetPropCallee
}

func (_getElem) opcode() opcode {
	return op_getElem
}

func (_getElemCallee) opcode() opcode {
	return op_getElemCallee
}

func (_dup) opcode() opcode {
	return op_dup
}

func (dupN) opcode() opcode {
	return opDupN
}

func (rdupN) opcode() opcode {
	return opRdupN
}

func (_newObject) opcode() opcode {
	return op_newObject
}

func (newArray) opcode() opcode {
	return opNewArray
}

func (*newRegexp) opcode() opcode {
	return opNewRegexp
}

func (setLocal) opcode() opcode {
	return opSetLocal
}

func (setLocalP) opcode() opcode {
	return opSetLocalP
}

func (setVar) opcode() opcode {
	return opSetVar
}

func (resolveVar1) opcode() opcode {
	return opResolveVar1
}

func (deleteVar) opcode() opcode {
	return opDeleteVar
}

func (deleteGlobal) opcode() opcode {
	return opDeleteGlobal
}

func (resolveVar1Strict) opcode() opcode {
	return opResolveVar1Strict
}

func (setGlobal) opcode() opcode {
	return opSetGlobal
}

func (setVarStrict) opcode() opcode {
	return opSetVarStrict
}

func (setVar1Strict) opcode() opcode {
	return opSetVar1Strict
}

func (setGlobalStrict) opcode() opcode {
	return opSetGlobalStrict
}

func (getLocal) opcode() opcode {
	return opGetLocal
}

func (getVar) opcode() opcode {
	return opGetVar
}

func (resolveVar) opcode() opcode {
	return opResolveVar
}

func (_getValue) opcode() opcode {
	return op_getValue
}

func (_putValue) opcode() opcode {
	return op_putValue
}

func (getVar1) opcode() opcode {
	return opGetVar1
}

func (getVar1Callee) opcode() opcode {
	return opGetVar1Callee
}

func (_pop) opcode() opcode {
	return op_pop
}

func (_swap) opcode() opcode {
	return op_swap
}

func (callEval) opcode() opcode {
	return opCallEval
}

func (callEvalStrict) opcode() opcode {
	return opCallEvalStrict
}

func (_boxThis) opcode() opcode {
	return op_boxThis
}

func (call) opcode() opcode {
	return opCall
}

func (enterFunc) opcode() opcode {
	return opEnterFunc
}

func (_ret) opcode() opcode {
	return op_ret
}

func (enterFuncStashless) opcode() opcode {
	return opEnterFuncStashless
}

func (_retStashless) opcode() opcode {
	return op_retStashless
}

func (*newFunc) opcode() opcode {
	return opNewFunc
}

func (bindName) opcode() opcode {
	return opBindName
}

func (jne) opcode() opcode {
	return opJne
}

func (jeq) opcode() opcode {
	return opJeq
}

func (jeq1) opcode() opcode {
	return opJeq1
}

func (jneq1) opcode() opcode {
	return opJneq1
}

func (_not) opcode() opcode {
	return op_not
}

func (_op_lt) opcode() opcode {
	return op_op_lt
}

func (_op_lte) opcode() opcode {
	return op_op_lte
}

func (_op_gt) opcode() opcode {
	return op_op_gt
}

func (_op_gte) opcode() opcode {
	return op_op_gte
}

func (_op_eq) opcode() opcode {
	return op_op_eq
}

func (_op_neq) opcode() opcode {
	return op_op_neq
}

func (_op_strict_eq) opcode() opcode {
	return op_op_strict_eq
}

func (_op_strict_neq) opcode() opcode {
	return op_op_strict_neq
}

func (_op_instanceof) opcode() opcode {
	return op_op_instanceof
}

func (_op_in) opcode() opcode {
	return op_op_in
}

func (t try) opcode() opcode {
	return opTry
}

func (_retFinally) opcode() opcode {
	return op_retFinally
}

func (enterCatch) opcode() opcode {
	return opEnterCatch
}

func (_throw) opcode() opcode {
	return op_throw
}

func (_new) opcode() opcode {
	return op_new
}

func (_typeof) opcode() opcode {
	return op_typeof
}

func (createArgs) opcode() opcode {
	return opCreateArgs
}

func (createArgsStrict) opcode() opcode {
	return opCreateArgsStrict
}

func (_enterWith) opcode() opcode {
	return op_enterWith
}

func (_leaveWith) opcode() opcode {
	return op_leaveWith
}

func (_enumerate) opcode() opcode {
	return op_enumerate
}

func (enumNext) opcode() opcode {
	return opEnumNext
}

func (_enumGet) opcode() opcode {
	return op_enumGet
}

func (_enumPop) opcode() opcode {
	return op_enumPop
}

type progReader struct {
	io.Reader
}

func (r *progReader) readUint8() (uint8, error) {
	var val uint8
	if err := binary.Read(r, binary.BigEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
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

func (w *progWriter) writeUint8(val uint8) error {
	return binary.Write(w, binary.BigEndian, val)
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

	length := len(prg.code)
	if err := writer.writeUint32(uint32(length)); err != nil {
		return nil, err
	}

	for _, ins := range prg.code {
		if err := writer.writeUint8(uint8(ins.opcode())); err != nil {
			return nil, err
		}
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

	length, err := reader.readUint32()
	if err != nil {
		return nil, err
	}
	prg.code = make([]instruction, length)

	for i := 0; i < int(length); i++ {
		// TODO deserialize actual instruction
		_, err := reader.readUint8()
		if err != nil {
			return nil, err
		}
	}

	return prg, nil
}
