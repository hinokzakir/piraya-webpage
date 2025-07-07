<template>
  <div class="app-root">
    <MainNavbar />
    <transition name="slide" mode="out-in">
      <router-view />
    </transition>
    <InfoCardsSection v-if="$route.path === '/'" />
    <EventsSection v-if="$route.path === '/'" />
    <SocialLinksSection v-if="$route.path === '/'" />
  </div>
</template>

<script>
import MainNavbar from "./components/MainNavbar.vue";
import InfoCardsSection from "./components/InfoCardsSection.vue";
import EventsSection from "./components/EventsSection.vue";
import SocialLinksSection from "./components/SocialLinksSection.vue";

export default {
  name: 'App',
  components: {
    MainNavbar,
    InfoCardsSection,
    EventsSection,
    SocialLinksSection
  }
}
</script>

<style>
body {
	background-color: rgb(74, 34, 34);
	background-image: url("https://www.cardiff.ac.uk/__data/assets/image/0005/981878/GettyImages-817621116.jpg")/* url("./assets/bg2.jpeg")*/;
	background-size:auto;
	background-attachment: fixed;
	margin: 0;
	padding: 0;
	height: 100%;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #000000;
  align-items: center;
  margin: 0;
  padding: 0;
  height: 100%;
}

/* Transition styles */
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.25s cubic-bezier(0.4,0,0.2,1);
}

.slide-enter-from {
  transform: translateX(100%);
}

.slide-leave-to {
  transform: translateX(-100%);
}

.app-root {
  min-height: 100vh;
  background: #000;
  margin: 0;
  padding: 0;
  height: 100vh;
}
</style>
