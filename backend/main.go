package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func main() {
	port := env("PORT", "8080")
	dbPath := env("DB_PATH", "./data/piraya.db")

	db, err := initDB(dbPath)
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}
	defer db.Close()

	s := &Server{
		db:        db,
		adminUser: env("ADMIN_USERNAME", "admin"),
		adminPass: env("ADMIN_PASSWORD", "piraya123"),
	}

	mux := http.NewServeMux()

	// Page visibility
	mux.HandleFunc("GET /api/pages", s.getPages)
	mux.HandleFunc("POST /api/admin/pages/{page}/toggle", s.togglePage)

	// Health
	mux.HandleFunc("GET /api/health", s.health)

	// Members
	mux.HandleFunc("GET /api/members", s.listMembers)
	mux.HandleFunc("POST /api/members", s.createMember)
	mux.HandleFunc("PUT /api/members/{id}", s.updateMember)
	mux.HandleFunc("DELETE /api/members/{id}", s.deleteMember)

	// Otympliga Spelen (OS)
	mux.HandleFunc("GET /api/status", s.osStatus)
	mux.HandleFunc("POST /api/register", s.osRegister)
	mux.HandleFunc("POST /api/admin/login", s.adminLogin)
	mux.HandleFunc("POST /api/admin/toggle", s.osToggle)
	mux.HandleFunc("PUT /api/admin/settings", s.osUpdateSettings)
	mux.HandleFunc("PUT /api/admin/teams/{id}", s.osUpdateTeam)
	mux.HandleFunc("DELETE /api/admin/teams/{id}", s.osDeleteTeam)

	// League of Legends (LoL)
	mux.HandleFunc("GET /api/lol/teams", s.listTournamentTeams("lol"))
	mux.HandleFunc("POST /api/lol/teams", s.createTournamentTeam("lol"))
	mux.HandleFunc("PUT /api/lol/teams/{id}", s.updateTournamentTeam("lol"))
	mux.HandleFunc("DELETE /api/lol/teams/{id}", s.deleteTournamentTeam("lol"))
	mux.HandleFunc("GET /api/lol/matches", s.listMatches("lol"))
	mux.HandleFunc("POST /api/lol/matches", s.createMatch("lol"))
	mux.HandleFunc("PUT /api/lol/matches/{id}", s.updateMatch("lol"))
	mux.HandleFunc("DELETE /api/lol/matches/{id}", s.deleteMatch("lol"))

	// Counter-Strike (CS)
	mux.HandleFunc("GET /api/cs/teams", s.listTournamentTeams("cs"))
	mux.HandleFunc("POST /api/cs/teams", s.createTournamentTeam("cs"))
	mux.HandleFunc("PUT /api/cs/teams/{id}", s.updateTournamentTeam("cs"))
	mux.HandleFunc("DELETE /api/cs/teams/{id}", s.deleteTournamentTeam("cs"))
	mux.HandleFunc("GET /api/cs/matches", s.listMatches("cs"))
	mux.HandleFunc("POST /api/cs/matches", s.createMatch("cs"))
	mux.HandleFunc("PUT /api/cs/matches/{id}", s.updateMatch("cs"))
	mux.HandleFunc("DELETE /api/cs/matches/{id}", s.deleteMatch("cs"))

	log.Printf("Go SQLite server listening on :%s (db=%s)", port, dbPath)
	log.Fatal(http.ListenAndServe(":"+port, cors(mux)))
}

func initDB(dbPath string) (*sql.DB, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)

	schema := `
	PRAGMA journal_mode=WAL;
	PRAGMA foreign_keys=ON;
	PRAGMA busy_timeout=5000;

	CREATE TABLE IF NOT EXISTS members (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		smeknamn TEXT NOT NULL,
		namn TEXT NOT NULL,
		roll TEXT NOT NULL DEFAULT '',
		quote TEXT NOT NULL DEFAULT '',
		bild TEXT NOT NULL DEFAULT '',
		telefonnummer TEXT NOT NULL DEFAULT ''
	);

	CREATE TABLE IF NOT EXISTS os_settings (
		id INTEGER PRIMARY KEY CHECK (id = 1),
		is_open INTEGER NOT NULL DEFAULT 0,
		max_teams INTEGER NOT NULL DEFAULT 20,
		max_gasque INTEGER NOT NULL DEFAULT 120
	);
	INSERT OR IGNORE INTO os_settings (id, is_open, max_teams, max_gasque) VALUES (1, 0, 20, 120);

	CREATE TABLE IF NOT EXISTS os_teams (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL DEFAULT '',
		is_late INTEGER NOT NULL DEFAULT 0,
		members_json TEXT NOT NULL DEFAULT '[]',
		registered_at TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS tournament_teams (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		game TEXT NOT NULL,
		name TEXT NOT NULL,
		group_name TEXT NOT NULL DEFAULT '',
		players_json TEXT NOT NULL DEFAULT '[]',
		created_at TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS page_settings (
		page TEXT PRIMARY KEY,
		visible INTEGER NOT NULL DEFAULT 0
	);
	INSERT OR IGNORE INTO page_settings (page, visible) VALUES ('os', 0), ('lol', 0), ('cs', 0);

	CREATE TABLE IF NOT EXISTS tournament_matches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		game TEXT NOT NULL,
		round TEXT NOT NULL,
		team1 TEXT NOT NULL,
		team2 TEXT NOT NULL,
		winner TEXT NOT NULL DEFAULT '',
		match_time TEXT NOT NULL DEFAULT ''
	);`

	if _, err := db.Exec(schema); err != nil {
		_ = db.Close()
		return nil, err
	}
	_, _ = db.Exec(`ALTER TABLE tournament_teams ADD COLUMN group_name TEXT NOT NULL DEFAULT ''`)
	return db, nil
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
