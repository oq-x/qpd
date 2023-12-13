package qpd

import "io"

type Decoder struct {
	r io.Reader
}

func (d Decoder) Decode(p *PlayerData) {
	d.readString(&p.Name)
	d.readShortLE(&p.X)
	d.readShortLE(&p.Y)
	d.readShortLE(&p.Z)
	d.readByte(&p.Yaw)
	d.readByte(&p.Pitch)
}

func (d Decoder) readByte(i *byte) {
	var b [1]byte
	d.r.Read(b[:])
	*i = b[0]
}

func (d Decoder) readShortLE(i *int16) {
	var b [2]byte
	d.r.Read(b[:])
	*i = int16(b[0]) | int16(b[1])<<8
}

func (d Decoder) readString(s *string) {
	var l int16
	d.readShortLE(&l)
	b := make([]byte, l)
	d.r.Read(b)

	*s = string(b)
}
