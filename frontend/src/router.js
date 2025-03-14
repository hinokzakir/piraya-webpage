import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './components/HomePage.vue';
import MembersPage from './components/MembersPage.vue';

const routes = [
  { path: '/', component: HomePage },
  { path: '/members', component: MembersPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;