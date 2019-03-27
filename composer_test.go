package composer

import (
	"bytes"
	"encoding/binary"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestWrite(c *C) {
	cmp := New()
	c.Assert(cmp.Byte(1).Uint16be(1).Uint16le(1).Make(),
		DeepEquals,
		[]byte{1, 0, 1, 1, 0})

	cmp = New()
	c.Assert(
		cmp.Uint16be(1).Uint16be(65534).Uint32be(1).Uint64be(1).Make(),
		DeepEquals,
		[]byte{0, 1, 255, 254, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1})

	cmp = New()
	c.Assert(
		cmp.Uint(binary.LittleEndian, 9, 3).Make(),
		DeepEquals,
		[]byte{9, 0, 0})

	cmp = New()
	c.Assert(
		cmp.Uint(binary.BigEndian, 9, 3).Make(),
		DeepEquals,
		[]byte{0, 0, 9})

	cmp = New()
	c.Assert(
		cmp.Int(binary.LittleEndian, -9, 3).Make(),
		DeepEquals,
		[]byte{247, 255, 255})

	cmp = New()
	c.Assert(
		cmp.Int(binary.BigEndian, -9, 3).Make(),
		DeepEquals,
		[]byte{255, 255, 247})

	cmp = New()
	c.Assert(
		cmp.Int(binary.LittleEndian, 2809, 2).Make(),
		DeepEquals,
		[]byte{0xf9, 0xa})

	cmp = New()
	c.Assert(
		cmp.Int(binary.LittleEndian, -2809, 2).Make(),
		DeepEquals,
		[]byte{0x7, 0xf5})

	cmp = New()
	c.Assert(
		cmp.Int(binary.BigEndian, 2809, 2).Make(),
		DeepEquals,
		[]byte{0xa, 0xf9})

	cmp = New()
	c.Assert(
		cmp.Int(binary.BigEndian, -2809, 2).Make(),
		DeepEquals,
		[]byte{0xf5, 0x7})
}

func (s *MySuite) TestRead(c *C) {
	cmp := NewWithRW(bytes.NewBuffer([]byte{1, 0, 2, 3, 0}))
	var b interface{}
	b, _ = cmp.ReadByte()
	c.Assert(b, Equals, byte(1))
	b, _ = cmp.ReadUint16be()
	c.Assert(b, Equals, uint16(2))
	b, _ = cmp.ReadUint16le()
	c.Assert(b, Equals, uint16(3))

	cmp = NewWithRW(bytes.NewBuffer([]byte{0, 1, 255, 254, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}))

	b, _ = cmp.ReadUint16be()
	c.Assert(b, Equals, uint16(1))
	b, _ = cmp.ReadUint16be()
	c.Assert(b, Equals, uint16(65534))
	b, _ = cmp.ReadUint32be()
	c.Assert(b, Equals, uint32(1))
	b, _ = cmp.ReadUint64be()
	c.Assert(b, Equals, uint64(1))

	cmp = NewWithRW(bytes.NewBuffer([]byte{9, 0, 0}))
	u := cmp.ReadUint(binary.LittleEndian, 3)
	c.Assert(u, DeepEquals, uint64(9))

	cmp = NewWithRW(bytes.NewBuffer([]byte{0, 0, 9}))
	u = cmp.ReadUint(binary.BigEndian, 3)
	c.Assert(u, DeepEquals, uint64(9))

	cmp = NewWithRW(bytes.NewBuffer([]byte{247, 255, 255}))
	z := cmp.ReadInt(binary.LittleEndian, 3)
	c.Assert(z, DeepEquals, int64(-9))

	cmp = NewWithRW(bytes.NewBuffer([]byte{255, 255, 247}))
	z = cmp.ReadInt(binary.BigEndian, 3)
	c.Assert(z, DeepEquals, int64(-9))

	cmp = NewWithRW(bytes.NewBuffer([]byte{0xf9, 0xa}))
	z = cmp.ReadInt(binary.LittleEndian, 2)
	c.Assert(z, DeepEquals, int64(2809))

	cmp = NewWithRW(bytes.NewBuffer([]byte{0x7, 0xf5}))
	z = cmp.ReadInt(binary.LittleEndian, 2)
	c.Assert(z, DeepEquals, int64(-2809))

	cmp = NewWithRW(bytes.NewBuffer([]byte{0xa, 0xf9}))
	z = cmp.ReadInt(binary.BigEndian, 2)
	c.Assert(z, DeepEquals, int64(2809))

	cmp = NewWithRW(bytes.NewBuffer([]byte{0xf5, 0x7}))
	z = cmp.ReadInt(binary.BigEndian, 2)
	c.Assert(z, DeepEquals, int64(-2809))
}
