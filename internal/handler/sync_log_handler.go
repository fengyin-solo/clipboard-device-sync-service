package handler

import (
	"net/http"

	"clipboard/internal/model"
	"clipboard/pkg/httpx"
)

func (s *Server) registerSyncLogRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/sync-logs", s.createSyncLog)
	mux.HandleFunc("GET /api/sync-logs", s.listSyncLogs)
	mux.HandleFunc("GET /api/sync-logs/{id}", s.getSyncLog)
	mux.HandleFunc("PUT /api/sync-logs/{id}", s.updateSyncLog)
	mux.HandleFunc("DELETE /api/sync-logs/{id}", s.deleteSyncLog)
}

type createSyncLogRequest struct {
	DeviceID   string `json:"device_id"`
	Direction  string `json:"direction"`
	EntryCount int    `json:"entry_count"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func (s *Server) createSyncLog(w http.ResponseWriter, r *http.Request) {
	var req createSyncLogRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	sl, err := s.svc.CreateSyncLog(model.SyncLog{
		DeviceID:   req.DeviceID,
		Direction:  req.Direction,
		EntryCount: req.EntryCount,
		Status:     req.Status,
		Message:    req.Message,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, sl)
}

func (s *Server) listSyncLogs(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.SyncLogFilter{
		DeviceID:  r.URL.Query().Get("device_id"),
		Direction: r.URL.Query().Get("direction"),
		Status:    r.URL.Query().Get("status"),
	}
	items, total, err := s.svc.ListSyncLogs(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getSyncLog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sl, err := s.svc.GetSyncLog(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, sl)
}

type updateSyncLogRequest struct {
	DeviceID   string `json:"device_id"`
	Direction  string `json:"direction"`
	EntryCount int    `json:"entry_count"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func (s *Server) updateSyncLog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateSyncLogRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	sl, err := s.svc.UpdateSyncLog(id, model.SyncLog{
		DeviceID:   req.DeviceID,
		Direction:  req.Direction,
		EntryCount: req.EntryCount,
		Status:     req.Status,
		Message:    req.Message,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, sl)
}

func (s *Server) deleteSyncLog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteSyncLog(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
