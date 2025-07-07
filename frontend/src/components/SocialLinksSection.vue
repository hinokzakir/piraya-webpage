<template>
  <section class="social-section">
    <h2 class="social-title">Hitta oss på sociala medier</h2>
    <div class="social-links">
      <a href="https://www.facebook.com/piraya.party.providers" target="_blank" rel="noopener" class="social-link">
        <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" fill="#1877F2"/><path d="M16 8.5h-2a1 1 0 0 0-1 1V12h3l-.5 3H13v7" stroke="#fff" stroke-width="2"/></svg>
        Facebook
      </a>
      <a href="https://www.instagram.com/piraya_party_providers/" target="_blank" rel="noopener" class="social-link">
        <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" fill="#E1306C"/><circle cx="12" cy="12" r="5" stroke="#fff" stroke-width="2"/><circle cx="17" cy="7" r="1.5" fill="#fff"/></svg>
        Instagram
      </a>
    </div>
  </section>
</template>

<style scoped>
.social-section {
  width: 100vw;
  background: #000;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 0 64px 0;
  z-index: 4;
}
.social-title {
  font-family: 'IM Fell English SC', serif;
  color: #fff;
  font-size: clamp(2rem, 6vw, 3.5rem);
  text-align: center;
  margin-top: 0;
  margin-bottom: 36px;
  letter-spacing: 0.04em;
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
}
.social-links {
  display: flex;
  flex-direction: row;
  gap: 48px;
  justify-content: center;
  align-items: center;
}
.social-link {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  font-size: 1.3rem;
  text-decoration: none;
  transition: transform 0.2s, color 0.2s;
  filter: drop-shadow(0 2px 8px #a8180c88);
}
.social-link:hover {
  color: #a8180c;
  transform: scale(1.08);
}
.social-link svg {
  margin-bottom: 8px;
}
@media (max-width: 600px) {
  .social-links {
    flex-direction: column;
    gap: 24px;
  }
}
</style> 