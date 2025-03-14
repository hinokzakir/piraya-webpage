<template>
	<nav>
		<ul>
			<img alt="Vue logo" width="100px" height="100px" src="../assets/logo.png">
			<div class="links">
				<li><router-link to="/">Hem</router-link></li>
				<li><router-link to="/members">Medlemmar</router-link></li>
				<li><router-link to="/about">Vem är vi?</router-link></li>
			</div>
		</ul>
	</nav>
</template>

<style scoped>
nav {
  background: #641818; /* Match your background */
  padding: 10px 20px;
}

ul {
  display: flex;
  align-items: center;
  list-style: none;
  padding: 0;
  margin: 0;
}

img {
  margin-left: 20px;
  height: 100px;
  width: 100px;
}

.links {
  display: flex;
  gap: 20px; /* Controls space between links */
  margin-left: auto; /* Pushes links all the way to the right */
  margin-right: 100px;
}

li {
  list-style: none;
}

a {
  color: white;
  text-decoration: none;
  font-size: 18px;
}

a:hover {
  text-decoration: underline;
}
</style>
