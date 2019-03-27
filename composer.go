package composer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type Composer struct {
	r io.Reader
	w io.Writer
	b *bytes.Buffer
}

func New() *Composer {
	return NewWithRW(nil)
}

func NewWithRW(rw io.ReadWriter) *Composer {
	return &Composer{rw, rw, &bytes.Buffer{}}
}

func NewWithR(r io.Reader) *Composer {
	return &Composer{r, nil, &bytes.Buffer{}}
}

func NewWithW(w io.Writer) *Composer {
	return &Composer{nil, w, &bytes.Buffer{}}
}

func (f *Composer) Byte(b byte) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Bytes(b []byte) *Composer {
	f.b.Write(b)
	return f
}

func (f *Composer) Uint8(b uint8) *Composer {
	return f.Byte(b)
}

func (f *Composer) Uintbe(b uint) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Uint16be(b uint16) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Uint32be(b uint32) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Uint64be(b uint64) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Uintle(b uint) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Uint16le(b uint16) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Uint32le(b uint32) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Uint64le(b uint64) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Int8(b int8) *Composer {
	return f.Byte(byte(b))
}

func (f *Composer) Intbe(b int) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Int16be(b int16) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Int32be(b int32) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Int64be(b int64) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Intle(b int) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Int16le(b int16) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Int32le(b int32) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Int64le(b int64) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Float32be(b float32) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Float64be(b float64) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Float32le(b float32) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Float64le(b float64) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Complex64be(b complex64) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Complex128be(b complex128) *Composer {
	binary.Write(f.b, binary.BigEndian, b)
	return f
}

func (f *Composer) Complex64le(b complex64) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Complex128le(b complex128) *Composer {
	binary.Write(f.b, binary.LittleEndian, b)
	return f
}

func (f *Composer) Uint(endianness binary.ByteOrder, t uint64, size int) *Composer {
	buf := make([]uint8, size)
	if endianness == binary.BigEndian {
		for i := 0; i < size; i++ {
			buf[i] = byte(t >> byte((size-i-1)*8))
		}
	} else {
		for i := 0; i < size; i++ {
			buf[i] = byte(t >> byte(i*8))
		}
	}
	f.b.Write(buf)
	return f
}

func (f *Composer) Int(endianness binary.ByteOrder, t int64, size int) *Composer {
	buf := make([]uint8, size)
	if endianness == binary.BigEndian {
		for i := 0; i < size; i++ {
			buf[i] = byte(t >> byte((size-i-1)*8))
		}
	} else {
		for i := 0; i < size; i++ {
			buf[i] = byte(t >> byte(i*8))
		}
	}
	f.b.Write(buf)
	return f
}

func (f *Composer) String(b string) *Composer {
	f.b.WriteString(b)
	return f
}

func (f *Composer) ReadByte() (b byte, err error) {
	var buf [1]byte
	_, err = io.ReadFull(f.r, buf[:])
	b = buf[0]
	return
}

func (f *Composer) ReadBytes(len int) (buf []byte, err error) {
	buf = make([]byte, len)
	_, err = io.ReadFull(f.r, buf)
	return
}

func (f *Composer) ReadBuf(buf []byte) (err error) {
	_, err = io.ReadFull(f.r, buf)
	return
}

func (f *Composer) ReadUint8() (byte, error) {
	return f.ReadByte()
}

func (f *Composer) ReadUintbe() (v uint, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadUint16be() (v uint16, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadUint32be() (v uint32, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadUint64be() (v uint64, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadUintle() (v uint, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadUint16le() (v uint16, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadUint32le() (v uint32, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadUint64le() (v uint64, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadInt8() (v int8, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadIntbe() (v int, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadInt16be() (v int16, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadInt32be() (v int32, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadInt64be() (v int64, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadIntle() (v int, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadInt16le() (v int16, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadInt32le() (v int32, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadInt64le() (v int64, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadFloat32be() (v float32, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadFloat64be() (v float64, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadFloat32le() (v float32, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadFloat64le() (v float64, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadComplex64be() (v complex64, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadComplex128be() (v complex128, err error) {
	err = binary.Read(f.r, binary.BigEndian, &v)
	return
}

func (f *Composer) ReadComplex64le() (v complex64, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadComplex128le() (v complex128, err error) {
	err = binary.Read(f.r, binary.LittleEndian, &v)
	return
}

func (f *Composer) ReadUint(endianness binary.ByteOrder, size int) uint64 {
	var v uint64
	buf := make([]uint8, size)
	f.r.Read(buf)
	if endianness == binary.BigEndian {
		for i := 0; i < size; i++ {
			t := buf[i]
			v = v | uint64(t)<<byte((size-i-1)*8)
		}
	} else {
		for i := 0; i < size; i++ {
			t := buf[i]
			v = v | uint64(t)<<byte(i*8)
		}
	}
	return v
}

func (f *Composer) ReadInt(endianness binary.ByteOrder, size int) int64 {
	var v int64
	buf := make([]uint8, size)
	f.r.Read(buf)
	if endianness == binary.BigEndian {
		for i := 0; i < size; i++ {
			t := buf[i]
			if i == 0 {
				v = v | int64(int8(t))<<byte((size-i-1)*8)
			} else {
				v = v | int64(t)<<byte((size-i-1)*8)
			}
		}
	} else {
		for i := 0; i < size; i++ {
			t := buf[i]
			if i != 0 {
				v = v | int64(int8(t))<<byte(i*8)
			} else {
				v = v | int64(t)<<byte(i*8)
			}
		}
	}
	return v
}

func (f *Composer) ReadString(len int) (v string, err error) {
	buf, err := f.ReadBytes(len)
	v = string(buf)
	return
}

func (f *Composer) Flush() (err error) {
	if f.w == nil {
		err = errors.New("Writer is not defined")
		return
	}
	_, err = f.w.Write(f.b.Bytes())
	if err != nil {
		return
	}
	f.b = &bytes.Buffer{}
	return
}

func (f *Composer) Make() []byte {
	return f.b.Bytes()
}
