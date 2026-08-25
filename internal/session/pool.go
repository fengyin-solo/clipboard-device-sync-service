package session

import "sync"

type Value struct {
	DeviceID string
	Payload  []byte
}

type Pool struct {
	pool sync.Pool
}

func NewPool() *Pool {
	p := &Pool{}
	p.pool.New = func() any { return &Value{} }
	return p
}

// Acquire 从池中取出一个 Value，并按入参填充字段。
// 每次取用都无条件清零既有字段后再填充，确保空入参不会带上一次复用残留的数据。
func (p *Pool) Acquire(deviceID string, payload []byte) *Value {
	value := p.pool.Get().(*Value)
	value.DeviceID = deviceID
	value.Payload = append(value.Payload[:0], payload...)
	return value
}

// Release 将 Value 归还到池中，并清空字段，避免下一次取用时残留旧数据。
func (p *Pool) Release(value *Value) {
	value.DeviceID = ""
	value.Payload = value.Payload[:0]
	p.pool.Put(value)
}
