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
				<div class="links" v-show="menuOpen">
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
nav {
	background: #641818;
	padding: 10px 20px;
	background: rgb(117, 59, 59);
	background: linear-gradient(0deg, rgba(117, 59, 59, 1) 7%, rgba(168, 24, 12, 1) 100%);
	border-radius: 10px;
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

.links {
	display: flex;
	gap: 20px;
	margin-left: auto;
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

/* Media Query for Mobile */
@media screen and (max-width: 768px) {
	.nav {
		background: #641818;
		padding: 10px 20px;
		background: rgb(117, 59, 59);
		background: linear-gradient(0deg, rgba(117, 59, 59, 1) 7%, rgba(168, 24, 12, 1) 100%);
		border-radius: 10px 10px 0 0;
		position: sticky;
		top: 0;
		width: auto;
		z-index: 1000;
	}

	.hamburger {
		display: block;
	}

	.links {
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

	/* Remove the display: none from .links and handle visibility through v-show */
	.links {
		display: flex;
		/* Base state for animation */
	}

	li {
		margin: 10px 0;
	}
}
</style>