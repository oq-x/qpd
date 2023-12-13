package qpd

import "io"

type PlayerData struct {
	Name       string
	X, Y, Z    int16
	Yaw, Pitch byte
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) Encoder {
	return Encoder{w}
}

func (e Encoder) Encode(p PlayerData) {
	e.writeString(p.Name)
	e.writeShortLE(p.X)
	e.writeShortLE(p.Y)
	e.writeShortLE(p.Z)
	e.writeByte(p.Yaw)
	e.writeByte(p.Pitch)
}

func (e Encoder) writeByte(i byte) {
	e.w.Write([]byte{i})
}

func (e Encoder) writeShortLE(i int16) {
	e.w.Write([]byte{byte(i), byte(i >> 8)})
}

func (e Encoder) writeString(s string) {
	e.writeShortLE(int16(len(s)))
	e.w.Write([]byte(s))
}
