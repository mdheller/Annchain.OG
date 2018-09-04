package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (t *Sequencer) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	err = t.TxBase.DecodeMsg(dc)
	if err != nil {
		return
	}
	t.Id, err = dc.ReadUint64()
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(t.ContractHashOrder) >= int(zb0002) {
		t.ContractHashOrder = (t.ContractHashOrder)[:zb0002]
	} else {
		t.ContractHashOrder = make([]Hash, zb0002)
	}
	for za0001 := range t.ContractHashOrder {
		err = t.ContractHashOrder[za0001].DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (t *Sequencer) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = t.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteUint64(t.Id)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(t.ContractHashOrder)))
	if err != nil {
		return
	}
	for za0001 := range t.ContractHashOrder {
		err = t.ContractHashOrder[za0001].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (t *Sequencer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, t.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o, err = t.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendUint64(o, t.Id)
	o = msgp.AppendArrayHeader(o, uint32(len(t.ContractHashOrder)))
	for za0001 := range t.ContractHashOrder {
		o, err = t.ContractHashOrder[za0001].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (t *Sequencer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	bts, err = t.TxBase.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	t.Id, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(t.ContractHashOrder) >= int(zb0002) {
		t.ContractHashOrder = (t.ContractHashOrder)[:zb0002]
	} else {
		t.ContractHashOrder = make([]Hash, zb0002)
	}
	for za0001 := range t.ContractHashOrder {
		bts, err = t.ContractHashOrder[za0001].UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (t *Sequencer) Msgsize() (s int) {
	s = 1 + t.TxBase.Msgsize() + msgp.Uint64Size + msgp.ArrayHeaderSize
	for za0001 := range t.ContractHashOrder {
		s += t.ContractHashOrder[za0001].Msgsize()
	}
	return
}