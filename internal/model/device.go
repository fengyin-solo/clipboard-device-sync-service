package model

import (
	"strings"
	"time"
)

const (
	DevicePlatformIOS     = "ios"
	DevicePlatformAndroid = "android"
	DevicePlatformDesktop = "desktop"
	DevicePlatformWeb     = "web"
)

type Device struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Platform   string    `json:"platform"`
	LastSyncAt time.Time `json:"last_sync_at"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
}

func (d *Device) Validate() error {
	d.Name = strings.TrimSpace(d.Name)
	if d.Name == "" {
		return NewValidationError("name", "设备名称不能为空")
	}
	if d.Platform == "" {
		return NewValidationError("platform", "平台不能为空")
	}
	if d.Platform != DevicePlatformIOS && d.Platform != DevicePlatformAndroid &&
		d.Platform != DevicePlatformDesktop && d.Platform != DevicePlatformWeb {
		return NewValidationError("platform", "平台类型不合法")
	}
	return nil
}

type DeviceFilter struct {
	Platform string
	Keyword  string
}

func (f DeviceFilter) Match(d *Device) bool {
	if f.Platform != "" && d.Platform != f.Platform {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(d.Name), k) {
			return false
		}
	}
	return true
}
