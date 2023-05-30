package parse

import (
	"math"
)

type ByteBuf struct {
	rIndex      int
	wIndex      int
	capacity    int
	bytes       []byte
	mark_rIndex int
	mark_wIndex int
}

func ToByteBuf_capacity(capacity int) *ByteBuf {
	return &ByteBuf{
		rIndex:      0,
		wIndex:      0,
		capacity:    capacity,
		bytes:       make([]byte, capacity),
		mark_rIndex: 0,
		mark_wIndex: 0,
	}
}

func ToByteBuf_empty() *ByteBuf {
	return &ByteBuf{
		rIndex:      0,
		wIndex:      0,
		capacity:    16,
		bytes:       make([]byte, 16),
		mark_rIndex: 0,
		mark_wIndex: 0,
	}
}

func ToByteBuf(bytes []byte) *ByteBuf {
	return &ByteBuf{
		rIndex:      0,
		wIndex:      len(bytes),
		capacity:    len(bytes),
		bytes:       bytes,
		mark_rIndex: 0,
		mark_wIndex: 0,
	}
}

func (b *ByteBuf) checkGrow(n int) {
	oldLen := b.capacity
	minGrowLen := b.wIndex + n - oldLen
	if minGrowLen > 0 {
		prefGrowLen := oldLen >> 1
		oldBytes := b.bytes
		var newBytes []byte
		if prefGrowLen >= minGrowLen {
			b.capacity = oldLen + prefGrowLen
		} else {
			b.capacity = oldLen + minGrowLen
		}
		newBytes = make([]byte, b.capacity)
		copy(newBytes, oldBytes[:b.wIndex])
		b.bytes = newBytes
	}
}

func (b *ByteBuf) ToBytes() []byte {
	return b.bytes[b.rIndex:b.wIndex]
}

func (b *ByteBuf) MarkReaderIndex() {
	b.mark_rIndex = b.rIndex
}

func (b *ByteBuf) MarkWriterIndex() {
	b.mark_wIndex = b.wIndex
}

func (b *ByteBuf) ResetReaderIndex() {
	b.rIndex = b.mark_rIndex
}

func (b *ByteBuf) ResetWriterIndex() {
	b.wIndex = b.mark_wIndex
}

func (b *ByteBuf) Readable() bool {
	return b.rIndex < b.wIndex
}

func (b *ByteBuf) ReaderIndex() int {
	return b.rIndex
}

func (b *ByteBuf) WriterIndex() int {
	return b.wIndex
}

func (b *ByteBuf) ReadableBytes() int {
	return b.wIndex - b.rIndex
}

func (b *ByteBuf) Get_uint8() uint8 {
	e := b.bytes[b.rIndex]
	return e
}

func (b *ByteBuf) Get_int8() int8 {
	e := b.bytes[b.rIndex]
	return int8(e)
}

func (b *ByteBuf) Get_uint16(bigEndian bool) uint16 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	if bigEndian {
		return (uint16(e1) << 8) | uint16(e2)
	} else {
		return (uint16(e2) << 8) | uint16(e1)
	}
}

func (b *ByteBuf) Get_int16(bigEndian bool) int16 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	if bigEndian {
		return (int16(e1) << 8) | int16(e2)
	} else {
		return (int16(e2) << 8) | int16(e1)
	}
}

func (b *ByteBuf) Get_uint32(bigEndian bool) uint32 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	if bigEndian {
		return (uint32(e1) << 24) | (uint32(e2) << 16) | (uint32(e3) << 8) | uint32(e4)
	} else {
		return (uint32(e4) << 24) | (uint32(e3) << 16) | (uint32(e2) << 8) | uint32(e1)
	}
}

func (b *ByteBuf) Get_int32(bigEndian bool) int32 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	if bigEndian {
		return (int32(e1) << 24) | (int32(e2) << 16) | (int32(e3) << 8) | int32(e4)
	} else {
		return (int32(e4) << 24) | (int32(e3) << 16) | (int32(e2) << 8) | int32(e1)
	}
}

func (b *ByteBuf) Get_uint64(bigEndian bool) uint64 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	e5 := b.bytes[b.rIndex+4]
	e6 := b.bytes[b.rIndex+5]
	e7 := b.bytes[b.rIndex+6]
	e8 := b.bytes[b.rIndex+7]
	if bigEndian {
		return (uint64(e1) << 56) | (uint64(e2) << 48) | (uint64(e3) << 40) | (uint64(e4) << 32) | (uint64(e5) << 24) | (uint64(e6) << 16) | (uint64(e7) << 8) | uint64(e8)
	} else {
		return (uint64(e8) << 56) | (uint64(e7) << 48) | (uint64(e6) << 40) | (uint64(e5) << 32) | (uint64(e4) << 24) | (uint64(e3) << 16) | (uint64(e2) << 8) | uint64(e1)
	}
}

func (b *ByteBuf) Get_int64(bigEndian bool) int64 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	e5 := b.bytes[b.rIndex+4]
	e6 := b.bytes[b.rIndex+5]
	e7 := b.bytes[b.rIndex+6]
	e8 := b.bytes[b.rIndex+7]
	if bigEndian {
		return (int64(e1) << 56) | (int64(e2) << 48) | (int64(e3) << 40) | (int64(e4) << 32) | (int64(e5) << 24) | (int64(e6) << 16) | (int64(e7) << 8) | int64(e8)
	} else {
		return (int64(e8) << 56) | (int64(e7) << 48) | (int64(e6) << 40) | (int64(e5) << 32) | (int64(e4) << 24) | (int64(e3) << 16) | (int64(e2) << 8) | int64(e1)
	}
}

func (b *ByteBuf) Get_float32(bigEndian bool) float32 {
	readUint32 := b.Get_uint32(bigEndian)
	return math.Float32frombits(readUint32)
}

func (b *ByteBuf) Get_float64(bigEndian bool) float64 {
	readUint64 := b.Get_uint64(bigEndian)
	return math.Float64frombits(readUint64)
}

func (b *ByteBuf) Get_bytes(n int) []byte {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	return bytes
}

func (b *ByteBuf) Get_string_utf8(n int) string {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	return string(bytes)
}

func (b *ByteBuf) Read_uint8() uint8 {
	e := b.bytes[b.rIndex]
	b.rIndex += 1
	return e
}

func (b *ByteBuf) Read_int8() int8 {
	e := b.bytes[b.rIndex]
	b.rIndex += 1
	return int8(e)
}

func (b *ByteBuf) Read_uint16(bigEndian bool) uint16 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	b.rIndex += 2
	if bigEndian {
		return (uint16(e1) << 8) | uint16(e2)
	} else {
		return (uint16(e2) << 8) | uint16(e1)
	}
}

func (b *ByteBuf) Read_int16(bigEndian bool) int16 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	b.rIndex += 2
	if bigEndian {
		return (int16(e1) << 8) | int16(e2)
	} else {
		return (int16(e2) << 8) | int16(e1)
	}
}

func (b *ByteBuf) Read_uint32(bigEndian bool) uint32 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	b.rIndex += 4
	if bigEndian {
		return (uint32(e1) << 24) | (uint32(e2) << 16) | (uint32(e3) << 8) | uint32(e4)
	} else {
		return (uint32(e4) << 24) | (uint32(e3) << 16) | (uint32(e2) << 8) | uint32(e1)
	}
}

func (b *ByteBuf) Read_int32(bigEndian bool) int32 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	b.rIndex += 4
	if bigEndian {
		return (int32(e1) << 24) | (int32(e2) << 16) | (int32(e3) << 8) | int32(e4)
	} else {
		return (int32(e4) << 24) | (int32(e3) << 16) | (int32(e2) << 8) | int32(e1)
	}
}

func (b *ByteBuf) Read_uint64(bigEndian bool) uint64 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	e5 := b.bytes[b.rIndex+4]
	e6 := b.bytes[b.rIndex+5]
	e7 := b.bytes[b.rIndex+6]
	e8 := b.bytes[b.rIndex+7]
	b.rIndex += 8
	if bigEndian {
		return (uint64(e1) << 56) | (uint64(e2) << 48) | (uint64(e3) << 40) | (uint64(e4) << 32) | (uint64(e5) << 24) | (uint64(e6) << 16) | (uint64(e7) << 8) | uint64(e8)
	} else {
		return (uint64(e8) << 56) | (uint64(e7) << 48) | (uint64(e6) << 40) | (uint64(e5) << 32) | (uint64(e4) << 24) | (uint64(e3) << 16) | (uint64(e2) << 8) | uint64(e1)
	}
}

func (b *ByteBuf) Read_int64(bigEndian bool) int64 {
	e1 := b.bytes[b.rIndex]
	e2 := b.bytes[b.rIndex+1]
	e3 := b.bytes[b.rIndex+2]
	e4 := b.bytes[b.rIndex+3]
	e5 := b.bytes[b.rIndex+4]
	e6 := b.bytes[b.rIndex+5]
	e7 := b.bytes[b.rIndex+6]
	e8 := b.bytes[b.rIndex+7]
	b.rIndex += 8
	if bigEndian {
		return (int64(e1) << 56) | (int64(e2) << 48) | (int64(e3) << 40) | (int64(e4) << 32) | (int64(e5) << 24) | (int64(e6) << 16) | (int64(e7) << 8) | int64(e8)
	} else {
		return (int64(e8) << 56) | (int64(e7) << 48) | (int64(e6) << 40) | (int64(e5) << 32) | (int64(e4) << 24) | (int64(e3) << 16) | (int64(e2) << 8) | int64(e1)
	}
}

func (b *ByteBuf) Read_float32(bigEndian bool) float32 {
	readUint32 := b.Read_uint32(bigEndian)
	return math.Float32frombits(readUint32)
}

func (b *ByteBuf) Read_float64(bigEndian bool) float64 {
	readUint64 := b.Read_uint64(bigEndian)
	return math.Float64frombits(readUint64)
}

func (b *ByteBuf) Read_bytes(n int) []byte {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	b.rIndex += n
	return bytes
}

func (b *ByteBuf) Read_string_utf8(n int) string {
	bytes := b.bytes[b.rIndex : b.rIndex+n]
	b.rIndex += n
	return string(bytes)
}

func (b *ByteBuf) Write_uint8(v uint8) {
	b.checkGrow(1)
	b.bytes[b.wIndex] = v
	b.wIndex++
}

func (b *ByteBuf) Write_int8(v int8) {
	b.checkGrow(1)
	b.bytes[b.wIndex] = byte(v)
	b.wIndex++
}

func (b *ByteBuf) Write_uint16(v uint16, bigEndian bool) {
	b.checkGrow(2)
	if bigEndian {
		b.bytes[b.wIndex] = uint8(v >> 8)
		b.bytes[b.wIndex+1] = uint8(v)
	} else {
		b.bytes[b.wIndex] = uint8(v)
		b.bytes[b.wIndex+1] = uint8(v >> 8)
	}
	b.wIndex += 2
}

func (b *ByteBuf) Write_int16(v int16, bigEndian bool) {
	b.checkGrow(2)
	if bigEndian {
		b.bytes[b.wIndex] = uint8(v >> 8)
		b.bytes[b.wIndex+1] = uint8(v)
	} else {
		b.bytes[b.wIndex] = uint8(v)
		b.bytes[b.wIndex+1] = uint8(v >> 8)
	}
	b.wIndex += 2
}

func (b *ByteBuf) Write_uint32(v uint32, bigEndian bool) {
	b.checkGrow(4)
	if bigEndian {
		b.bytes[b.wIndex] = uint8(v >> 24)
		b.bytes[b.wIndex+1] = uint8(v >> 16)
		b.bytes[b.wIndex+2] = uint8(v >> 8)
		b.bytes[b.wIndex+3] = uint8(v)
	} else {
		b.bytes[b.wIndex] = uint8(v)
		b.bytes[b.wIndex+1] = uint8(v >> 8)
		b.bytes[b.wIndex+2] = uint8(v >> 16)
		b.bytes[b.wIndex+3] = uint8(v >> 24)
	}
	b.wIndex += 4
}

func (b *ByteBuf) Write_int32(v int32, bigEndian bool) {
	b.checkGrow(4)
	if bigEndian {
		b.bytes[b.wIndex] = uint8(v >> 24)
		b.bytes[b.wIndex+1] = uint8(v >> 16)
		b.bytes[b.wIndex+2] = uint8(v >> 8)
		b.bytes[b.wIndex+3] = uint8(v)
	} else {
		b.bytes[b.wIndex] = uint8(v)
		b.bytes[b.wIndex+1] = uint8(v >> 8)
		b.bytes[b.wIndex+2] = uint8(v >> 16)
		b.bytes[b.wIndex+3] = uint8(v >> 24)
	}
	b.wIndex += 4
}

func (b *ByteBuf) Write_uint64(v uint64, bigEndian bool) {
	b.checkGrow(8)
	if bigEndian {
		b.bytes[b.wIndex] = uint8(v >> 56)
		b.bytes[b.wIndex+1] = uint8(v >> 48)
		b.bytes[b.wIndex+2] = uint8(v >> 40)
		b.bytes[b.wIndex+3] = uint8(v >> 32)
		b.bytes[b.wIndex+4] = uint8(v >> 24)
		b.bytes[b.wIndex+5] = uint8(v >> 16)
		b.bytes[b.wIndex+6] = uint8(v >> 8)
		b.bytes[b.wIndex+7] = uint8(v)
	} else {
		b.bytes[b.wIndex] = uint8(v)
		b.bytes[b.wIndex+1] = uint8(v >> 8)
		b.bytes[b.wIndex+2] = uint8(v >> 16)
		b.bytes[b.wIndex+3] = uint8(v >> 24)
		b.bytes[b.wIndex+4] = uint8(v >> 32)
		b.bytes[b.wIndex+5] = uint8(v >> 40)
		b.bytes[b.wIndex+6] = uint8(v >> 48)
		b.bytes[b.wIndex+7] = uint8(v >> 56)
	}
	b.wIndex += 8
}

func (b *ByteBuf) Write_int64(v int64, bigEndian bool) {
	b.checkGrow(8)
	if bigEndian {
		b.bytes[b.wIndex] = uint8(v >> 56)
		b.bytes[b.wIndex+1] = uint8(v >> 48)
		b.bytes[b.wIndex+2] = uint8(v >> 40)
		b.bytes[b.wIndex+3] = uint8(v >> 32)
		b.bytes[b.wIndex+4] = uint8(v >> 24)
		b.bytes[b.wIndex+5] = uint8(v >> 16)
		b.bytes[b.wIndex+6] = uint8(v >> 8)
		b.bytes[b.wIndex+7] = uint8(v)
	} else {
		b.bytes[b.wIndex] = uint8(v)
		b.bytes[b.wIndex+1] = uint8(v >> 8)
		b.bytes[b.wIndex+2] = uint8(v >> 16)
		b.bytes[b.wIndex+3] = uint8(v >> 24)
		b.bytes[b.wIndex+4] = uint8(v >> 32)
		b.bytes[b.wIndex+5] = uint8(v >> 40)
		b.bytes[b.wIndex+6] = uint8(v >> 48)
		b.bytes[b.wIndex+7] = uint8(v >> 56)
	}
	b.wIndex += 8
}

func (b *ByteBuf) Write_float32(v float32, bigEndian bool) {
	b.Write_uint32(math.Float32bits(v), bigEndian)
}

func (b *ByteBuf) Write_float64(v float64, bigEndian bool) {
	b.Write_uint64(math.Float64bits(v), bigEndian)
}

func (b *ByteBuf) Write_bytes(bytes []byte) {
	b.checkGrow(len(bytes))
	copy(b.bytes[b.wIndex:], bytes)
	b.wIndex += len(bytes)
}

func (b *ByteBuf) Write_zero(n int) {
	if n <= 0 {
		return
	}
	b.checkGrow(n)
	b.wIndex += n
}

func (b *ByteBuf) Write_string_utf8(v string) {
	bytes := []byte(v)
	b.Write_bytes(bytes)
}

func (b *ByteBuf) Skip(n int) {
	b.rIndex += n
}
