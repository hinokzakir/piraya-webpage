package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	db         *sql.DB
	adminUser  string
	adminPass  string
	uploadsDir string
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
	GroupName string   `json:"groupName"`
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

// --- Health, Pages & Admin Login ---

func (s *Server) getPages(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query(`SELECT page, visible FROM page_settings`)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	pages := map[string]bool{"os": false, "lol": false, "cs": false}
	for rows.Next() {
		var page string
		var vis int
		if err := rows.Scan(&page, &vis); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		pages[page] = vis != 0
	}
	writeJSON(w, http.StatusOK, pages)
}

func (s *Server) togglePage(w http.ResponseWriter, r *http.Request) {
	page := r.PathValue("page")
	if page != "os" && page != "lol" && page != "cs" {
		writeErr(w, http.StatusBadRequest, "invalid page")
		return
	}
	if _, err := s.db.Exec(`UPDATE page_settings SET visible = 1 - visible WHERE page=?`, page); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	s.getPages(w, r)
}

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
			`SELECT id, game, name, group_name, players_json, created_at FROM tournament_teams WHERE game=? ORDER BY id ASC`,
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
			if err := rows.Scan(&t.ID, &t.Game, &t.Name, &t.GroupName, &pj, &t.CreatedAt); err != nil {
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
			`INSERT INTO tournament_teams (game, name, group_name, players_json, created_at) VALUES (?, ?, ?, ?, ?)`,
			game, t.Name, t.GroupName, string(pj), now,
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
			`UPDATE tournament_teams SET name=?, group_name=?, players_json=? WHERE id=? AND game=?`,
			t.Name, t.GroupName, string(pj), id, game,
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

// --- Batch JSON Import ---

func (s *Server) importJSON(w http.ResponseWriter, r *http.Request) {
	target := r.PathValue("target")
	if target != "os" && target != "lol" && target != "cs" && target != "members" {
		writeErr(w, http.StatusBadRequest, "invalid target (must be os, lol, cs, or members)")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil || len(strings.TrimSpace(string(bodyBytes))) == 0 {
		writeErr(w, http.StatusBadRequest, "tom eller ogiltig JSON")
		return
	}

	clearExisting := r.URL.Query().Get("clear") == "true"

	tx, err := s.db.Begin()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback()

	now := time.Now().UTC().Format(time.RFC3339)

	if target == "os" {
		if clearExisting {
			if _, err := tx.Exec(`DELETE FROM os_teams`); err != nil {
				writeErr(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		var teams []OSTeam
		if err := json.Unmarshal(bodyBytes, &teams); err != nil {
			var wrapper struct {
				Teams []OSTeam `json:"teams"`
			}
			if err2 := json.Unmarshal(bodyBytes, &wrapper); err2 != nil {
				writeErr(w, http.StatusBadRequest, "Ogiltig OS-JSON (förväntade lista av lag eller { teams: [...] })")
				return
			}
			teams = wrapper.Teams
		}

		if len(teams) == 0 {
			writeErr(w, http.StatusBadRequest, "Inga lag hittades i JSON")
			return
		}

		count := 0
		for _, t := range teams {
			if strings.TrimSpace(t.Name) == "" {
				continue
			}
			if t.Members == nil {
				t.Members = []OSTeamMember{}
			}
			mb, _ := json.Marshal(t.Members)
			regAt := t.RegisteredAt
			if regAt == "" {
				regAt = now
			}
			_, err := tx.Exec(`
				INSERT INTO os_teams (name, description, is_late, members_json, registered_at)
				VALUES (?, ?, ?, ?, ?)
				ON CONFLICT(name) DO UPDATE SET
					description = excluded.description,
					is_late = excluded.is_late,
					members_json = excluded.members_json,
					registered_at = excluded.registered_at
			`, t.Name, t.Description, b2i(t.IsLate), string(mb), regAt)
			if err != nil {
				writeErr(w, http.StatusInternalServerError, err.Error())
				return
			}
			count++
		}

		if err := tx.Commit(); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"message": fmt.Sprintf("Importerade %d lag för Otympliga Spelen", count),
			"count":   count,
		})
		return
	}

	if target == "members" {
		if clearExisting {
			if _, err := tx.Exec(`DELETE FROM members`); err != nil {
				writeErr(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		var members []Member
		if err := json.Unmarshal(bodyBytes, &members); err != nil {
			var wrapper struct {
				Members []Member `json:"members"`
			}
			if err2 := json.Unmarshal(bodyBytes, &wrapper); err2 != nil {
				writeErr(w, http.StatusBadRequest, "Ogiltig medlems-JSON (förväntade lista eller { members: [...] })")
				return
			}
			members = wrapper.Members
		}

		if len(members) == 0 {
			writeErr(w, http.StatusBadRequest, "Inga medlemmar hittades i JSON")
			return
		}

		count := 0
		for _, m := range members {
			if strings.TrimSpace(m.Smeknamn) == "" || strings.TrimSpace(m.Namn) == "" {
				continue
			}
			// Alla members ska inte ha nån profilbild vid import: bild = ""
			if clearExisting {
				_, err := tx.Exec(
					`INSERT INTO members (smeknamn, namn, roll, quote, bild, telefonnummer) VALUES (?, ?, ?, ?, '', ?)`,
					m.Smeknamn, m.Namn, m.Roll, m.Quote, m.Telefonnummer,
				)
				if err != nil {
					writeErr(w, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				var existingID int64
				err := tx.QueryRow(`SELECT id FROM members WHERE LOWER(smeknamn) = LOWER(?) LIMIT 1`, m.Smeknamn).Scan(&existingID)
				if err == nil {
					_, err = tx.Exec(
						`UPDATE members SET namn=?, roll=?, quote=?, bild='', telefonnummer=? WHERE id=?`,
						m.Namn, m.Roll, m.Quote, m.Telefonnummer, existingID,
					)
				} else {
					_, err = tx.Exec(
						`INSERT INTO members (smeknamn, namn, roll, quote, bild, telefonnummer) VALUES (?, ?, ?, ?, '', ?)`,
						m.Smeknamn, m.Namn, m.Roll, m.Quote, m.Telefonnummer,
					)
				}
				if err != nil {
					writeErr(w, http.StatusInternalServerError, err.Error())
					return
				}
			}
			count++
		}

		if err := tx.Commit(); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"message": fmt.Sprintf("Importerade %d medlemmar (utan profilbild)", count),
			"count":   count,
		})
		return
	}

	// target == "lol" || target == "cs"
	if clearExisting {
		if _, err := tx.Exec(`DELETE FROM tournament_teams WHERE game=?`, target); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		if _, err := tx.Exec(`DELETE FROM tournament_matches WHERE game=?`, target); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var teams []TournamentTeam
	var matches []TournamentMatch

	var wrapper struct {
		Teams   []TournamentTeam  `json:"teams"`
		Matches []TournamentMatch `json:"matches"`
	}
	if err := json.Unmarshal(bodyBytes, &wrapper); err == nil && (len(wrapper.Teams) > 0 || len(wrapper.Matches) > 0) {
		teams = wrapper.Teams
		matches = wrapper.Matches
	} else {
		var rawList []json.RawMessage
		if err := json.Unmarshal(bodyBytes, &rawList); err == nil && len(rawList) > 0 {
			for _, raw := range rawList {
				var m TournamentMatch
				var t TournamentTeam
				_ = json.Unmarshal(raw, &m)
				_ = json.Unmarshal(raw, &t)
				if m.Round != "" && m.Team1 != "" && m.Team2 != "" {
					matches = append(matches, m)
				} else if t.Name != "" {
					teams = append(teams, t)
				}
			}
		} else {
			writeErr(w, http.StatusBadRequest, "Ogiltig turnerings-JSON (förväntade { teams: [...], matches: [...] } eller lista)")
			return
		}
	}

	if len(teams) == 0 && len(matches) == 0 {
		writeErr(w, http.StatusBadRequest, "Inga lag eller matcher hittades i JSON")
		return
	}

	teamCount := 0
	for _, t := range teams {
		if strings.TrimSpace(t.Name) == "" {
			continue
		}
		if t.Players == nil {
			t.Players = []string{}
		}
		pj, _ := json.Marshal(t.Players)
		created := t.CreatedAt
		if created == "" {
			created = now
		}
		_, err := tx.Exec(`
			INSERT INTO tournament_teams (game, name, group_name, players_json, created_at)
			VALUES (?, ?, ?, ?, ?)
		`, target, t.Name, t.GroupName, string(pj), created)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		teamCount++
	}

	matchCount := 0
	for _, m := range matches {
		if strings.TrimSpace(m.Round) == "" || strings.TrimSpace(m.Team1) == "" || strings.TrimSpace(m.Team2) == "" {
			continue
		}
		_, err := tx.Exec(`
			INSERT INTO tournament_matches (game, round, team1, team2, winner, match_time)
			VALUES (?, ?, ?, ?, ?, ?)
		`, target, m.Round, m.Team1, m.Team2, m.Winner, m.MatchTime)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		matchCount++
	}

	if err := tx.Commit(); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"success":         true,
		"message":         fmt.Sprintf("Importerade %d lag och %d matcher för %s", teamCount, matchCount, strings.ToUpper(target)),
		"importedTeams":   teamCount,
		"importedMatches": matchCount,
	})
}

// --- Member Images & File Uploads ---

func cleanFileName(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	replacer := strings.NewReplacer(
		"å", "a", "ä", "a", "ö", "o",
		"é", "e", "è", "e",
		" ", "_",
	)
	s = replacer.Replace(s)
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func matchNickname(fileNameBase, memberNick string) bool {
	fClean := cleanFileName(fileNameBase)
	mClean := cleanFileName(memberNick)
	if fClean == "" || mClean == "" {
		return false
	}
	if fClean == mClean || strings.EqualFold(fileNameBase, memberNick) {
		return true
	}
	fNoSep := strings.ReplaceAll(strings.ReplaceAll(fClean, "_", ""), "-", "")
	mNoSep := strings.ReplaceAll(strings.ReplaceAll(mClean, "_", ""), "-", "")
	if fNoSep == mNoSep {
		return true
	}
	if strings.HasPrefix(fNoSep, mNoSep) || strings.Contains(fNoSep, mNoSep) {
		return true
	}
	return false
}

func (s *Server) serveMemberImage(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(r.PathValue("filename"))
	if filename == "." || filename == "/" || filename == "" {
		writeErr(w, http.StatusBadRequest, "ogiltigt filnamn")
		return
	}
	filePath := filepath.Join(s.uploadsDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, filePath)
}

func (s *Server) uploadMemberImages(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		writeErr(w, http.StatusBadRequest, "Kunde inte läsa uppladdade filer: "+err.Error())
		return
	}
	if r.MultipartForm == nil || len(r.MultipartForm.File) == 0 {
		writeErr(w, http.StatusBadRequest, "Inga filer mottogs")
		return
	}

	var allFiles []*multipart.FileHeader
	for _, headers := range r.MultipartForm.File {
		allFiles = append(allFiles, headers...)
	}

	if len(allFiles) == 0 {
		writeErr(w, http.StatusBadRequest, "Inga filer hittades i formuläret")
		return
	}

	rows, err := s.db.Query(`SELECT id, smeknamn, namn FROM members`)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	type memberInfo struct {
		id       int64
		smeknamn string
		namn     string
	}
	var members []memberInfo
	for rows.Next() {
		var mi memberInfo
		if err := rows.Scan(&mi.id, &mi.smeknamn, &mi.namn); err == nil {
			members = append(members, mi)
		}
	}

	if len(members) == 0 {
		writeErr(w, http.StatusBadRequest, "Det finns inga medlemmar i databasen att koppla bilderna till. Importera medlemmar först!")
		return
	}

	if err := os.MkdirAll(s.uploadsDir, 0o755); err != nil {
		writeErr(w, http.StatusInternalServerError, "Kunde inte skapa uppladdningsmapp: "+err.Error())
		return
	}

	var matched []string
	var unmatched []string

	for _, fh := range allFiles {
		ext := strings.ToLower(filepath.Ext(fh.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".gif" {
			unmatched = append(unmatched, fh.Filename+" (ej bildformat)")
			continue
		}

		baseName := strings.TrimSuffix(fh.Filename, filepath.Ext(fh.Filename))

		var targetMember *memberInfo
		for i := range members {
			if matchNickname(baseName, members[i].smeknamn) {
				targetMember = &members[i]
				break
			}
		}
		if targetMember == nil {
			for i := range members {
				if matchNickname(baseName, members[i].namn) {
					targetMember = &members[i]
					break
				}
			}
		}

		if targetMember == nil {
			unmatched = append(unmatched, fh.Filename)
			continue
		}

		file, err := fh.Open()
		if err != nil {
			unmatched = append(unmatched, fh.Filename+" (kunde inte öppnas)")
			continue
		}

		savedName := fmt.Sprintf("%s%s", cleanFileName(targetMember.smeknamn), ext)
		destPath := filepath.Join(s.uploadsDir, savedName)
		dst, err := os.Create(destPath)
		if err != nil {
			file.Close()
			unmatched = append(unmatched, fh.Filename+" (kunde inte sparas)")
			continue
		}
		_, err = io.Copy(dst, file)
		file.Close()
		dst.Close()
		if err != nil {
			unmatched = append(unmatched, fh.Filename+" (fel vid skrivning)")
			continue
		}

		imageURL := "/api/uploads/members/" + savedName
		_, err = s.db.Exec(`UPDATE members SET bild=? WHERE id=?`, imageURL, targetMember.id)
		if err != nil {
			unmatched = append(unmatched, fh.Filename+" (databasfel)")
			continue
		}

		matched = append(matched, fmt.Sprintf("%s (%s -> %s)", targetMember.smeknamn, fh.Filename, savedName))
	}

	msg := fmt.Sprintf("Kopplade %d bild(er) till medlemmar.", len(matched))
	if len(unmatched) > 0 {
		msg += fmt.Sprintf(" %d fil(er) kunde inte matchas.", len(unmatched))
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"success":   true,
		"message":   msg,
		"matched":   matched,
		"unmatched": unmatched,
	})
}
