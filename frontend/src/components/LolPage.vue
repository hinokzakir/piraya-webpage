<template>
  <div class="lol-container">
    <h1>League Of Legends Turneringen</h1>
    <div class="warning-banner">
      <strong>VIKTIGT:</strong> Matcher kan börja tidigare eller senare än schemalagt beroende på hur matcherna går. Ha koll på Discord-servern och streamen för uppdateringar!
    </div>
    <div class="subtitle">
      PRISERNA FÖR TURNERINGEN ÄR:<br> 1:a plats: 1500kr till laget<br> 2:a plats: 500kr till laget<br> Best Highlight: 300kr till lag- eller soloplay.
    </div>
    <div class="subtitle">
      Klicka här för vår discord server och twitch!<br><br>
      <div class="button-container">
        <a href="https://discord.gg/QUQTannz" target="_blank" class="discord-btn">
          <img src="https://logos-world.net/wp-content/uploads/2020/12/Discord-Logo.png" height="110px" width="200px" alt="Discord" class="icon" />
        </a>
        <a href="https://twitch.tv/pirayagamingproviders" target="_blank" class="twitch-btn">
          <img src="https://upload.wikimedia.org/wikipedia/commons/a/a7/Twitch_BlackLogo.svg" height="80px" width="200px" alt="Twitch"  class="icon" />
        </a>
      </div>
    </div>
    <h2 class="bracket-title">FAQ</h2>
    <div class="info-boxes">
      <div class="info-box" v-for="(section, index) in sections" :key="index">
        <h2>{{ section.title }}</h2>
        <div class="box-content" v-html="section.content"></div>
      </div>
    </div>

    <!-- Tournament Bracket -->
    <div class="bracket-section">
      <h2 class="bracket-title">Turnerings Bracket</h2>
      <div class="bracket-container">
        <!-- Round of 16 -->
        <div class="bracket-round round-of-16">
          <div class="round-label">Åttondelsfinaler</div>
          <div class="match" v-for="(match, i) in bracket.roundOf16" :key="'r16-' + i">
            <div class="match-time" v-if="match.time">{{ match.time }}</div>
            <div class="team blue-team" :class="{ winner: match.winner === match.team1 }">{{ match.team1 }}</div>
            <div class="team red-team" :class="{ winner: match.winner === match.team2 }">{{ match.team2 }}</div>
          </div>
        </div>

        <!-- Quarter Finals -->
        <div class="bracket-round quarter-finals">
          <div class="round-label">Kvartsfinaler</div>
          <div class="match" v-for="(match, i) in bracket.quarterFinals" :key="'qf-' + i">
            <div class="match-time" v-if="match.time">{{ match.time }}</div>
            <div class="team blue-team" :class="{ winner: match.winner === match.team1 }">{{ match.team1 }}</div>
            <div class="team red-team" :class="{ winner: match.winner === match.team2 }">{{ match.team2 }}</div>
          </div>
        </div>

        <!-- Semi Finals -->
        <div class="bracket-round semi-finals">
          <div class="round-label">Semifinaler</div>
          <div class="match" v-for="(match, i) in bracket.semiFinals" :key="'sf-' + i">
            <div class="match-time" v-if="match.time">{{ match.time }}</div>
            <div class="team blue-team" :class="{ winner: match.winner === match.team1 }">{{ match.team1 }}</div>
            <div class="team red-team" :class="{ winner: match.winner === match.team2 }">{{ match.team2 }}</div>
          </div>
        </div>

        <!-- Final -->
        <div class="bracket-round final">
          <div class="round-label">Final (BO3)</div>
          <div class="match final-match">
            <div class="match-time" v-if="bracket.final.time">{{ bracket.final.time }}</div>
            <div class="team blue-team" :class="{ winner: bracket.final.winner === bracket.final.team1 }">{{ bracket.final.team1 }}</div>
            <div class="team red-team" :class="{ winner: bracket.final.winner === bracket.final.team2 }">{{ bracket.final.team2 }}</div>
          </div>
        </div>

        <!-- Champion -->
        <div class="bracket-round champion">
          <div class="round-label">🏆 Mästare</div>
          <div class="champion-box">
            <div class="champion-name">{{ bracket.champion }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LolPage',
  data() {
    return {
      // Tournament bracket data - easy to update with team names
      bracket: {
        // Round of 16 (8 matches)
        roundOf16: [
          { team1: 'TeamKadeem', team2: 'The Wizards of Cz', winner: 'The Wizards of Cz', time: '16:00 3/12' },
          { team1: 'Jontes Änglar', team2: 'Baron breewster', winner: 'Baron breewster', time: '16:45 3/12' },
          { team1: 'INTedgers', team2: 'Grit Gaming', winner: 'Grit Gaming', time: '17:30 3/12' },
          { team1: 'airfryerGroup', team2: 'Oskars gossar', winner: 'airfryerGroup', time: '18:15 3/12' },
          { team1: 'Data Party (Freakoff)', team2: 'svep67', winner: 'svep67', time: '19:00 3/12' },
          { team1: 'FF15', team2: 'Larkers Segjärn', winner: 'Larkers Segjärn', time: '19:45 3/12' },
          { team1: 'Poppyprovider', team2: 'SVP Jihad', winner: 'Poppyprovider', time: '20:30 3/12' },
          { team1: 'Pappas lena händer', team2: 'SkibidiHuzzlers', winner: 'Pappas lena händer', time: '21:15 3/12' }
        ],
        // Quarter Finals (4 matches)
        quarterFinals: [
          { team1: 'The Wizards of Cz', team2: 'Baron breewster', winner: 'Baron breewster', time: '17:00 4/12' },
          { team1: 'Grit Gaming', team2: 'airfryerGroup', winner: 'airfryerGroup', time: '17:45 4/12' },
          { team1: 'svep67', team2: 'Larkers Segjärn', winner: 'svep67', time: '18:30 4/12' },
          { team1: 'Poppyprovider', team2: 'Pappas lena händer', winner: 'Pappas lena händer', time: '19:15 4/12' }
        ],
        // Semi Finals (2 matches)
        semiFinals: [
          { team1: 'Baron Brewster', team2: 'airfryerGroup', winner: 'airfryerGroup', time: '20:00 4/12' },
          { team1: 'svep67', team2: 'Pappas lena händer', winner: 'svep67', time: '20:45 4/12' }
        ],
        // Final (1 match - BO3)
        final: { team1: 'airfryerGroup', team2: 'svep67', winner: null, time: '18:30 5/12' },
        // Champion
        champion: 'TBD'
      },
      sections: [
        {
          title: 'Tid och plats',
          content: `
            <p>Alla matcher spelas hemma</p>
            <p>OBS: tiderna är ungefärliga, ha koll på streamen och discord</p>
            <h3>Onsdag 3/12</h3>
            <p>16:00 - 22:00<br>Max 1 match per lag</p>
            <h3>Torsdag 4/12</h3>
            <p>17:00 - 21:30<br>1 - 2 matcher per lag som kvarstår</p>
            <h3>Fredag 5/12</h3>
            <p>18:00 - 22:00</p>
          `
        },
        {
          title: 'Streaming',
          content: `
            <p>Alla matcher streamas på vår Twitch:</p>
            <a href="https://twitch.tv/pirayagamingproviders" target="_blank" class="twitch-btn">
              <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
                <path d="M11.571 4.714h1.715v5.143H11.57zm4.715 0H18v5.143h-1.714zM6 0L1.714 4.286v15.428h5.143V24l4.286-4.286h3.428L22.286 12V0zm14.571 11.143l-3.428 3.428h-3.429l-3 3v-3H6.857V1.714h13.714z"/>
              </svg>
              twitch.tv/pirayagamingproviders
            </a>
            <p>På fredagen visas matcherna med live commentary på Kårhuset Origo.</p>
          `
        },
        {
          title: 'Viktigt: Server',
          content: `
            <p>Spelas på <strong>EU West</strong>.</p>
            <p>Ingen crossplay. Alla konton måste vara på EUW och ha minst 20 champs för Draft Pick.</p>
          `
        },
        {
          title: 'Lagkrav',
          content: `
            <p>5 spelare per lag.</p>
            <p>Minst 4 måste studera eller nyligen studerat på Umeå Universitet.</p>
          `
        },
        {
          title: 'Spelupplägg',
          content: `
            <ul>
              <li>Vanlig bracket (Single Elimination)</li>
              <li>Finalen spelas Bäst av 3 (BO3)</li>
              <li>Spelschema lottas och läggs upp före start</li>
              <li>Alla matcher streamas på Twitch</li>
            </ul>
          `
        },
        {
          title: 'CS-turneringen',
          content: `
            <p><strong>Ersätter detta CS-turneringen?</strong></p>
            <p>Nej! CS-turneringen planeras till Q1 2026.</p>
          `
        },
        {
          title: 'Mer frågor?',
          content: `
            <p>Kom förbi och snacka med oss på Tisdagar i MIT eller skicka DM till oss på Instagram.</p>
          `
        }
      ]
    };
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');

.subtitle {
  width: 100%;
  max-width: 99%;
  padding-top: 2%;
  padding-bottom: 2%;
  border-radius: 18px;
  margin-bottom: 32px;
  color: #e0e0e0;
  font-size: 1.8rem;
  text-align: center;
  font-family: 'IM Fell English SC', serif;
  background-color: #a8180c;
  background: rgba(20, 20, 20, 0.92);
  text-shadow: 0 2px 16px #000000, 1 1px 0 #000;
}

.warning-banner {
  background: #e3eb05;
  color: #000000;
  padding: 12px 20px;
  border-radius: 8px;
  font-family: 'IM Fell English SC', serif;
  margin-bottom: 24px;
  text-align: center;
  box-shadow: 0 2px 16px rgba(0,0,0,0.45);
}

.lol-container {
  min-height: 100vh;
  background: #000;
  padding: 48px 8vw;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.button-container {
  gap: 100px;
}
  

h1 {
  font-family: 'IM Fell English SC', serif;
  color: #fff;
  font-size: clamp(2.2rem, 6vw, 4.5rem);
  text-align: center;
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
  margin-top: 0;
  margin-bottom: 48px;
  letter-spacing: 0.04em;
}

.info-boxes {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 32px;
  width: 100%;
  max-width: 1400px;
}

.info-box {
  background: rgba(20, 20, 20, 0.92);
  padding: 28px 24px;
  border-radius: 18px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.45);
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  transition: transform 0.2s;
}

.info-box:hover {
  transform: translateY(-4px) scale(1.02);
  box-shadow: 0 8px 32px #a8180c55;
}

.info-box h2 {
  font-family: 'IM Fell English SC', serif;
  font-size: 1.5rem;
  color: #fff;
  margin-top: 0;
  margin-bottom: 16px;
  text-shadow: 0 2px 8px #a8180c;
  border-bottom: 2px solid #a8180c;
  padding-bottom: 8px;
}

.box-content {
  font-size: 1.05rem;
  line-height: 1.6;
}

.box-content h3 {
  font-family: 'IM Fell English SC', serif;
  font-size: 1.2rem;
  color: #e8e8e8;
  margin: 16px 0 8px 0;
  text-shadow: 0 1px 4px #a8180c88;
}

.box-content p {
  margin: 8px 0;
  color: #e0e0e0;
}

.box-content ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.box-content li {
  padding: 6px 0;
  padding-left: 20px;
  position: relative;
  color: #e0e0e0;
}

.box-content li::before {
  content: '•';
  color: #a8180c;
  font-size: 1.2rem;
  position: absolute;
  left: 0;
}

.box-content a {
  color: #ffd700;
  text-decoration: none;
  transition: color 0.2s;
}

.box-content a:hover {
  color: #ffec8b;
  text-decoration: underline;
}

.box-content .twitch-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: #9146ff;
  color: #fff;
  padding: 12px 20px;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: bold;
  text-decoration: none;
  transition: background 0.2s, transform 0.2s;
  margin: 12px 0;
}

.box-content .twitch-btn:hover {
  background: #772ce8;
  transform: scale(1.05);
  color: #fff;
  text-decoration: none;
}

.box-content .twitch-btn svg {
  flex-shrink: 0;
}

.box-content strong {
  color: #fff;
}

@media (max-width: 1100px) {
  .info-boxes {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 700px) {
  .lol-container {
    padding: 24px 4vw;
  }

  .info-boxes {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  h1 {
    font-size: clamp(1.2rem, 8vw, 2rem);
    margin-bottom: 24px;
  }

  .info-box {
    padding: 20px 16px;
  }

  .info-box h2 {
    font-size: 1.2rem;
  }

  .box-content {
    font-size: 0.95rem;
  }
}

/* Tournament Bracket Styles */
.bracket-section {
  width: 100%;
  max-width: 1400px;
  margin-top: 64px;
  padding: 32px;
  background: rgba(20, 20, 20, 0.92);
  border-radius: 18px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.45);
}

.bracket-title {
  font-family: 'IM Fell English SC', serif;
  font-size: 2rem;
  color: #fff;
  text-align: center;
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
  margin-top: 0;
  margin-bottom: 32px;
  letter-spacing: 0.04em;
}

.bracket-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
  overflow-x: auto;
  padding: 16px 0;
}

.bracket-round {
  display: flex;
  flex-direction: column;
  gap: 16px;
  min-width: 160px;
}

.round-label {
  font-family: 'IM Fell English SC', serif;
  font-size: 1rem;
  color: #a8180c;
  text-align: center;
  margin-bottom: 8px;
  text-shadow: 0 1px 4px #000;
  font-weight: bold;
}

.match {
  background: rgba(40, 40, 40, 0.9);
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #333;
  transition: transform 0.2s, box-shadow 0.2s;
}

.match-time {
  font-family: 'IM Fell English SC', serif;
  font-size: 2rem;
  color: #ffd700;
  text-align: center;
  padding: 6px 8px;
  background: rgba(0, 0, 0, 0.5);
  border-bottom: 1px solid #444;
  font-weight: bold;
}

.match:hover {
  transform: scale(1.02);
  box-shadow: 0 4px 16px #a8180c44;
}

.team {
  font-family: 'IM Fell English SC', serif;
  font-size: 0.9rem;
  color: #ccc;
  padding: 10px 12px;
  border-bottom: 1px solid #444;
  transition: background 0.2s, color 0.2s;
  position: relative;
}

.team:last-child {
  border-bottom: none;
}

/* Blue team (left/home) */
.team.blue-team {
  border-left: 4px solid #1e90ff;
  background: linear-gradient(90deg, rgba(30, 144, 255, 0.15) 0%, transparent 100%);
}

/* Red team (right/away) */
.team.red-team {
  border-left: 4px solid #dc3545;
  background: linear-gradient(90deg, rgba(220, 53, 69, 0.15) 0%, transparent 100%);
}

.team.winner {
  font-weight: bold;
  color: #ffd700;
  text-shadow: 0 0 10px #ffd700, 0 0 20px #ffd700, 0 0 30px #ffb700;
  animation: winner-glow 2s ease-in-out infinite alternate;
}

@keyframes winner-glow {
  from {
    text-shadow: 0 0 10px #ffd700, 0 0 20px #ffd700, 0 0 30px #ffb700;
  }
  to {
    text-shadow: 0 0 15px #ffd700, 0 0 30px #ffd700, 0 0 45px #ffb700;
  }
}

.team.blue-team.winner {
  border-left-color: #1e90ff;
}

.team.red-team.winner {
  border-left-color: #dc3545;
}

/* Round of 16 spacing */
.round-of-16 .match {
  margin-bottom: 0;
}

/* Quarter Finals spacing */
.quarter-finals {
  justify-content: space-around;
}

.quarter-finals .match {
  margin-bottom: 48px;
}

.quarter-finals .match:last-child {
  margin-bottom: 0;
}

/* Semi Finals spacing */
.semi-finals {
  justify-content: space-around;
}

.semi-finals .match {
  margin-bottom: 120px;
}

.semi-finals .match:last-child {
  margin-bottom: 0;
}

/* Final styling */
.final-match {
  border: 2px solid #a8180c;
  box-shadow: 0 0 20px #a8180c44;
}

.final-match .team {
  font-size: 1rem;
  padding: 14px 16px;
}

/* Champion styling */
.champion {
  align-items: center;
}

.champion-box {
  background: linear-gradient(135deg, #ffd700 0%, #ffb700 50%, #ffd700 100%);
  border-radius: 12px;
  padding: 20px 24px;
  text-align: center;
  box-shadow: 0 4px 24px rgba(255, 215, 0, 0.4);
  border: 2px solid #fff;
}

.champion-name {
  font-family: 'IM Fell English SC', serif;
  font-size: 1.2rem;
  color: #000;
  font-weight: bold;
  text-shadow: 0 1px 2px rgba(255,255,255,0.5);
}

/* Mobile responsive */
@media (max-width: 1100px) {
  .bracket-container {
    flex-wrap: wrap;
    justify-content: center;
  }

  .bracket-round {
    margin-bottom: 32px;
  }

  .quarter-finals .match,
  .semi-finals .match {
    margin-bottom: 16px;
  }
}

@media (max-width: 700px) {
  .bracket-section {
    padding: 16px;
    margin-top: 32px;
  }

  .bracket-title {
    font-size: 1.4rem;
    margin-bottom: 24px;
  }

  .bracket-container {
    flex-direction: column;
    align-items: stretch;
  }

  .bracket-round {
    min-width: 100%;
    margin-bottom: 24px;
  }

  .match {
    max-width: 100%;
  }

  .team {
    font-size: 0.85rem;
    padding: 8px 10px;
  }

  .quarter-finals .match,
  .semi-finals .match {
    margin-bottom: 8px;
  }
}
</style>
