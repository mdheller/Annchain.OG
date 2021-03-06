package tx_types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/annchain/OG/common"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BlsSigSet) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.PublicKey, err = dc.ReadBytes(z.PublicKey)
	if err != nil {
		return
	}
	z.BlsSignature, err = dc.ReadBytes(z.BlsSignature)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BlsSigSet) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PublicKey)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.BlsSignature)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BlsSigSet) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendBytes(o, z.PublicKey)
	o = msgp.AppendBytes(o, z.BlsSignature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BlsSigSet) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.PublicKey, bts, err = msgp.ReadBytesBytes(bts, z.PublicKey)
	if err != nil {
		return
	}
	z.BlsSignature, bts, err = msgp.ReadBytesBytes(bts, z.BlsSignature)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BlsSigSet) Msgsize() (s int) {
	s = 1 + msgp.BytesPrefixSize + len(z.PublicKey) + msgp.BytesPrefixSize + len(z.BlsSignature)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sequencer) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	err = z.TxBase.DecodeMsg(dc)
	if err != nil {
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			return
		}
		z.Issuer = nil
	} else {
		if z.Issuer == nil {
			z.Issuer = new(common.Address)
		}
		err = z.Issuer.DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	err = z.BlsJointSig.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.StateRoot.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Sequencer) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	if z.Issuer == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Issuer.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = z.BlsJointSig.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.StateRoot.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Sequencer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	if z.Issuer == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Issuer.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o, err = z.BlsJointSig.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.BlsJointPubKey.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.StateRoot.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sequencer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	bts, err = z.TxBase.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Issuer = nil
	} else {
		if z.Issuer == nil {
			z.Issuer = new(common.Address)
		}
		bts, err = z.Issuer.UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	bts, err = z.BlsJointSig.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.BlsJointPubKey.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.StateRoot.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Sequencer) Msgsize() (s int) {
	s = 1 + z.TxBase.Msgsize()
	if z.Issuer == nil {
		s += msgp.NilSize
	} else {
		s += z.Issuer.Msgsize()
	}
	s += z.BlsJointSig.Msgsize() + z.BlsJointPubKey.Msgsize() + z.StateRoot.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SequencerJson) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = z.TxBaseJson.DecodeMsg(dc)
	if err != nil {
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			return
		}
		z.Issuer = nil
	} else {
		if z.Issuer == nil {
			z.Issuer = new(common.Address)
		}
		err = z.Issuer.DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	err = z.BlsJointSig.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SequencerJson) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.TxBaseJson.EncodeMsg(en)
	if err != nil {
		return
	}
	if z.Issuer == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Issuer.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = z.BlsJointSig.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.BlsJointPubKey.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SequencerJson) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.TxBaseJson.MarshalMsg(o)
	if err != nil {
		return
	}
	if z.Issuer == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Issuer.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o, err = z.BlsJointSig.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.BlsJointPubKey.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SequencerJson) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.TxBaseJson.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Issuer = nil
	} else {
		if z.Issuer == nil {
			z.Issuer = new(common.Address)
		}
		bts, err = z.Issuer.UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	bts, err = z.BlsJointSig.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.BlsJointPubKey.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SequencerJson) Msgsize() (s int) {
	s = 1 + z.TxBaseJson.Msgsize()
	if z.Issuer == nil {
		s += msgp.NilSize
	} else {
		s += z.Issuer.Msgsize()
	}
	s += z.BlsJointSig.Msgsize() + z.BlsJointPubKey.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sequencers) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Sequencers, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Sequencer)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Sequencers) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Sequencers) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sequencers) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Sequencers, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Sequencer)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sequencers) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
