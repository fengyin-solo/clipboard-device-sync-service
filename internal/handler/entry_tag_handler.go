package handler

import (
	"net/http"

	"clipboard/internal/model"
	"clipboard/pkg/httpx"
)

func (s *Server) registerEntryTagRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/entry-tags", s.createEntryTag)
	mux.HandleFunc("GET /api/entry-tags", s.listEntryTags)
	mux.HandleFunc("GET /api/entry-tags/{id}", s.getEntryTag)
	mux.HandleFunc("DELETE /api/entry-tags/{id}", s.deleteEntryTag)
}

type createEntryTagRequest struct {
	EntryID string `json:"entry_id"`
	TagID   string `json:"tag_id"`
}

func (s *Server) createEntryTag(w http.ResponseWriter, r *http.Request) {
	var req createEntryTagRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	et, err := s.svc.CreateEntryTag(model.EntryTag{
		EntryID: req.EntryID,
		TagID:   req.TagID,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, et)
}

func (s *Server) listEntryTags(w http.ResponseWriter, r *http.Request) {
	filter := model.EntryTagFilter{
		EntryID: r.URL.Query().Get("entry_id"),
		TagID:   r.URL.Query().Get("tag_id"),
	}
	items, err := s.svc.ListEntryTags(filter)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, items)
}

func (s *Server) getEntryTag(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	et, err := s.svc.GetEntryTag(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, et)
}

func (s *Server) deleteEntryTag(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteEntryTag(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
