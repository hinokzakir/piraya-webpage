<template>
	<nav>
		<ul>
			<img alt="Vue logo" width="75px" height="75px" src="../assets/logo.png">
			<div class="links">
				<li><router-link to="/">Hem</router-link></li>
				<li><router-link to="/members">Medlemmar</router-link></li>
				<li><router-link to="/about">Om oss</router-link></li>
			</div>
		</ul>
	</nav>
</template>

<style scoped>
nav {
  background: #641818; /* Match your background */
  padding: 10px 20px;
  background: rgb(117,59,59);
  background: linear-gradient(0deg, rgba(117,59,59,1) 7%, rgba(168,24,12,1) 100%);
  border-radius: 10px;
  position: sticky;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
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
  color: firebrick;
  background-color: white;
}
</style>
