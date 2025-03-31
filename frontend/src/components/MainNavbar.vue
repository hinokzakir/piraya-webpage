<template>
	<nav>
		<ul>
			<img alt="Vue logo" width="75px" height="75px" src="../assets/logo.png">
			<div class="hamburger" @click="toggleMenu">
				<span></span>
				<span></span>
				<span></span>
			</div>
			<div class="links" :class="{ 'show': menuOpen }">
				<li><router-link to="/" @click="closeMenu">Hem</router-link></li>
				<li><router-link to="/members" @click="closeMenu">Medlemmar</router-link></li>
				<li><router-link to="/about" @click="closeMenu">Om oss</router-link></li>
			</div>
		</ul>
	</nav>
</template>

<script>
export default {
	data() {
		return {
			menuOpen: false
		}
	},
	methods: {
		toggleMenu() {
			this.menuOpen = !this.menuOpen
		},
		closeMenu() {
			this.menuOpen = false
		}
	}
}
</script>

<style scoped>
nav {
	background: #641818;
	padding: 10px 20px;
	background: rgb(117, 59, 59);
	background: linear-gradient(0deg, rgba(117, 59, 59, 1) 7%, rgba(168, 24, 12, 1) 100%);
	border-radius: 10px;
	position: sticky;
	top: 0;
	left: 0;
	right: 0;
	width: 97%;
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

.hamburger {
	display: none;
	/* Hidden by default on larger screens */
	cursor: pointer;
	margin-left: auto;
	margin-right: 20px;
}

.hamburger span {
	display: block;
	width: 30px;
	height: 3px;
	background: white;
	margin: 5px 0;
	transition: 0.3s;
}

.links {
	display: flex;
	gap: 20px;
	margin-left: auto;
	margin-right: 100px;
	transition: 0.3s;
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

/* Media Query for Mobile */
@media screen and (max-width: 768px) {
	.hamburger {
		display: block;
	}

	.links {
		display: none;
		/* Hidden by default on mobile */
		position: absolute;
		top: 100%;
		left: 0;
		right: 0;
		background: rgb(117, 59, 59);
		flex-direction: column;
		margin: 0;
		padding: 20px;
		border-radius: 0 0 10px 10px;
	}

	.links.show {
		display: flex;
		/* Show when menuOpen is true */
	}

	li {
		margin: 10px 0;
	}
}
</style>