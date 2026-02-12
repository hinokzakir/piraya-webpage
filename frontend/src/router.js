import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './components/HomePage.vue';
import MembersPage from './components/MembersPage.vue';
import AboutPage from './components/AboutPage.vue';
import LolPage from './components/LolPage.vue';
import OsPage from './components/OsPage.vue';
import AdminPage from './components/AdminPage.vue';

const routes = [
  { path: '/', component: HomePage },
  { path: '/members', component: MembersPage },
  { path: '/about', component: AboutPage },
  { path: '/lol', component: LolPage },
  { path: '/os', component: OsPage },
  { path: '/admin', component: AdminPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;