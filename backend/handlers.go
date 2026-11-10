package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	db        *sql.DB
	adminUser string
	adminPass string
}

type Member struct {
	ID            int64  `json:"id"`
	Smeknamn      string `json:"Smeknamn"`
	Namn          string `json:"Namn"`
	Roll          string `json:"Roll"`
	Quote         string `json:"quote"`
	Bild          string `json:"bild"`
	Telefonnummer string `json:"telefonnummer"`
}

type OSTeamMember struct {
	Name         string `json:"name"`
	Personnumber string `json:"personnumber"`
	Nickname     string `json:"nickname"`
	Role         string `json:"role"`
	Allergies    string `json:"allergies"`
	Drink        string `json:"drink"`
	Gasque       bool   `json:"gasque"`
	IsLate       bool   `json:"isLate"`
}

type OSTeam struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	IsLate       bool           `json:"isLate"`
	Members      []OSTeamMember `json:"members"`
	RegisteredAt string         `json:"registeredAt"`
}

type TournamentTeam struct {
	ID        int64    `json:"id"`
	Game      string   `json:"game"`
	Name      string   `json:"name"`
	Players   []string `json:"players"`
	CreatedAt string   `json:"createdAt"`
}

type TournamentMatch struct {
	ID        int64  `json:"id"`
	Game      string `json:"game"`
	Round     string `json:"round"`
	Team1     string `json:"team1"`
	Team2     string `json:"team2"`
	Winner    string `json:"winner"`
	MatchTime string `json:"matchTime"`
}

// --- Helpers ---

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

func decode(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

func pathID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// --- Health & Admin Login ---

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	if err := s.db.Ping(); err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) adminLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := decode(r, &body); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}
	if body.Username == s.adminUser && body.Password == s.adminPass {
		writeJSON(w, http.StatusOK, map[string]any{"success": true, "token": "piraya-admin-token"})
		return
	}
	writeJSON(w, http.StatusUnauthorized, map[string]any{"success": false, "error": "invalid credentials"})
}

// --- Members CRUD ---

func (s *Server) listMembers(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query(`SELECT id, smeknamn, namn, roll, quote, bild, telefonnummer FROM members ORDER BY id ASC`)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	out := []Member{}
	for rows.Next() {
		var m Member
		if err := rows.Scan(&m.ID, &m.Smeknamn, &m.Namn, &m.Roll, &m.Quote, &m.Bild, &m.Telefonnummer); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		out = append(out, m)
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) createMember(w http.ResponseWriter, r *http.Request) {
	var m Member
	if err := decode(r, &m); err != nil || strings.TrimSpace(m.Smeknamn) == "" || strings.TrimSpace(m.Namn) == "" {
		writeErr(w, http.StatusBadRequest, "Smeknamn and Namn are required")
		return
	}
	res, err := s.db.Exec(
		`INSERT INTO members (smeknamn, namn, roll, quote, bild, telefonnummer) VALUES (?, ?, ?, ?, ?, ?)`,
		m.Smeknamn, m.Namn, m.Roll, m.Quote, m.Bild, m.Telefonnummer,
	)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	m.ID, _ = res.LastInsertId()
	writeJSON(w, http.StatusCreated, m)
}

func (s *Server) updateMember(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	var m Member
	if err := decode(r, &m); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}
	res, err := s.db.Exec(
		`UPDATE members SET smeknamn=?, namn=?, roll=?, quote=?, bild=?, telefonnummer=? WHERE id=?`,
		m.Smeknamn, m.Namn, m.Roll, m.Quote, m.Bild, m.Telefonnummer, id,
	)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		writeErr(w, http.StatusNotFound, "member not found")
		return
	}
	m.ID = id
	writeJSON(w, http.StatusOK, m)
}

func (s *Server) deleteMember(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	res, err := s.db.Exec(`DELETE FROM members WHERE id=?`, id)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		writeErr(w, http.StatusNotFound, "member not found")
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"deleted": true})
}

// --- Otympliga Spelen (OS) ---

func (s *Server) osStatus(w http.ResponseWriter, r *http.Request) {
	var isOpenInt, maxTeams, maxGasque int
	err := s.db.QueryRow(`SELECT is_open, max_teams, max_gasque FROM os_settings WHERE id=1`).Scan(&isOpenInt, &maxTeams, &maxGasque)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	rows, err := s.db.Query(`SELECT id, name, description, is_late, members_json, registered_at FROM os_teams ORDER BY id ASC`)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	teams := []OSTeam{}
	usedGasque := 0
	for rows.Next() {
		var t OSTeam
		var lateInt int
		var membersJSON string
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &lateInt, &membersJSON, &t.RegisteredAt); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		t.IsLate = lateInt != 0
		_ = json.Unmarshal([]byte(membersJSON), &t.Members)
		if t.Members == nil {
			t.Members = []OSTeamMember{}
		}
		for _, m := range t.Members {
			if m.Gasque {
				usedGasque++
			}
		}
		teams = append(teams, t)
	}

	spotsLeft := max(0, maxTeams-len(teams))
	gasqueLeft := max(0, maxGasque-usedGasque)
	writeJSON(w, http.StatusOK, map[string]any{
		"isOpen":     isOpenInt != 0,
		"isFull":     spotsLeft == 0,
		"spotsLeft":  spotsLeft,
		"gasqueLeft": gasqueLeft,
		"maxTeams":   maxTeams,
		"maxGasque":  maxGasque,
		"teams":      teams,
	})
}

func (s *Server) osRegister(w http.ResponseWriter, r *http.Request) {
	var req OSTeam
	if err := decode(r, &req); err != nil || strings.TrimSpace(req.Name) == "" || len(req.Members) == 0 {
		writeErr(w, http.StatusBadRequest, "name and members are required")
		return
	}

	tx, err := s.db.Begin()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback()

	var isOpenInt, maxTeams, maxGasque int
	if err := tx.QueryRow(`SELECT is_open, max_teams, max_gasque FROM os_settings WHERE id=1`).Scan(&isOpenInt, &maxTeams, &maxGasque); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if isOpenInt == 0 {
		writeErr(w, http.StatusForbidden, "Registration is closed")
		return
	}

	var teamCount int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM os_teams`).Scan(&teamCount); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if teamCount >= maxTeams {
		writeErr(w, http.StatusConflict, "No team slots left")
		return
	}

	reqGasque := 0
	for _, m := range req.Members {
		if m.Gasque {
			reqGasque++
		}
	}
	if reqGasque > 0 {
		rows, err := tx.Query(`SELECT members_json FROM os_teams`)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		usedGasque := 0
		for rows.Next() {
			var mj string
			if err := rows.Scan(&mj); err != nil {
				rows.Close()
				writeErr(w, http.StatusInternalServerError, err.Error())
				return
			}
			var ms []OSTeamMember
			_ = json.Unmarshal([]byte(mj), &ms)
			for _, m := range ms {
				if m.Gasque {
					usedGasque++
				}
			}
		}
		rows.Close()
		if usedGasque+reqGasque > maxGasque {
			writeErr(w, http.StatusConflict, "No gasque tickets left")
			return
		}
	}

	membersBytes, _ := json.Marshal(req.Members)
	now := time.Now().UTC().Format(time.RFC3339)
	res, err := tx.Exec(
		`INSERT INTO os_teams (name, description, is_late, members_json, registered_at) VALUES (?, ?, 0, ?, ?)`,
		req.Name, req.Description, string(membersBytes), now,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			writeErr(w, http.StatusConflict, "Team name already exists")
			return
		}
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	req.ID, _ = res.LastInsertId()
	req.RegisteredAt = now
	writeJSON(w, http.StatusCreated, map[string]any{"message": "Team registered", "team": req})
}

func (s *Server) osToggle(w http.ResponseWriter, r *http.Request) {
	var isOpen int
	err := s.db.QueryRow(`UPDATE os_settings SET is_open = 1 - is_open WHERE id=1 RETURNING is_open`).Scan(&isOpen)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"isOpen": isOpen != 0})
}

func (s *Server) osUpdateSettings(w http.ResponseWriter, r *http.Request) {
	var body struct {
		IsOpen    bool `json:"isOpen"`
		MaxTeams  int  `json:"maxTeams"`
		MaxGasque int  `json:"maxGasque"`
	}
	if err := decode(r, &body); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}
	_, err := s.db.Exec(
		`UPDATE os_settings SET is_open=?, max_teams=?, max_gasque=? WHERE id=1`,
		b2i(body.IsOpen), body.MaxTeams, body.MaxGasque,
	)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, body)
}

func (s *Server) osUpdateTeam(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req OSTeam
	if err := decode(r, &req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}
	if req.Members == nil {
		req.Members = []OSTeamMember{}
	}
	membersBytes, _ := json.Marshal(req.Members)
	res, err := s.db.Exec(
		`UPDATE os_teams SET name=?, description=?, is_late=?, members_json=? WHERE id=?`,
		req.Name, req.Description, b2i(req.IsLate), string(membersBytes), id,
	)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		writeErr(w, http.StatusNotFound, "team not found")
		return
	}
	req.ID = id
	writeJSON(w, http.StatusOK, req)
}

func (s *Server) osDeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	res, err := s.db.Exec(`DELETE FROM os_teams WHERE id=?`, id)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		writeErr(w, http.StatusNotFound, "team not found")
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"deleted": true})
}

// --- Tournaments (LoL & Counter-Strike) ---

func (s *Server) listTournamentTeams(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := s.db.Query(
			`SELECT id, game, name, players_json, created_at FROM tournament_teams WHERE game=? ORDER BY id ASC`,
			game,
		)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		teams := []TournamentTeam{}
		for rows.Next() {
			var t TournamentTeam
			var pj string
			if err := rows.Scan(&t.ID, &t.Game, &t.Name, &pj, &t.CreatedAt); err != nil {
				writeErr(w, http.StatusInternalServerError, err.Error())
				return
			}
			_ = json.Unmarshal([]byte(pj), &t.Players)
			if t.Players == nil {
				t.Players = []string{}
			}
			teams = append(teams, t)
		}
		writeJSON(w, http.StatusOK, teams)
	}
}

func (s *Server) createTournamentTeam(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t TournamentTeam
		if err := decode(r, &t); err != nil || strings.TrimSpace(t.Name) == "" {
			writeErr(w, http.StatusBadRequest, "name is required")
			return
		}
		if t.Players == nil {
			t.Players = []string{}
		}
		pj, _ := json.Marshal(t.Players)
		now := time.Now().UTC().Format(time.RFC3339)
		res, err := s.db.Exec(
			`INSERT INTO tournament_teams (game, name, players_json, created_at) VALUES (?, ?, ?, ?)`,
			game, t.Name, string(pj), now,
		)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		t.ID, _ = res.LastInsertId()
		t.Game = game
		t.CreatedAt = now
		writeJSON(w, http.StatusCreated, t)
	}
}

func (s *Server) updateTournamentTeam(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, "invalid id")
			return
		}
		var t TournamentTeam
		if err := decode(r, &t); err != nil || strings.TrimSpace(t.Name) == "" {
			writeErr(w, http.StatusBadRequest, "name is required")
			return
		}
		if t.Players == nil {
			t.Players = []string{}
		}
		pj, _ := json.Marshal(t.Players)
		res, err := s.db.Exec(
			`UPDATE tournament_teams SET name=?, players_json=? WHERE id=? AND game=?`,
			t.Name, string(pj), id, game,
		)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			writeErr(w, http.StatusNotFound, "team not found")
			return
		}
		t.ID = id
		t.Game = game
		writeJSON(w, http.StatusOK, t)
	}
}

func (s *Server) deleteTournamentTeam(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, "invalid id")
			return
		}
		res, err := s.db.Exec(`DELETE FROM tournament_teams WHERE id=? AND game=?`, id, game)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			writeErr(w, http.StatusNotFound, "team not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]bool{"deleted": true})
	}
}

func (s *Server) listMatches(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := s.db.Query(
			`SELECT id, game, round, team1, team2, winner, match_time FROM tournament_matches WHERE game=? ORDER BY id ASC`,
			game,
		)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		matches := []TournamentMatch{}
		for rows.Next() {
			var m TournamentMatch
			if err := rows.Scan(&m.ID, &m.Game, &m.Round, &m.Team1, &m.Team2, &m.Winner, &m.MatchTime); err != nil {
				writeErr(w, http.StatusInternalServerError, err.Error())
				return
			}
			matches = append(matches, m)
		}
		writeJSON(w, http.StatusOK, matches)
	}
}

func (s *Server) createMatch(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m TournamentMatch
		if err := decode(r, &m); err != nil || strings.TrimSpace(m.Round) == "" || strings.TrimSpace(m.Team1) == "" || strings.TrimSpace(m.Team2) == "" {
			writeErr(w, http.StatusBadRequest, "round, team1, and team2 are required")
			return
		}
		res, err := s.db.Exec(
			`INSERT INTO tournament_matches (game, round, team1, team2, winner, match_time) VALUES (?, ?, ?, ?, ?, ?)`,
			game, m.Round, m.Team1, m.Team2, m.Winner, m.MatchTime,
		)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		m.ID, _ = res.LastInsertId()
		m.Game = game
		writeJSON(w, http.StatusCreated, m)
	}
}

func (s *Server) updateMatch(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, "invalid id")
			return
		}
		var m TournamentMatch
		if err := decode(r, &m); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid json")
			return
		}
		res, err := s.db.Exec(
			`UPDATE tournament_matches SET round=?, team1=?, team2=?, winner=?, match_time=? WHERE id=? AND game=?`,
			m.Round, m.Team1, m.Team2, m.Winner, m.MatchTime, id, game,
		)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			writeErr(w, http.StatusNotFound, "match not found")
			return
		}
		m.ID = id
		m.Game = game
		writeJSON(w, http.StatusOK, m)
	}
}

func (s *Server) deleteMatch(game string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, "invalid id")
			return
		}
		res, err := s.db.Exec(`DELETE FROM tournament_matches WHERE id=? AND game=?`, id, game)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			writeErr(w, http.StatusNotFound, "match not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]bool{"deleted": true})
	}
}
