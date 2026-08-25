package document

import "fmt"

type Document struct {
	ID      string
	Payload string
	Ready   bool
}

type Builder struct{}

func (Builder) Build(id, payload string) (doc *Document, err error) {
	working := &Document{ID: id}
	defer func() {
		if recovered := recover(); recovered != nil {
			// 失败时不外泄半成品：doc 保持 nil，仅返回错误。
			err = fmt.Errorf("document build failed: %v", recovered)
		}
	}()
	working.Payload = payload
	if payload == "panic" {
		panic("decoder panic")
	}
	working.Ready = true
	return working, nil
}
