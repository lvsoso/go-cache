package http

import (
	"bytes"
	"log"
	"net/http"
)

type rebalanceHandler struct {
	*Server
}

func (h *rebalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	go h.rebalance()
}

func (h *rebalanceHandler) rebalance() {
	log.Println("start rebalance")
	s := h.NewScanner()
	defer s.Close()
	c := &http.Client{}
	for s.Scan() {
		k := s.Key()
		n, ok := h.ShouldProcess(k)
		log.Println("k -> ", k)
		log.Println("n -> ", n)
		log.Println("ok -> ", ok)
		if !ok {
			r, _ := http.NewRequest(http.MethodPut, "http://"+n+":12345/cache/"+k, bytes.NewReader(s.Value()))
			c.Do(r)
			h.Del(k)
		}
	}
}

func (s *Server) rebalanceHandler() http.Handler {
	return &rebalanceHandler{s}
}
