package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *SequencerHeader) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "SeqHash":
			err = z.SeqHash.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "SeqId":
			z.SeqId, err = dc.ReadUint64()
			if err != nil {
				return
			}
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
func (z *SequencerHeader) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "SeqHash"
	err = en.Append(0x82, 0xa7, 0x53, 0x65, 0x71, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = z.SeqHash.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "SeqId"
	err = en.Append(0xa5, 0x53, 0x65, 0x71, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.SeqId)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SequencerHeader) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "SeqHash"
	o = append(o, 0x82, 0xa7, 0x53, 0x65, 0x71, 0x48, 0x61, 0x73, 0x68)
	o, err = z.SeqHash.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "SeqId"
	o = append(o, 0xa5, 0x53, 0x65, 0x71, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.SeqId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SequencerHeader) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "SeqHash":
			bts, err = z.SeqHash.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "SeqId":
			z.SeqId, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
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
func (z *SequencerHeader) Msgsize() (s int) {
	s = 1 + 8 + z.SeqHash.Msgsize() + 6 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SequencerHeaders) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(SequencerHeaders, zb0002)
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
				(*z)[zb0001] = new(SequencerHeader)
			}
			var field []byte
			_ = field
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "SeqHash":
					err = (*z)[zb0001].SeqHash.DecodeMsg(dc)
					if err != nil {
						return
					}
				case "SeqId":
					(*z)[zb0001].SeqId, err = dc.ReadUint64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z SequencerHeaders) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0004 := range z {
		if z[zb0004] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "SeqHash"
			err = en.Append(0x82, 0xa7, 0x53, 0x65, 0x71, 0x48, 0x61, 0x73, 0x68)
			if err != nil {
				return
			}
			err = z[zb0004].SeqHash.EncodeMsg(en)
			if err != nil {
				return
			}
			// write "SeqId"
			err = en.Append(0xa5, 0x53, 0x65, 0x71, 0x49, 0x64)
			if err != nil {
				return
			}
			err = en.WriteUint64(z[zb0004].SeqId)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SequencerHeaders) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		if z[zb0004] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "SeqHash"
			o = append(o, 0x82, 0xa7, 0x53, 0x65, 0x71, 0x48, 0x61, 0x73, 0x68)
			o, err = z[zb0004].SeqHash.MarshalMsg(o)
			if err != nil {
				return
			}
			// string "SeqId"
			o = append(o, 0xa5, 0x53, 0x65, 0x71, 0x49, 0x64)
			o = msgp.AppendUint64(o, z[zb0004].SeqId)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SequencerHeaders) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(SequencerHeaders, zb0002)
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
				(*z)[zb0001] = new(SequencerHeader)
			}
			var field []byte
			_ = field
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "SeqHash":
					bts, err = (*z)[zb0001].SeqHash.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				case "SeqId":
					(*z)[zb0001].SeqId, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SequencerHeaders) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		if z[zb0004] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 8 + z[zb0004].SeqHash.Msgsize() + 6 + msgp.Uint64Size
		}
	}
	return
}