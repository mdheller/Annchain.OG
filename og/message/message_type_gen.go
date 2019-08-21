package message

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MessageCounter) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z MessageCounter) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MessageCounter) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageCounter) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MessageCounter) Msgsize() (s int) {
	s = 1
	return
}

// DecodeMsg implements msgp.Decodable
func (z *OGMessageType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint16
		zb0001, err = dc.ReadUint16()
		if err != nil {
			return
		}
		(*z) = OGMessageType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z OGMessageType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint16(uint16(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z OGMessageType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint16(o, uint16(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OGMessageType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint16
		zb0001, bts, err = msgp.ReadUint16Bytes(bts)
		if err != nil {
			return
		}
		(*z) = OGMessageType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OGMessageType) Msgsize() (s int) {
	s = msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MsgKey) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z MsgKey) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MsgKey) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MsgKey) Msgsize() (s int) {
	s = 1
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SendingType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint8
		zb0001, err = dc.ReadUint8()
		if err != nil {
			return
		}
		(*z) = SendingType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z SendingType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SendingType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SendingType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint8
		zb0001, bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			return
		}
		(*z) = SendingType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SendingType) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StatusData) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.ProtocolVersion, err = dc.ReadUint32()
	if err != nil {
		return
	}
	z.NetworkId, err = dc.ReadUint64()
	if err != nil {
		return
	}
	err = z.CurrentBlock.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.GenesisBlock.DecodeMsg(dc)
	if err != nil {
		return
	}
	z.CurrentId, err = dc.ReadUint64()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StatusData) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.ProtocolVersion)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.NetworkId)
	if err != nil {
		return
	}
	err = z.CurrentBlock.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.GenesisBlock.EncodeMsg(en)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.CurrentId)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StatusData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o = msgp.AppendUint32(o, z.ProtocolVersion)
	o = msgp.AppendUint64(o, z.NetworkId)
	o, err = z.CurrentBlock.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.GenesisBlock.MarshalMsg(o)
	if err != nil {
		return
	}
	o = msgp.AppendUint64(o, z.CurrentId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StatusData) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.ProtocolVersion, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		return
	}
	z.NetworkId, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	bts, err = z.CurrentBlock.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.GenesisBlock.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	z.CurrentId, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StatusData) Msgsize() (s int) {
	s = 1 + msgp.Uint32Size + msgp.Uint64Size + z.CurrentBlock.Msgsize() + z.GenesisBlock.Msgsize() + msgp.Uint64Size
	return
}


// DecodeMsg implements msgp.Decodable
func (z *SignedOgPartnerMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = z.BftMessage.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "BftMessage")
		return
	}
	z.TermId, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "SessionId")
		return
	}
	err = z.Signature.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	err = z.PublicKey.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "PublicKey")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SignedOgPartnerMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.BftMessage.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "BftMessage")
		return
	}
	err = en.WriteUint32(z.TermId)
	if err != nil {
		err = msgp.WrapError(err, "SessionId")
		return
	}
	err = z.Signature.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	err = z.PublicKey.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "PublicKey")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SignedOgPartnerMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.BftMessage.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "BftMessage")
		return
	}
	o = msgp.AppendUint32(o, z.TermId)
	o, err = z.Signature.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	o, err = z.PublicKey.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "PublicKey")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SignedOgPartnerMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.BftMessage.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "BftMessage")
		return
	}
	z.TermId, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "SessionId")
		return
	}
	bts, err = z.Signature.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	bts, err = z.PublicKey.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "PublicKey")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedOgPartnerMessage) Msgsize() (s int) {
	s = 1 + z.BftMessage.Msgsize() + msgp.Uint32Size + z.Signature.Msgsize() + z.PublicKey.Msgsize()
	return
}