<template>
  <MainNavbar />
  <router-view />
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
	background-image: url("./assets/bg1.jpeg");
	background-repeat: repeat-x;
	background-size: auto, auto;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #000000;
}
</style>
