package frame

type Decoder struct {
	buffer []byte
}

type Frame struct {
	Payload []byte
}

func NewDecoder(size int) *Decoder {
	return &Decoder{buffer: make([]byte, size)}
}

func (d *Decoder) Decode(text string) Frame {
	n := copy(d.buffer, text)
	return Frame{Payload: d.buffer[:n]}
}
