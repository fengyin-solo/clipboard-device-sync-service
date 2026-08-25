package handler

import (
	"net/http"
	"strconv"

	"clipboard/internal/model"
	"clipboard/pkg/httpx"
)

func (s *Server) registerEntryRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/entries", s.createEntry)
	mux.HandleFunc("GET /api/entries", s.listEntries)
	mux.HandleFunc("GET /api/entries/{id}", s.getEntry)
	mux.HandleFunc("PUT /api/entries/{id}", s.updateEntry)
	mux.HandleFunc("DELETE /api/entries/{id}", s.deleteEntry)
	mux.HandleFunc("POST /api/entries/batch", s.batchCreateEntries)
	mux.HandleFunc("POST /api/entries/batch-delete", s.batchDeleteEntries)
	mux.HandleFunc("POST /api/entries/batch-status", s.batchUpdateEntryStatus)
	mux.HandleFunc("POST /api/entries/batch-pin", s.batchPinEntries)
}

type createEntryRequest struct {
	Content      string `json:"content"`
	ContentType  string `json:"content_type"`
	GroupID      string `json:"group_id"`
	SizeBytes    int64  `json:"size_bytes"`
	Pinned       bool   `json:"pinned"`
	SourceDevice string `json:"source_device"`
	Status       string `json:"status"`
}

func (s *Server) createEntry(w http.ResponseWriter, r *http.Request) {
	var req createEntryRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	entry, err := s.svc.CreateEntry(model.Entry{
		Content:      req.Content,
		ContentType:  req.ContentType,
		GroupID:      req.GroupID,
		SizeBytes:    req.SizeBytes,
		Pinned:       req.Pinned,
		SourceDevice: req.SourceDevice,
		Status:       req.Status,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, entry)
}

func (s *Server) listEntries(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	pinnedStr := r.URL.Query().Get("pinned")
	var pinned *bool
	if pinnedStr != "" {
		b, _ := strconv.ParseBool(pinnedStr)
		pinned = &b
	}
	filter := model.EntryFilter{
		GroupID:     r.URL.Query().Get("group_id"),
		ContentType: r.URL.Query().Get("content_type"),
		Status:      r.URL.Query().Get("status"),
		Keyword:     r.URL.Query().Get("keyword"),
		Pinned:      pinned,
	}
	items, total, err := s.svc.ListEntries(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	entry, err := s.svc.GetEntry(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, entry)
}

type updateEntryRequest struct {
	Content      string `json:"content"`
	ContentType  string `json:"content_type"`
	GroupID      string `json:"group_id"`
	SizeBytes    int64  `json:"size_bytes"`
	Pinned       bool   `json:"pinned"`
	SourceDevice string `json:"source_device"`
	Status       string `json:"status"`
}

func (s *Server) updateEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateEntryRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	entry, err := s.svc.UpdateEntry(id, model.Entry{
		Content:      req.Content,
		ContentType:  req.ContentType,
		GroupID:      req.GroupID,
		SizeBytes:    req.SizeBytes,
		Pinned:       req.Pinned,
		SourceDevice: req.SourceDevice,
		Status:       req.Status,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, entry)
}

func (s *Server) deleteEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteEntry(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type batchCreateEntriesRequest struct {
	Items []createEntryRequest `json:"items"`
}

func (s *Server) batchCreateEntries(w http.ResponseWriter, r *http.Request) {
	var req batchCreateEntriesRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	inputs := make([]model.Entry, 0, len(req.Items))
	for _, item := range req.Items {
		inputs = append(inputs, model.Entry{
			Content:      item.Content,
			ContentType:  item.ContentType,
			GroupID:      item.GroupID,
			SizeBytes:    item.SizeBytes,
			Pinned:       item.Pinned,
			SourceDevice: item.SourceDevice,
			Status:       item.Status,
		})
	}
	entries, err := s.svc.BatchCreateEntries(inputs)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, entries)
}

type batchDeleteEntriesRequest struct {
	IDs []string `json:"ids"`
}

func (s *Server) batchDeleteEntries(w http.ResponseWriter, r *http.Request) {
	var req batchDeleteEntriesRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	if err := s.svc.BatchDeleteEntries(req.IDs); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type batchUpdateEntryStatusRequest struct {
	IDs    []string `json:"ids"`
	Status string   `json:"status"`
}

func (s *Server) batchUpdateEntryStatus(w http.ResponseWriter, r *http.Request) {
	var req batchUpdateEntryStatusRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	if err := s.svc.BatchUpdateEntryStatus(req.IDs, req.Status); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type batchPinEntriesRequest struct {
	IDs    []string `json:"ids"`
	Pinned bool     `json:"pinned"`
}

func (s *Server) batchPinEntries(w http.ResponseWriter, r *http.Request) {
	var req batchPinEntriesRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	if err := s.svc.BatchPinEntries(req.IDs, req.Pinned); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
