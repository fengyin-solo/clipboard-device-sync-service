package handler

import (
	"net/http"

	"clipboard/internal/model"
	"clipboard/pkg/httpx"
)

func (s *Server) registerGroupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/groups", s.createGroup)
	mux.HandleFunc("GET /api/groups", s.listGroups)
	mux.HandleFunc("GET /api/groups/{id}", s.getGroup)
	mux.HandleFunc("PUT /api/groups/{id}", s.updateGroup)
	mux.HandleFunc("DELETE /api/groups/{id}", s.deleteGroup)
}

type createGroupRequest struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	SortOrder int    `json:"sort_order"`
}

func (s *Server) createGroup(w http.ResponseWriter, r *http.Request) {
	var req createGroupRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	group, err := s.svc.CreateGroup(model.Group{
		Name:      req.Name,
		Color:     req.Color,
		SortOrder: req.SortOrder,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, group)
}

func (s *Server) listGroups(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.GroupFilter{
		Keyword: r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListGroups(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getGroup(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	group, err := s.svc.GetGroup(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, group)
}

type updateGroupRequest struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	SortOrder int    `json:"sort_order"`
}

func (s *Server) updateGroup(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateGroupRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	group, err := s.svc.UpdateGroup(id, model.Group{
		Name:      req.Name,
		Color:     req.Color,
		SortOrder: req.SortOrder,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, group)
}

func (s *Server) deleteGroup(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteGroup(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
