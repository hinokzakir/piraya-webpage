<template>
  <div class="lol-container">
    <h1>League Of Legends Turneringen</h1>
    <div class="info-boxes">
      <div class="info-box" v-for="(section, index) in sections" :key="index">
        <h2>{{ section.title }}</h2>
        <div class="box-content" v-html="section.content"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LolPage',
  data() {
    return {
      sections: [
        {
          title: 'Tid och plats',
          content: `
            <p>Alla matcher spelas hemma</p>
            <h3>Onsdag 3/12</h3>
            <p>16:00 - 23:00<br>Max 1 match per lag</p>
            <h3>Torsdag 4/12</h3>
            <p>16:00 - 23:00<br>1 - 2 matcher per lag som kvarstår</p>
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

.lol-container {
  min-height: 100vh;
  background: #000;
  padding: 48px 8vw;
  display: flex;
  flex-direction: column;
  align-items: center;
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
</style>
