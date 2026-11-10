<template>
  <div class="cs-container">
    <h1>Counter-Strike Turneringen</h1>

    <div class="section">
      <h2>Anmälda Lag</h2>
      <p v-if="teams.length === 0">Inga lag inlagda ännu.</p>
      <div class="grid">
        <div v-for="team in teams" :key="team.id" class="card">
          <h3>{{ team.name }}</h3>
          <p v-if="team.players && team.players.length">{{ team.players.join(', ') }}</p>
        </div>
      </div>
    </div>

    <div class="section">
      <h2>Matcher</h2>
      <p v-if="matches.length === 0">Inga matcher schemalagda ännu.</p>
      <div class="grid">
        <div v-for="match in matches" :key="match.id" class="card">
          <h3>{{ match.round }}</h3>
          <p v-if="match.matchTime" class="meta">{{ match.matchTime }}</p>
          <p>{{ match.team1 }} vs {{ match.team2 }}</p>
          <p v-if="match.winner" class="winner">Vinnare: {{ match.winner }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CsPage',
  data() {
    return {
      teams: [],
      matches: []
    };
  },
  async mounted() {
    try {
      const [teamsRes, matchesRes] = await Promise.all([
        fetch('/api/cs/teams'),
        fetch('/api/cs/matches')
      ]);
      if (teamsRes.ok) this.teams = await teamsRes.json();
      if (matchesRes.ok) this.matches = await matchesRes.json();
    } catch (e) {
      console.error('Kunde inte hämta CS-data:', e);
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');

.cs-container {
  min-height: 100vh;
  background: #000;
  color: #fff;
  padding: 120px 8vw 48px;
  font-family: 'IM Fell English SC', serif;
  text-align: center;
}
h1 {
  font-size: clamp(2rem, 6vw, 4rem);
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
  margin-bottom: 40px;
}
h2 {
  font-size: 2rem;
  color: #a8180c;
  margin-bottom: 20px;
}
.section {
  max-width: 1100px;
  margin: 0 auto 48px;
}
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 20px;
}
.card {
  background: rgba(20, 20, 20, 0.92);
  border-radius: 14px;
  padding: 20px;
  border-left: 4px solid #a8180c;
}
.card h3 {
  margin-top: 0;
}
.meta {
  color: #ccc;
}
.winner {
  color: #ffd700;
  font-weight: bold;
}
</style>
