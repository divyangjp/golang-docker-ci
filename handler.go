package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"io"
	"net/http"
)

type handler struct {
	key   []byte
	stats map[string]uint64
}

func (h handler) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

func (h handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	h.stats["requests"] += 1

	body, _ := io.ReadAll(r.Body)
	out := createMAC(body, h.key)

	enc := json.NewEncoder(w)
	enc.Encode(out)
}

func (h handler) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	enc.Encode(h.stats)
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
