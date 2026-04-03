package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// ─── Response types ───────────────────────────────────────────────────────────

type EduIDResponse struct {
	EduID string `json:"eduID"`
}

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

// ─── EduID generation ─────────────────────────────────────────────────────────

// generateEduID returns "<email>+<6-digit-number>", e.g. "alice@bar.com+483921".
func generateEduID(email string) string {
	n := rand.Intn(900000) + 100000 // guaranteed 6 digits: 100000–999999
	return fmt.Sprintf("%s+%06d", email, n)
}

// ─── Handlers ─────────────────────────────────────────────────────────────────

// GET /user-attributes/edu-id?email=<email>
func handleGetEduID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET is supported")
		return
	}

	email := strings.TrimSpace(r.URL.Query().Get("email"))
	if email == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "Query parameter 'email' is required")
		return
	}

	eduID := generateEduID(email)

	log.Printf("[%s] GET /user-attributes/edu-id email=%s → eduID=%s",
		time.Now().Format(time.RFC3339), email, eduID)

	writeJSON(w, http.StatusOK, EduIDResponse{EduID: eduID})
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, errCode, desc string) {
	writeJSON(w, status, ErrorResponse{Error: errCode, Description: desc})
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	rand.Seed(time.Now().UnixNano()) //nolint:staticcheck

	mux := http.NewServeMux()
	mux.HandleFunc("/user-attributes/edu-id", handleGetEduID)

	addr := ":8081"
	log.Printf("EduID Service starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
