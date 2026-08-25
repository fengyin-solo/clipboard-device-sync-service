package handler

import (
	"net/http"

	"clipboard/pkg/httpx"
)

func (s *Server) registerStatsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/stats/overview", s.statsOverview)
}

func (s *Server) statsOverview(w http.ResponseWriter, r *http.Request) {
	stats, err := s.svc.StatsOverview()
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, stats)
}
