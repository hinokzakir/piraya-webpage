<template>
  <section class="info-cards-section">
    <h2 class="info-cards-title">Vad gör vi?</h2>
    <div class="info-cards">
      <div class="info-card">
        <div class="info-icon">
          <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#a8180c" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 12l10-7 10 7-10 7-10-7z"/><path d="M22 12v5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5"/></svg>
        </div>
        <h3>Bidrar till Insparken</h3>
        <p>Vi är med och arrangerar insparken för nya studenter och ser till att alla känner sig välkomna!</p>
      </div>
      <div class="info-card">
        <div class="info-icon">
          <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#a8180c" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2"/><path d="M16 2v4"/><path d="M8 2v4"/><path d="M3 10h18"/></svg>
        </div>
        <h3>Event & Fester</h3>
        <p>Vi anordnar fester, sittningar och roliga evenemang för att skapa gemenskap och minnen för livet.</p>
      </div>
      <div class="info-card">
        <div class="info-icon">
          <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#a8180c" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 19a6 6 0 0 0-12 0"/><circle cx="12" cy="7" r="4"/><path d="M16 11a4 4 0 0 1 0 8"/><path d="M8 11a4 4 0 0 0 0 8"/></svg>
        </div>
        <h3>Samarbete</h3>
        <p>Vi samarbetar med andra festerier och föreningar vid universitetet för ännu bättre studentliv.</p>
      </div>
    </div>
  </section>
</template>

<style scoped>
.info-cards-section {
  width: 100vw;
  background: #000;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0 0 64px 0;
  z-index: 4;
}
.info-cards-title {
  font-family: 'IM Fell English SC', serif;
  color: #fff;
  font-size: clamp(2.2rem, 7vw, 4.2rem);
  text-align: center;
  margin-top: 0;
  margin-bottom: 36px;
  letter-spacing: 0.04em;
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
}
.info-cards {
  display: flex;
  flex-direction: row;
  gap: 32px;
  width: 100%;
  max-width: 1200px;
  justify-content: center;
  align-items: stretch;
  margin-top: 48px;
}
.info-card {
  background: rgba(20,20,20,0.97);
  border-radius: 18px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.45);
  padding: 32px 24px 24px 24px;
  flex: 1 1 0;
  min-width: 220px;
  max-width: 340px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  transition: transform 0.2s;
}
.info-card:hover {
  transform: translateY(-4px) scale(1.03);
  box-shadow: 0 8px 32px #a8180c55;
}
.info-icon {
  margin-bottom: 18px;
}
.info-card h3 {
  font-size: 1.3rem;
  margin-bottom: 0.5em;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  letter-spacing: 0.04em;
}
.info-card p {
  font-size: 1.05rem;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  margin: 0;
}
@media (max-width: 900px) {
  .info-cards {
    flex-direction: column;
    gap: 24px;
    align-items: center;
    margin-top: 32px;
  }
  .info-card {
    max-width: 95vw;
    min-width: 0;
    width: 100%;
  }
}
</style> 