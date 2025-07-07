<template>
  <div class="app-root">
    <MainNavbar />
    <transition name="slide" mode="out-in">
      <router-view />
    </transition>
  </div>
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
}
</style>
