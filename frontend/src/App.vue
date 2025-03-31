<template>
  <MainNavbar />
  <transition name="slide" mode="out-in">
	<router-view />
  </transition>
</template>

<script>
import MainNavbar from "./components/MainNavbar.vue";

export default {
  name: 'App',
  components: {
    MainNavbar
  }
}
</script>

<style>
body {
	background-color: rgb(74, 34, 34);
	background-image: url("./assets/bg1.jpeg"), url("./assets/bg2.jpeg");
	background-repeat: no-repeat, no-repeat;
	background-size: cover, cover;
	background-position: left top, right bottom;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #000000;
  align-items: center;
}

/* Transition styles */
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}

.slide-enter-from {
  transform: translateX(100%);
}

.slide-leave-to {
  transform: translateX(-100%);
}
</style>
