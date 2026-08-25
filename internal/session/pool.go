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

func (p *Pool) Acquire(deviceID string, payload []byte) *Value {
	value := p.pool.Get().(*Value)
	value.DeviceID = deviceID
	value.Payload = append(value.Payload[:0], payload...)
	return value
}

func (p *Pool) Release(value *Value) {
	value.DeviceID = ""
	value.Payload = value.Payload[:0]
	p.pool.Put(value)
}
