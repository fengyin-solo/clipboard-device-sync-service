package model

import (
	"strings"
	"time"
)

type Group struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g *Group) Validate() error {
	g.Name = strings.TrimSpace(g.Name)
	if g.Name == "" {
		return NewValidationError("name", "分组名称不能为空")
	}
	if len(g.Name) > 100 {
		return NewValidationError("name", "分组名称过长")
	}
	g.Color = strings.TrimSpace(g.Color)
	if g.Color == "" {
		g.Color = "#000000"
	}
	return nil
}

type GroupFilter struct {
	Keyword string
}

func (f GroupFilter) Match(g *Group) bool {
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(g.Name), k) {
			return false
		}
	}
	return true
}
