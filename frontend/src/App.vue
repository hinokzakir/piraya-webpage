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
	background-image: url("https://www.cardiff.ac.uk/__data/assets/image/0005/981878/GettyImages-817621116.jpg")/* url("./assets/bg2.jpeg")*/;
	background-size:auto;
	background-attachment: fixed;
	margin: 0;
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
