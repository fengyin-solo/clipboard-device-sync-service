package model

import (
	"strings"
	"time"
)

type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Tag) Validate() error {
	t.Name = strings.TrimSpace(t.Name)
	if t.Name == "" {
		return NewValidationError("name", "标签名称不能为空")
	}
	if len(t.Name) > 50 {
		return NewValidationError("name", "标签名称过长")
	}
	t.Color = strings.TrimSpace(t.Color)
	if t.Color == "" {
		t.Color = "#3b82f6"
	}
	return nil
}

type TagFilter struct {
	Keyword string
}

func (f TagFilter) Match(t *Tag) bool {
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(t.Name), k) {
			return false
		}
	}
	return true
}
