<template>
  <div class="admin-page">
    <h1>Piraya Admin</h1>

    <!-- Login Form -->
    <div v-if="!token" class="login-container">
      <form @submit.prevent="login">
        <div class="form-group">
          <label>Användarnamn</label>
          <input v-model="username" type="text" required />
        </div>
        <div class="form-group">
          <label>Lösenord</label>
          <input v-model="password" type="password" required />
        </div>
        <button type="submit">Logga in</button>
      </form>
    </div>

    <!-- Dashboard -->
    <div v-else class="dashboard">
      <div class="controls">
        <h2>Status: <span :class="{'open': isOpen, 'closed': !isOpen}">{{ isOpen ? 'ÖPPEN' : 'STÄNGD' }}</span></h2>
        <button @click="toggleStatus" class="toggle-btn">
          {{ isOpen ? 'Stäng Anmälan' : 'Öppna Anmälan' }}
        </button>
      </div>

      <div class="teams-list">
        <h2>Anmälda Lag ({{ teams.length }})</h2>
        <table>
          <thead>
            <tr>
              <th>Lag</th>
              <th>Medlemmar</th>
              <th>Reg. Datum</th>
              <th>Åtgärder</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="team in teams" :key="team.id">
              <td>
                <input v-if="editingTeamId === team.id" v-model="editForm.name" class="edit-input" />
                <span v-else>{{ team.name }}</span>
              </td>
              <td>
                <span v-if="editingTeamId !== team.id">{{ team.members.length }}</span>
                <button v-if="editingTeamId !== team.id" @click="viewMembers(team)" class="small-btn">Visa</button>
              </td>
              <td>{{ new Date(team.registeredAt).toLocaleString() }}</td>
              <td>
                <button v-if="editingTeamId !== team.id" @click="startEdit(team)" class="action-btn edit">Ändra</button>
                <button v-if="editingTeamId !== team.id" @click="deleteTeam(team.id)" class="action-btn delete">Ta bort</button>
                
                <button v-if="editingTeamId === team.id" @click="saveEdit(team.id)" class="action-btn save">Spara</button>
                <button v-if="editingTeamId === team.id" @click="cancelEdit" class="action-btn cancel">Avbryt</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Members Modal (Simple View) -->
      <div v-if="selectedTeam" class="modal-overlay" @click.self="selectedTeam = null">
        <div class="modal">
          <h3>Medlemmar för {{ selectedTeam.name }}</h3>
          <ul>
            <li v-for="(m, i) in selectedTeam.members" :key="i">
              <strong>{{ m.name }}</strong> ({{ m.nickname }}) - {{ m.role }} - {{ m.drink }} [{{ m.personnumber }}]
            </li>
          </ul>
          <button @click="selectedTeam = null">Stäng</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminPage',
  data() {
    return {
      username: '',
      password: '',
      token: null,
      isOpen: false,
      teams: [],
      editingTeamId: null,
      editForm: {},
      selectedTeam: null
    };
  },
  methods: {
    async login() {
      try {
        const ip = process.env.VUE_APP_OTYMPLIGASPELEN_IP || 'localhost';
        const baseUrl = `http://${ip}`;

        const res = await fetch(`${baseUrl}/api/admin/login`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ username: this.username, password: this.password })
        });
        const data = await res.json();
        if (data.success) {
          this.token = data.token;
          this.fetchData();
        } else {
          alert('Fel inloggningsuppgifter');
        }
      } catch (e) {
        console.error(e);
      }
    },
    async fetchData() {
      // i dont know how to do this
      await this.refreshStatus();
    },
    async refreshStatus() {
        const ip = process.env.VUE_APP_OTYMPLIGASPELEN_IP || 'localhost';
        const port = process.env.VUE_APP_OTYMPLIGASPELEN_PORT || '3000';
        const baseUrl = `http://${ip}:${port}`;
        
        // Just calling status endpoint
        const res = await fetch(`${baseUrl}/api/status`);
        const data = await res.json();
        this.isOpen = data.isOpen;
        this.teams = data.teams || [];
    },
    async toggleStatus() {
      const ip = process.env.VUE_APP_OTYMPLIGASPELEN_IP || 'localhost';
      const port = process.env.VUE_APP_OTYMPLIGASPELEN_PORT || '3000';
      const baseUrl = `http://${ip}:${port}`;

      const res = await fetch(`${baseUrl}/api/admin/toggle`, { method: 'POST' });
      const data = await res.json();
      this.isOpen = data.isOpen;
      // If we just opened, fetch teams
      if (this.isOpen) this.refreshStatus();
    },
    deleteTeam(id) {
        if(!confirm('Är du säker?')) return;
        const ip = process.env.VUE_APP_OTYMPLIGASPELEN_IP || 'localhost';
        const port = process.env.VUE_APP_OTYMPLIGASPELEN_PORT || '3000';
        const baseUrl = `http://${ip}:${port}`;

        fetch(`${baseUrl}/api/admin/teams/${id}`, { method: 'DELETE' })
            .then(() => this.refreshStatus());
    },
    startEdit(team) {
        this.editingTeamId = team.id;
        this.editForm = JSON.parse(JSON.stringify(team));
    },
    cancelEdit() {
        this.editingTeamId = null;
        this.editForm = {};
    },
    saveEdit(id) {
        const ip = process.env.VUE_APP_OTYMPLIGASPELEN_IP || 'localhost';
        const port = process.env.VUE_APP_OTYMPLIGASPELEN_PORT || '3000';
        const baseUrl = `http://${ip}:${port}`;

        fetch(`${baseUrl}/api/admin/teams/${id}`, { 
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(this.editForm)
        }).then(() => {
            this.editingTeamId = null;
            this.refreshStatus();
        });
    },
    viewMembers(team) {
        this.selectedTeam = team;
    }
  }
};
</script>

<style scoped>
.admin-page {
  padding: 100px 20px;
  color: white;
  min-height: 100vh;
  text-align: center;
  background: #111;
  font-family: sans-serif;
}

.login-container {
  max-width: 300px;
  margin: 0 auto;
  background: #222;
  padding: 20px;
  border-radius: 8px;
}

.form-group {
    margin-bottom: 15px;
    text-align: left;
}
.form-group label {
    display: block;
    margin-bottom: 5px;
}
input {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
}

button {
    padding: 10px 20px;
    cursor: pointer;
}

.dashboard {
    max-width: 800px;
    margin: 0 auto;
}

.controls {
    background: #222;
    padding: 20px;
    border-radius: 8px;
    margin-bottom: 30px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.open { color: #4caf50; }
.closed { color: #f44336; }

.toggle-btn {
    background: #a8180c;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 1.1rem;
}

table {
    width: 100%;
    border-collapse: collapse;
    background: #1a1a1a;
}

th, td {
    padding: 12px;
    text-align: left;
    border-bottom: 1px solid #333;
}

th { background: #333; }

.action-btn {
    margin-right: 5px;
    padding: 5px 10px;
    font-size: 0.8rem;
}

.delete { background: #f44336; color: white; border: none; }
.edit { background: #2196f3; color: white; border: none; }
.save { background: #4caf50; color: white; border: none; }
.cancel { background: #ccc; border: none; }

.small-btn {
    font-size: 0.8rem;
    padding: 2px 5px;
}

.modal-overlay {
    position: fixed;
    top: 0; left: 0; right: 0; bottom: 0;
    background: rgba(0,0,0,0.8);
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal {
    background: #222;
    padding: 30px;
    border-radius: 8px;
    max-width: 600px;
    max-height: 80vh;
    overflow-y: auto;
    text-align: left;
}
.modal ul { padding-left: 20px; }
.modal li { margin-bottom: 8px; }

.edit-input {
    background: #333;
    border: 1px solid #555;
    color: white;
}
</style>
