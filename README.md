# Piraya Party Providers

A website for the student organization Piraya Party Providers @ Umeå University.  
Live at: [https://pirayapartyproviders.se](https://pirayapartyproviders.se)

### Tech Stack

- **Frontend:** Vue.js (served via Nginx reverse proxy)
- **Backend:** Go REST API
- **Database:** SQLite (persisted under `./data/`)
- **Deployment:** Docker Compose on Google Cloud VM

### Local Development

1. **Clone this repo**
Clone this repo and cd into the folder in your terminal.

2. **Configure environment:**
   ```bash
   cp .env.example .env
   ```
   *(Set your desired admin password in `.env`)*

3. **Run with Docker:**
   ```bash
   docker compose up --build
   ```