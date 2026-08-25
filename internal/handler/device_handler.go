package handler

import (
	"net/http"

	"clipboard/internal/model"
	"clipboard/pkg/httpx"
)

func (s *Server) registerDeviceRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/devices", s.createDevice)
	mux.HandleFunc("GET /api/devices", s.listDevices)
	mux.HandleFunc("GET /api/devices/{id}", s.getDevice)
	mux.HandleFunc("PUT /api/devices/{id}", s.updateDevice)
	mux.HandleFunc("DELETE /api/devices/{id}", s.deleteDevice)
}

type createDeviceRequest struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
}

func (s *Server) createDevice(w http.ResponseWriter, r *http.Request) {
	var req createDeviceRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	device, err := s.svc.CreateDevice(model.Device{
		Name:     req.Name,
		Platform: req.Platform,
		Token:    req.Token,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, device)
}

func (s *Server) listDevices(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.DeviceFilter{
		Platform: r.URL.Query().Get("platform"),
		Keyword:  r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListDevices(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getDevice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	device, err := s.svc.GetDevice(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, device)
}

type updateDeviceRequest struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
}

func (s *Server) updateDevice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateDeviceRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	device, err := s.svc.UpdateDevice(id, model.Device{
		Name:     req.Name,
		Platform: req.Platform,
		Token:    req.Token,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, device)
}

func (s *Server) deleteDevice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteDevice(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
