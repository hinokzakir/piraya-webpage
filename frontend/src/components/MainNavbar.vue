<template>
	<nav>
		<ul>
			<img alt="Vue logo" width="75px" height="75px" src="../assets/logo.png">
			<div class="hamburger" @click="toggleMenu" :class="{ 'open': menuOpen }">
				<span></span>
				<span></span>
				<span></span>
			</div>
			<transition name="menu-slide">
				<div class="links" :class="{ 'mobile-active': menuOpen }">
					<li><router-link to="/" @click="closeMenu">Hem</router-link></li>
					<li><router-link to="/members" @click="closeMenu">Medlemmar</router-link></li>
					<li><router-link to="/about" @click="closeMenu">Om oss</router-link></li>
				</div>
			</transition>
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
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');
nav {
	/* background: #641818; */
	/* padding: 10px 20px; */
	background: linear-gradient(180deg, rgba(168,24,12,0.7) 0%, rgba(168,24,12,0.3) 60%, rgba(168,24,12,0) 100%);
	padding: 10px 20px;
	position: sticky;
	top: 0;
	left: 0;
	width: auto;
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
	transition: all 0.3s ease;
}

.hamburger.open span:nth-child(1) {
	transform: rotate(45deg) translate(5px, 5px);
}

.hamburger.open span:nth-child(2) {
	opacity: 0;
}

.hamburger.open span:nth-child(3) {
	transform: rotate(-45deg) translate(7px, -7px);
}

/* Desktop styles */
.links {
	display: flex;
	gap: 20px;
	margin-left: auto;
	margin-right: 100px;
}

li {
	list-style: none;
	font-family: 'IM Fell English SC', serif;
	font-size: 1.35em;
}

a {
	color: white;
	text-decoration: none;
	font-size: 1.25em;
	font-family: 'IM Fell English SC', serif;
}

a:hover {
	text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #fff, 0 0 20px #fff, 0 0 30px #fff, 0 0 40px #fff, 0 0 50px #fff;
}

/* Menu slide animation */
.menu-slide-enter-active,
.menu-slide-leave-active {
	transition: all 0.3s ease;
}

.menu-slide-enter-from,
.menu-slide-leave-to {
	transform: translateY(-100%);
	opacity: 0;
}

/* Mobile styles */
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

	.links.mobile-active {
		display: flex;
		/* Show when menu is open */
	}

	li {
		margin: 10px 0;
	}
}
</style>