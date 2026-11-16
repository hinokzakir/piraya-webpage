<template>
  <div class="os-page">
    <div class="content-container">
      <h1>Otympliga Spelen</h1>

      <div v-if="isOpen" class="status-bar">
        <div v-if="isFull" class="status-full">
            🔴 EVENTET ÄR FULLBOKAT
        </div>
        <div v-else class="status-open">
            🟢 Lag Platser kvar: <strong>{{ spotsLeft }}</strong>
            <span v-if="gasqueLeft !== undefined" style="margin-left: 15px;">
                (Gasque: <strong>{{ gasqueLeft }}</strong> kvar)
            </span>
        </div>
      </div>

      <div v-if="!isOpen" class="locked-message">
        <h2>När öppnar anmälan?</h2>
        <p class="deadline-text">17/2 (1337+666+420+67) / (69-67)</p>
      </div>

      <div v-else class="open-content">
        
        <div class="teams-list-section">
            <h2>Anmälda Lag</h2>
            <div v-if="teams.length === 0" class="no-teams">Inga lag anmälda än. Bli först!</div>
            <div class="teams-grid">
                <div 
                    v-for="team in teams" 
                    :key="team.id" 
                    class="team-card" 
                    :class="{ 'late-entry': team.isLate }"
                    @click="toggleTeamDetails(team.id)"
                >
                    <div class="team-header">
                        <h3>{{ team.name }}</h3>
                        <span v-if="team.isLate" class="late-badge">⚠️ RESERVPLATS</span>
                    </div>
                    <p>{{ team.members.length }} Medlemmar <span v-if="team.registeredAt" style="font-size:0.8em; opacity:0.7">({{ formatTime(team.registeredAt) }})</span></p>
                    
                    <div v-if="expandedTeamId === team.id" class="team-details">
                        <div v-for="(member, mIndex) in team.members" :key="mIndex" class="member-detail">
                            <span class="member-role">{{ member.nickname }}</span> ({{ member.role }})
                            <span v-if="member.isLate" class="member-late-mark" title="Sen anmälan">*</span>
                            <br>
                            <span class="drink-info">🍺 {{ member.drink }}</span>
                            <span class="gasque-info" v-if="member.gasque">🍽️ Gasque</span>
                        </div>
                    </div>
                </div>
            </div>
         </div>

        <div v-if="!isFull">
            <button v-if="!showForm" class="submit-btn" @click="showForm = true" style="margin-top: 0;">
                Anmäl ett lag
            </button>
        </div>
        <div v-else class="full-message">
            <p>Anmälan är stängd då maxgränsen är nådd.</p>
        </div>

        <div v-if="showForm && !isFull" class="signup-form">
          <h2>Anmälan</h2>
          <p v-if="spotsLeft < 6" class="warning-text">
            OBS: Det är få platser kvar! Om ditt lag gör att gränsen passeras kommer ni hamna på reservlista.
          </p>
          <form @submit.prevent="submitTeam">
            <div class="form-group">
              <label for="teamName">Lagnamn:</label>
              <input type="text" id="teamName" v-model="formRef.name" required placeholder="ta nåt nice" />
            </div>
            
            <div class="form-group">
              <label for="teamDescription">Lagbeskrivning:</label>
              <input type="text" id="teamDescription" v-model="formRef.description" required placeholder="skriv nåt bra info om laget" />
            </div>

            <div class="members-section">
              <h3>Lagmedlemmar</h3>
              <div v-for="(member, index) in formRef.members" :key="index" class="member-card">
                <div class="member-header">
                  <h4>Medlem {{ index + 1 }}</h4>
                  <button type="button" class="remove-btn" @click="removeMember(index)" v-if="formRef.members.length > 1">
                    &times;
                  </button>
                </div>
                
                <div class="form-group">
                  <label>Namn:</label>
                  <input type="text" v-model="member.name" required placeholder="Förnamn Efternamn" />
                </div>

                <div class="form-group">
                  <label>Personnummer (10 siffror):</label>
                  <input type="text" v-model="member.personnumber" required placeholder="ÅÅMMDDXXXX" pattern="\d{10}" title="10 siffror utan bindestreck" />
                </div>

                <div class="form-group">
                  <label>Smeknamn:</label>
                  <input type="text" v-model="member.nickname" required placeholder="ta nåt nice ba" />
                </div>

                <div class="form-group">
                  <label>Funktion i laget:</label>
                  <input type="text" v-model="member.role" required placeholder="tänk inte så länge på de" />
                </div>
                
                <div class="form-group">
                  <label>Allergier:</label>
                  <input type="text" v-model="member.allergies" placeholder="lämna tomt om inga" />
                </div>

                <div class="form-group">
                  <label>Dryckesval:</label>
                  <select v-model="member.drink" required>
                    <option disabled value="">Välj dryck</option>
                    <option value="beer">Öl</option>
                    <option value="cider">Cider</option>
                    <option value="soda">Läsk</option>
                  </select>
                </div>

                <div class="form-group checkbox-group">
                    <label class="checkbox-label">
                        <input type="checkbox" v-model="member.gasque">
                        Vill gå på Gasque?
                    </label>
                </div>
              </div>
              
              <button type="button" class="add-btn" @click="addMember">
                + Lägg till medlem
              </button>
            </div>

            <div class="form-actions">
                <button type="button" class="cancel-btn" @click="showForm = false">
                    Avbryt
                </button>
                <button type="submit" class="submit-btn" :disabled="submitting">
                    {{ submitting ? 'Skickar...' : 'Skicka Anmälan' }}
                </button>
            </div>
          </form>
        </div>


      </div>

      <div class="faq-section">
        <h2>Vanliga Frågor</h2>
        <div class="faq-grid">
            <div v-for="(item, index) in faqItems" :key="index" class="faq-item">
                <h3>{{ item.question }}</h3>
                <div class="answer" v-html="item.answer"></div>
            </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OsPage',
  data() {
    return {
      isOpen: false,
      isFull: false,      // NY: Kollar om det är fullt
      spotsLeft: 0,       // NY: Antal platser kvar
      gasqueLeft: undefined, // NY: Antal gasque-biljetter kvar
      showForm: false,
      teams: [],
      expandedTeamId: null,
      submitting: false,
      formRef: {
        name: '',
        description: '',
        members: []
      },
      pollingInterval: null,
      faqItems: [
        {
            question: 'När ska dem Otympliga spelen äga rum?',
            answer: '5-7 mars'
        },
        {
            question: 'Vart anmäler jag mig?',
            answer: 'Bror du är ju här?'
        },
        {
            question: 'När kan jag anmäla mig?',
            answer: '17/2, (1337+666+420+67) / (69-67)'
        },
        {
            question: 'Hur många lagmedlemmar får man vara?',
            answer: 'Som vanligt, 6-8 personer/lag'
        },
        {
            question: 'Pris för att delta på detta event?',
            answer: 'det håller på att beräknas, rookies håller på räkna för hand'
        },
        {
            question: 'Vad ingår i dessa biljetter?',
            answer: '<strong>OS innebär:</strong><br>Inträde på pubkväll 5e.<br>Fulsittning 5e och 6e.<br>Otympliga lekar på Campusängen 6e och 7e.<br>Ett glas blåbärsoppa<br>Några öl var.<br>Ett märke.<br>Kul.<br><br><strong>OS+gasque innebär:</strong><br>Allt ovanstående + tvårätters gasque på Origo 7e som avslutning på OS. (Begränsat antal platser)'
        }
      ]
    };
  },
  methods: {
    async fetchStatus() {
      try {
        const res = await fetch('/api/status');
        const data = await res.json();
        
        this.isOpen = data.isOpen;
        this.teams = data.teams || [];
        // Uppdatera status för platser
        this.spotsLeft = data.spotsLeft;
        this.isFull = data.isFull;
        this.gasqueLeft = data.gasqueLeft; // Kan vara undefined om servern inte stödjer det än

      } catch (err) {
        console.error('Failed to fetch status:', err);
      }
    },
    toggleTeamDetails(id) {
        if (this.expandedTeamId === id) {
            this.expandedTeamId = null;
        } else {
            this.expandedTeamId = id;
        }
    },
    formatTime(isoString) {
        if(!isoString) return '';
        const d = new Date(isoString);
        return d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
    },
    addMember() {
      this.formRef.members.push({
        name: '',
        personnumber: '',
        nickname: '',
        role: '',
        allergies: '',
        drink: '',
        gasque: false
      });
    },
    removeMember(index) {
      this.formRef.members.splice(index, 1);
    },
    async submitTeam() {
      this.submitting = true;
      try {
          // Förbered datan enligt serverns nya struktur
          // (Servern vill ha 'members' som en array av objekt)
          const payload = {
              name: this.formRef.name,
              // Vi skickar description även om servern kanske inte sparar den i nuvarande DB-schema
              description: this.formRef.description, 
              members: this.formRef.members
          };

          const res = await fetch('/api/register', {
              method: 'POST',
              headers: { 'Content-Type': 'application/json' },
              body: JSON.stringify(payload)
          });
          
          const result = await res.json();

          if (res.ok) {
              let msg = `Anmälan mottagen för laget: ${this.formRef.name}!`;
              // Om servern säger att laget blev flaggat som sent
              if (result.message && result.message.includes('Late')) {
                  msg += '\n\nOBS: Ni hamnade på reservlista/sen anmälan då platserna tog slut.';
              }
              alert(msg);
              
              this.formRef = { name: '', description: '', members: [] }; // Reset
              this.addMember(); 
              this.fetchStatus(); 
              this.showForm = false;
          } else {
              if (result.error === 'No gasque tickets left') {
                  alert('Tyvärr, gasque-biljetterna är slut för en eller flera medlemmar som valt det.');
              } else if (result.error === 'No team slots left') {
                  alert('Tyvärr, det finns inga lagplatser kvar.');
              } else {
                  alert('Fel vid anmälan: ' + (result.error || 'Okänt fel'));
              }
          }
      } catch (e) {
          console.error(e);
          alert('Kunde inte ansluta till servern.');
      } finally {
          this.submitting = false;
      }
    }
  },
  mounted() {
    this.addMember(); // Init form with 1 member
    this.fetchStatus();
    this.pollingInterval = setInterval(this.fetchStatus, 5000); // Polling oftare (5 sek) vid släpp!
  },
  beforeUnmount() {
    clearInterval(this.pollingInterval);
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');

.os-page {
  min-height: 100vh;
  background: rgba(0, 0, 0, 0); 
  padding-top: 50px;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
}

.content-container {
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
  padding-bottom: 50px;
}

h1 {
  font-size: clamp(2.5rem, 5vw, 4rem);
  margin-bottom: 20px;
  text-shadow: 0 2px 10px #a8180c;
  color: #fff;
}

/* Status Bar */
.status-bar {
    margin-bottom: 40px;
    font-size: 1.5rem;
    padding: 10px;
    background: rgba(0,0,0,0.5);
    border-radius: 8px;
    display: inline-block;
}
.status-open { color: #4caf50; }
.status-full { color: #ff4444; font-weight: bold; animation: pulse 2s infinite;}
.warning-text { color: #ff9800; font-weight: bold; margin-bottom: 20px;}

@keyframes pulse {
    0% { opacity: 1; }
    50% { opacity: 0.7; }
    100% { opacity: 1; }
}

h2 {
  font-size: 2rem;
  margin-bottom: 20px;
  color: #a8180c;
}

/* Teams List */
.teams-list-section {
    margin-bottom: 50px;
}

.teams-grid {
    display: flex;
    flex-direction: column;
    gap: 15px;
}

.team-card {
    background: rgba(255, 255, 255, 0.05);
    padding: 15px;
    border-radius: 8px;
    border-left: 4px solid #4caf50; /* Grön som default */
    cursor: pointer;
    transition: background 0.2s;
    text-align: left;
}

/* Stil för sena lag */
.team-card.late-entry {
    border-left-color: #ff4444;
    background: rgba(255, 68, 68, 0.05);
}

.team-card:hover {
    background: rgba(255, 255, 255, 0.1);
}

.team-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.late-badge {
    background: #ff4444;
    color: white;
    padding: 2px 8px;
    border-radius: 4px;
    font-size: 0.8rem;
    font-family: sans-serif;
    font-weight: bold;
}

.team-details {
    margin-top: 15px;
    padding-top: 15px;
    border-top: 1px solid #444;
}

.member-detail {
    margin-bottom: 8px;
    font-size: 1rem;
    color: #ccc;
    line-height: 1.4;
}

.member-role {
    color: #fff;
    font-weight: bold;
}

.drink-info {
    font-size: 0.85rem;
    color: #888;
    margin-left: 5px;
}

.gasque-info {
    font-size: 0.85rem;
    color: #ffd700;
    margin-left: 10px;
    font-weight: bold;
}

/* FAQ, Form, etc (Mostly unchanged styles below) */
.faq-section {
    margin-top: 80px;
    margin-bottom: 50px;
    text-align: left;
    background: rgba(20, 20, 20, 0.6);
    padding: 30px;
    border-radius: 15px;
    border: 1px solid #333;
}

.faq-section h2 {
    text-align: center;
    border-bottom: 1px solid #444;
    padding-bottom: 15px;
    margin-bottom: 25px;
    margin-top: 25px;
}

.faq-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 25px;
}

.faq-item h3 {
    color: #a8180c;
    margin-bottom: 10px;
    font-size: 1.3rem;
}

.answer {
    color: #ddd;
    line-height: 1.5;
}

.locked-message, .full-message {
  background: rgba(20, 20, 20, 0.9);
  padding: 60px;
  border-radius: 15px;
  box-shadow: 0 0 20px rgba(168, 24, 12, 0.3);
  border: 1px solid #333;
}

.deadline-text {
  font-size: 1.5rem;
  color: #ddd;
}

.signup-form {
  background: rgba(20, 20, 20, 0.92);
  padding: 30px;
  border-radius: 15px;
  box-shadow: 0 0 20px rgba(255, 255, 255, 0.1);
  text-align: left;
  margin-top: 20px;
}

.signup-form h2 {
    text-align: center;
    margin-bottom: 30px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-size: 1.2rem;
  color: #ddd;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #444;
  background: #111;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  font-size: 1.1rem;
  box-sizing: border-box; 
}

.form-group input:focus,
.form-group select:focus {
  outline: none;
  border-color: #a8180c;
  box-shadow: 0 0 8px rgba(168, 24, 12, 0.5);
}

.checkbox-group {
    margin-top: 15px;
}

.checkbox-label {
    display: flex;
    align-items: center;
    gap: 10px;
    cursor: pointer;
    font-size: 1.1rem;
    color: #fff;
}

.checkbox-label input[type="checkbox"] {
    width: 20px;
    height: 20px;
    cursor: pointer;
    margin: 0;
}

.member-card {
  background: rgba(255, 255, 255, 0.05);
  padding: 20px;
  border-radius: 10px;
  margin-bottom: 25px;
  border-left: 4px solid #a8180c;
}

.member-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.member-header h4 {
  margin: 0;
  color: #a8180c;
  font-size: 1.3rem;
}

.remove-btn {
  background: transparent;
  border: 1px solid #ff4444;
  color: #ff4444;
  font-size: 1.5rem;
  line-height: 1;
  padding: 0 8px 4px 8px;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.2s;
}

.remove-btn:hover {
  background: #ff4444;
  color: white;
}

.add-btn {
  width: 100%;
  padding: 15px;
  background: transparent;
  border: 2px dashed #444;
  color: #888;
  font-size: 1.2rem;
  font-family: 'IM Fell English SC', serif;
  cursor: pointer;
  border-radius: 10px;
  transition: all 0.3s;
  margin-top: 10px;
}

.add-btn:hover {
  border-color: #a8180c;
  color: #a8180c;
  background: rgba(168, 24, 12, 0.05);
}

.submit-btn {
  display: block;
  width: 100%;
  padding: 15px;
  background: #a8180c;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.5rem;
  font-family: 'IM Fell English SC', serif;
  cursor: pointer;
  transition: background 0.3s, transform 0.2s;
  box-shadow: 0 4px 15px rgba(168, 24, 12, 0.4);
}

.form-actions {
    display: flex;
    gap: 15px;
    margin-top: 40px;
}

.cancel-btn {
    flex: 1;
    padding: 15px;
    background: transparent;
    border: 1px solid #aaa;
    color: #aaa;
    border-radius: 8px;
    font-size: 1.5rem;
    font-family: 'IM Fell English SC', serif;
    cursor: pointer;
    transition: all 0.3s;
}

.cancel-btn:hover {
    border-color: #fff;
    color: #fff;
    background: rgba(255, 255, 255, 0.1);
}

.submit-btn {
    flex: 2;
}

.submit-btn:disabled {
    background: #555;
    cursor: not-allowed;
    box-shadow: none;
}

.submit-btn:hover:not(:disabled) {
  background: #c91e0e;
  transform: translateY(-2px);
}

.submit-btn:active:not(:disabled) {
  transform: translateY(1px);
}

@media (max-width: 600px) {
  .signup-form {
    padding: 20px;
  }
}
</style>