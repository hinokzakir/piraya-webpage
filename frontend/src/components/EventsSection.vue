<template>
  <section class="events-section">
    <h2 class="events-title">Kommande event</h2>
    <div v-if="loading" class="events-loading">Laddar event...</div>
    <div v-else-if="error" class="events-error">Kunde inte hämta event just nu.</div>
    <div v-else class="events-cards">
      <div v-for="event in events" :key="event.id" class="event-card">
        <div class="event-card-content">
          <h3>{{ event.title }}</h3>
          <p class="event-date">{{ formatDate(event.startDate) }}<span v-if="event.endDate && !isSameDay(event.startDate, event.endDate)"> – {{ formatDate(event.endDate) }}</span></p>
          <p class="event-desc">{{ event.description }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { ref, onMounted } from 'vue';
import { db } from '../firebase';
import { collection, getDocs, query, orderBy } from 'firebase/firestore';

function toDate(ts) {
  // Firestore timestamp till JS Date
  if (!ts) return null;
  if (typeof ts.toDate === 'function') return ts.toDate();
  if (ts.seconds) return new Date(ts.seconds * 1000);
  return new Date(ts);
}

export default {
  name: 'EventsSection',
  setup() {
    const events = ref([]);
    const loading = ref(true);
    const error = ref(false);

    const fetchEvents = async () => {
      loading.value = true;
      error.value = false;
      try {
        const q = query(collection(db, 'events'), orderBy('startDate', 'asc'));
        const querySnapshot = await getDocs(q);
        events.value = querySnapshot.docs.map(doc => {
          const data = doc.data();
          return {
            id: doc.id,
            ...data,
            startDate: toDate(data.startDate),
            endDate: toDate(data.endDate)
          };
        });
        console.log('Events fetched:', events.value);
      } catch (e) {
        console.error('Error fetching events:', e);
        error.value = true;
      } finally {
        loading.value = false;
      }
    };

    const formatDate = (dateObj) => {
      if (!dateObj) return '';
      return dateObj.toLocaleDateString('sv-SE', { year: 'numeric', month: 'short', day: 'numeric' });
    };

    const isSameDay = (d1, d2) => {
      if (!d1 || !d2) return false;
      return d1.getFullYear() === d2.getFullYear() && d1.getMonth() === d2.getMonth() && d1.getDate() === d2.getDate();
    };

    onMounted(fetchEvents);

    return { events, loading, error, formatDate, isSameDay };
  }
};
</script>

<style scoped>
.events-section {
  width: 100vw;
  background: #000;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 0 64px 0;
  z-index: 4;
}
.events-title {
  font-family: 'IM Fell English SC', serif;
  color: #fff;
  font-size: clamp(2rem, 6vw, 3.5rem);
  text-align: center;
  margin-top: 0;
  margin-bottom: 36px;
  letter-spacing: 0.04em;
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
}
.events-loading, .events-error {
  color: #fff;
  font-size: 1.2rem;
  margin: 32px 0;
  font-family: 'IM Fell English SC', serif;
}
.events-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 32px;
  width: 100%;
  max-width: 1200px;
  justify-content: center;
  align-items: stretch;
}
.event-card {
  background: linear-gradient(135deg, rgba(168,24,12,0.95) 0%, rgba(220,60,40,0.85) 100%);
  border-radius: 18px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.45);
  padding: 24px 18px 18px 18px;
  min-width: 300px;
  max-width: 480px;
  flex: 1 1 340px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: left;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  transition: transform 0.2s;
}
.event-card:hover {
  transform: translateY(-4px) scale(1.03);
  box-shadow: 0 8px 32px #a8180c55;
}
.event-card-content h3 {
  font-size: 1.7rem;
  margin-bottom: 0.2em;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  letter-spacing: 0.04em;
}
.event-date {
  color: #111;
  font-size: 1.1rem;
  margin-bottom: 0.2em;
}
.event-desc {
  font-size: 1.05rem;
  color: #fff;
  font-family: 'IM Fell English SC', serif;
  margin: 0;
  margin-bottom: 0.1em;
}
@media (max-width: 900px) {
  .events-cards {
    flex-direction: column;
    gap: 10px;
    align-items: center;
    padding: 0;
  }
  .event-card {
    max-width: 80vw;
    width: 80vw;
    margin: 0 auto;
    box-sizing: border-box;
    padding: 10px 4px 8px 4px;
    align-items: center;
    text-align: center;
    flex: 0 0 auto;
  }
  .events-title {
    font-size: clamp(1.1rem, 8vw, 2.2rem);
    margin-bottom: 12px;
  }
  .event-card-content h3 {
    font-size: 1rem;
    margin-bottom: 0.3em;
  }
  .event-card p, .event-date {
    font-size: 0.92rem;
    margin-bottom: 0.2em;
  }
}
</style> 