<template>
  <div class="admin-page">
    <h1>Piraya Admin</h1>

    <div v-if="!token" class="box login-box">
      <form @submit.prevent="login">
        <div class="form-group"><label>Användarnamn</label><input v-model="username" required /></div>
        <div class="form-group"><label>Lösenord</label><input v-model="password" type="password" required /></div>
        <button type="submit">Logga in</button>
      </form>
    </div>

    <div v-else class="dashboard">
      <div class="tabs">
        <button v-for="(label, key) in tabs" :key="key" :class="{ active: tab === key }" @click="setTab(key)">{{ label }}</button>
        <button class="logout-btn" @click="logout">Logga ut</button>
      </div>

      <p v-if="errorMsg" class="error-banner">{{ errorMsg }}</p>

      <div v-if="tab === 'os'">
        <div class="box controls">
          <h2>Anmälan: <span :class="isOpen ? 'open' : 'closed'">{{ isOpen ? 'ÖPPEN' : 'STÄNGD' }}</span></h2>
          <button @click="toggleStatus" class="toggle-btn">{{ isOpen ? 'Stäng Anmälan' : 'Öppna Anmälan' }}</button>
        </div>
        <form class="box inline-form" @submit.prevent="saveSettings">
          <label>Max lag: <input v-model.number="settings.maxTeams" type="number" min="0" required /></label>
          <label>Max gasque: <input v-model.number="settings.maxGasque" type="number" min="0" required /></label>
          <button type="submit" class="btn save">Spara inställningar</button>
        </form>
      </div>

      <div v-if="tab === 'os' || tab === 'lol' || tab === 'cs'" class="box controls">
        <h2>Sida i navbar: <span :class="pages[tab] ? 'open' : 'closed'">{{ pages[tab] ? 'PÅ' : 'AV' }}</span></h2>
        <button @click="togglePage(tab)" class="toggle-btn">{{ pages[tab] ? 'Slå av' : 'Slå på' }}</button>
      </div>

      <div v-for="sec in sections" :key="sec.id" class="section">
        <h2>{{ sec.title }} ({{ sec.items.length }})</h2>

        <form v-if="sec.create" class="box inline-form" @submit.prevent="createItem(sec)">
          <template v-for="f in sec.fields" :key="f.key">
            <select v-if="f.options" v-model="sec.draft[f.key]" :required="f.req">
              <option v-if="f.req" disabled value="">Välj {{ f.label }}</option>
              <option v-for="o in f.options" :key="o" :value="o">{{ o || ('Ingen ' + f.label) }}</option>
            </select>
            <input v-else v-model="sec.draft[f.key]" :placeholder="f.label + (f.req ? ' *' : '')" :required="f.req" />
          </template>
          <button type="submit" class="btn save">Lägg till</button>
          <button v-if="sec.isMatch" type="button" class="btn edit" @click="generateGroupMatches(sec)">Skapa gruppmatcher</button>
        </form>

        <table>
          <thead>
            <tr>
              <th v-for="f in sec.fields" :key="f.key">{{ f.label }}</th>
              <th v-if="sec.hasMembers">Medlemmar</th>
              <th v-if="sec.hasMembers">Reg. Datum</th>
              <th>Åtgärder</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in sec.items" :key="item.id">
              <td v-for="f in sec.fields" :key="f.key">
                <template v-if="editKey === sec.id + ':' + item.id">
                  <select v-if="f.options" v-model="editForm[f.key]">
                    <option v-for="o in f.options" :key="o" :value="o">{{ o || '-' }}</option>
                  </select>
                  <input v-else v-model="editForm[f.key]" :type="f.type || 'text'" :class="{ 'checkbox-input': f.type === 'checkbox' }" />
                </template>
                <span v-else-if="sec.isMatch && f.key === 'winner'" class="winner-btns">
                  <button type="button" class="btn" :class="item.winner === item.team1 ? 'save' : 'pick'" @click="setWinner(sec, item, item.team1)">{{ item.team1 }}</button>
                  <button type="button" class="btn" :class="item.winner === item.team2 ? 'save' : 'pick'" @click="setWinner(sec, item, item.team2)">{{ item.team2 }}</button>
                  <button v-if="item.winner" type="button" class="btn cancel" @click="setWinner(sec, item, '')">×</button>
                </span>
                <span v-else>{{ formatVal(item[f.key], f) }}</span>
              </td>
              <td v-if="sec.hasMembers">
                {{ (item.members || []).length }}
                <button @click="selectedTeam = item" class="small-btn">Visa</button>
              </td>
              <td v-if="sec.hasMembers">{{ item.registeredAt ? new Date(item.registeredAt).toLocaleString() : '' }}</td>
              <td>
                <span v-if="editKey !== sec.id + ':' + item.id">
                  <button @click="startEdit(sec, item)" class="btn edit">Ändra</button>
                  <button @click="deleteItem(sec, item.id)" class="btn delete">Ta bort</button>
                </span>
                <span v-else>
                  <button @click="saveEdit(sec, item.id)" class="btn save">Spara</button>
                  <button @click="editKey = null" class="btn cancel">Avbryt</button>
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="selectedTeam" class="modal-overlay" @click.self="selectedTeam = null">
        <div class="box modal">
          <h3>Medlemmar för {{ selectedTeam.name }}</h3>
          <ul>
            <li v-for="(m, i) in selectedTeam.members" :key="i">
              <strong>{{ m.name }}</strong> ({{ m.nickname }}) - {{ m.role }} - {{ m.drink }}
              <span v-if="m.gasque"> - Gasque</span>
              <span v-if="m.allergies"> - Allergi: {{ m.allergies }}</span>
              [{{ m.personnumber }}]
            </li>
          </ul>
          <button @click="selectedTeam = null">Stäng</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const GAME_SECTIONS = (game, label) => [
  {
    id: game + '-teams',
    title: `${label} – Lag`,
    url: `/api/${game}/teams`,
    create: true,
    fields: [
      { key: 'name', label: 'Lagnamn', req: true },
      { key: 'groupName', label: 'Grupp', options: ['', 'A', 'B', 'C', 'D'] },
      { key: 'players', label: 'Spelare (kommaseparerade)', isList: true }
    ]
  },
  {
    id: game + '-matches',
    title: `${label} – Matcher`,
    url: `/api/${game}/matches`,
    create: true,
    isMatch: true,
    fields: [
      { key: 'round', label: 'Runda', req: true, options: ['Grupp A', 'Grupp B', 'Grupp C', 'Grupp D', 'Kvartsfinal', 'Semifinal', 'Final'] },
      { key: 'team1', label: 'Lag 1', req: true },
      { key: 'team2', label: 'Lag 2', req: true },
      { key: 'winner', label: 'Vinnare' },
      { key: 'matchTime', label: 'Tid' }
    ]
  }
];

export default {
  name: 'AdminPage',
  data() {
    return {
      username: '',
      password: '',
      token: sessionStorage.getItem('piraya_admin_token') || null,
      tabs: { os: 'Otympliga Spelen', members: 'Medlemmar', lol: 'LoL', cs: 'Counter-Strike' },
      tab: 'os',
      errorMsg: '',
      isOpen: false,
      pages: { os: false, lol: false, cs: false },
      settings: { maxTeams: 20, maxGasque: 120 },
      sections: [],
      editKey: null,
      editForm: {},
      selectedTeam: null
    };
  },
  mounted() {
    if (this.token) this.loadTab();
  },
  methods: {
    async api(path, opts = {}) {
      this.errorMsg = '';
      const headers = { 'Content-Type': 'application/json', Authorization: `Bearer ${this.token}` };
      const res = await fetch(path, { ...opts, headers });
      const data = await res.json().catch(() => ({}));
      if (!res.ok) {
        this.errorMsg = data.error || `Fel (${res.status})`;
        throw new Error(this.errorMsg);
      }
      return data;
    },
    async login() {
      try {
        const data = await this.api('/api/admin/login', {
          method: 'POST',
          body: JSON.stringify({ username: this.username, password: this.password })
        });
        this.token = data.token;
        sessionStorage.setItem('piraya_admin_token', data.token);
        this.loadTab();
      } catch (e) {
        alert(e.message || 'Fel inloggningsuppgifter');
      }
    },
    logout() {
      this.token = null;
      sessionStorage.removeItem('piraya_admin_token');
    },
    setTab(key) {
      this.tab = key;
      this.editKey = null;
      this.loadTab();
    },
    async loadTab() {
      this.pages = await this.api('/api/pages');
      if (this.tab === 'os') {
        const data = await this.api('/api/status');
        this.isOpen = data.isOpen;
        this.settings = {
          maxTeams: data.maxTeams !== undefined ? data.maxTeams : 20,
          maxGasque: data.maxGasque !== undefined ? data.maxGasque : 120
        };
        this.sections = [{
          id: 'os-teams',
          title: 'Anmälda Lag',
          url: '/api/admin/teams',
          hasMembers: true,
          items: data.teams || [],
          fields: [
            { key: 'name', label: 'Lag', req: true },
            { key: 'description', label: 'Beskrivning' },
            { key: 'isLate', label: 'Sen', type: 'checkbox' }
          ]
        }];
        return;
      }

      const defs = this.tab === 'members'
        ? [{
            id: 'members',
            title: 'Medlemmar',
            url: '/api/members',
            create: true,
            fields: [
              { key: 'Smeknamn', label: 'Smeknamn', req: true },
              { key: 'Namn', label: 'Namn', req: true },
              { key: 'Roll', label: 'Roll' },
              { key: 'quote', label: 'Citat' },
              { key: 'bild', label: 'Bild' },
              { key: 'telefonnummer', label: 'Telefon' }
            ]
          }]
        : GAME_SECTIONS(this.tab, this.tab === 'lol' ? 'League of Legends' : 'Counter-Strike');

      this.sections = await Promise.all(defs.map(async s => {
        const draft = {};
        s.fields.forEach(f => { if (f.options) draft[f.key] = ''; });
        return { ...s, draft, items: await this.api(s.url) };
      }));
    },
    formatVal(val, f) {
      if (f.type === 'checkbox') return val ? 'Ja' : 'Nej';
      if (f.isList) return (val || []).join(', ');
      return val;
    },
    toPayload(sec, raw) {
      const body = JSON.parse(JSON.stringify(raw));
      sec.fields.forEach(f => {
        if (f.isList && typeof body[f.key] === 'string') {
          body[f.key] = body[f.key].split(',').map(s => s.trim()).filter(Boolean);
        }
      });
      return body;
    },
    async togglePage(page) {
      this.pages = await this.api(`/api/admin/pages/${page}/toggle`, { method: 'POST' });
      window.dispatchEvent(new Event('pages-updated'));
    },
    async toggleStatus() {
      const data = await this.api('/api/admin/toggle', { method: 'POST' });
      this.isOpen = data.isOpen;
      await this.loadTab();
    },
    async saveSettings() {
      await this.api('/api/admin/settings', {
        method: 'PUT',
        body: JSON.stringify({ isOpen: this.isOpen, maxTeams: Number(this.settings.maxTeams), maxGasque: Number(this.settings.maxGasque) })
      });
      await this.loadTab();
    },
    async createItem(sec) {
      await this.api(sec.url, { method: 'POST', body: JSON.stringify(this.toPayload(sec, sec.draft)) });
      await this.loadTab();
    },
    async setWinner(sec, match, winner) {
      await this.api(`${sec.url}/${match.id}`, { method: 'PUT', body: JSON.stringify({ ...match, winner }) });
      await this.loadTab();
    },
    async generateGroupMatches(sec) {
      const teams = (this.sections[0] && this.sections[0].items) || [];
      const existing = new Set(sec.items.map(m => `${m.round}:${[m.team1, m.team2].sort().join('|')}`));
      for (const g of ['A', 'B', 'C', 'D']) {
        const groupTeams = teams.filter(t => (t.groupName || '').toUpperCase() === g);
        const round = `Grupp ${g}`;
        for (let i = 0; i < groupTeams.length; i++) {
          for (let j = i + 1; j < groupTeams.length; j++) {
            const key = `${round}:${[groupTeams[i].name, groupTeams[j].name].sort().join('|')}`;
            if (!existing.has(key)) {
              await this.api(sec.url, {
                method: 'POST',
                body: JSON.stringify({ round, team1: groupTeams[i].name, team2: groupTeams[j].name, winner: '', matchTime: '' })
              });
            }
          }
        }
      }
      await this.loadTab();
    },
    startEdit(sec, item) {
      this.editKey = `${sec.id}:${item.id}`;
      const copy = JSON.parse(JSON.stringify(item));
      sec.fields.forEach(f => {
        if (f.isList) copy[f.key] = (copy[f.key] || []).join(', ');
      });
      this.editForm = copy;
    },
    async saveEdit(sec, id) {
      await this.api(`${sec.url}/${id}`, { method: 'PUT', body: JSON.stringify(this.toPayload(sec, this.editForm)) });
      this.editKey = null;
      await this.loadTab();
    },
    async deleteItem(sec, id) {
      if (!confirm('Är du säker?')) return;
      await this.api(`${sec.url}/${id}`, { method: 'DELETE' });
      await this.loadTab();
    }
  }
};
</script>

<style scoped>
.admin-page { padding: 100px 20px; color: white; min-height: 100vh; text-align: center; background: #111; font-family: sans-serif; }
.box { background: #222; padding: 20px; border-radius: 8px; }
.login-box { max-width: 300px; margin: 0 auto; }
.dashboard { max-width: 1100px; margin: 0 auto; }
.form-group { margin-bottom: 15px; text-align: left; }
.form-group label { display: block; margin-bottom: 5px; }
input, select { width: 100%; padding: 8px; box-sizing: border-box; background: #333; border: 1px solid #555; color: white; border-radius: 4px; }
.checkbox-input { width: auto; }
button { padding: 10px 20px; cursor: pointer; }
.tabs { display: flex; flex-wrap: wrap; gap: 8px; margin-bottom: 24px; justify-content: center; }
.tabs button { background: #222; color: #ccc; border: 1px solid #444; border-radius: 4px; }
.tabs button.active { background: #a8180c; color: white; border-color: #a8180c; }
.tabs .logout-btn { margin-left: auto; background: #333; color: white; }
.error-banner { background: #f44336; padding: 10px; border-radius: 4px; margin-bottom: 16px; }
.controls { margin-bottom: 20px; display: flex; justify-content: space-between; align-items: center; }
.inline-form { margin-bottom: 24px; display: flex; flex-wrap: wrap; gap: 10px; align-items: center; }
.inline-form input, .inline-form select { width: auto; flex: 1 1 160px; }
.winner-btns { display: flex; flex-wrap: wrap; gap: 4px; }
.pick { background: #333; }
.inline-form label { display: flex; align-items: center; gap: 8px; }
.inline-form label input { width: 90px; }
.open { color: #4caf50; }
.closed { color: #f44336; }
.toggle-btn { background: #a8180c; color: white; border: none; border-radius: 4px; font-size: 1.1rem; }
.section { margin-bottom: 36px; }
table { width: 100%; border-collapse: collapse; background: #1a1a1a; }
th, td { padding: 12px; text-align: left; border-bottom: 1px solid #333; }
th { background: #333; }
.btn { margin-right: 5px; padding: 5px 10px; font-size: 0.85rem; border: none; color: white; }
.delete { background: #f44336; }
.edit { background: #2196f3; }
.save { background: #4caf50; }
.cancel { background: #ccc; color: #111; }
.small-btn { font-size: 0.8rem; padding: 2px 6px; margin-left: 6px; }
.modal-overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.8); display: flex; justify-content: center; align-items: center; }
.modal { padding: 30px; max-width: 600px; max-height: 80vh; overflow-y: auto; text-align: left; }
.modal ul { padding-left: 20px; }
.modal li { margin-bottom: 8px; }
</style>
