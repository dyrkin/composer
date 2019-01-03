package composer

import (
	"bytes"
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
}
