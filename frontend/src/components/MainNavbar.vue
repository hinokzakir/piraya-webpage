<template>
	<nav :class="{ 'mobile-menu-open': menuOpen }">
		<ul>
			<img alt="Vue logo" width="75px" height="75px" src="../assets/logo.png">
			<div class="hamburger" @click="toggleMenu" :class="{ 'open': menuOpen }">
				<span></span>
				<span></span>
				<span></span>
			</div>
			<transition name="menu-slide">
				<div class="links" :class="{ 'mobile-active': menuOpen }">
					<li v-for="(item, idx) in navItems" :key="idx">
						<router-link :to="item.to" @click="closeMenu">{{ item.label }}</router-link>
					</li>
				</div>
			</transition>
		</ul>
	</nav>
</template>

<script>
export default {
	data() {
		return {
			menuOpen: false,
			navItems: [
				{ to: '/', label: 'Hem' },
				{ to: '/members', label: 'Medlemmar' },
				{ to: '/about', label: 'Om oss' },
				{ to: '/os', label: 'OS' }
			]
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
	position: fixed;
	top: 0;
	left: 0;
	width: 100vw;
	z-index: 10000;
	transition: background 0.3s;
}

nav.mobile-menu-open {
	background: rgba(168,24,12,0.7) !important;
	background-image: none !important;
	transition: background 0.3s;
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
	margin-right: 32px;
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
	transition: all 0.4s cubic-bezier(0.4,0,0.2,1);
}

.menu-slide-enter-from,
.menu-slide-leave-to {
	opacity: 0;
	transform: translateY(-30px);
}

/* Mobile styles */
@media screen and (max-width: 768px) {
	nav {
		height: 70px;
		padding-top: 18px;
		padding-bottom: 18px;
	}
	.hamburger {
		display: block;
	}

	img {
		display: none;
	}

	.links {
		display: none;
		position: absolute;
		top: 100%;
		left: 0;
		right: 0;
		background: rgb(117, 59, 59);
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin: 0;
		padding: 30px 0 30px 0;
		border-radius: 0 0 16px 16px;
		min-height: 0;
		transition: min-height 0.4s cubic-bezier(0.4,0,0.2,1), background 0.3s;
		overflow: hidden;
	}

	.links.mobile-active {
		display: flex;
		min-height: 300px;
		background: rgba(168,24,12,0.7) !important;
		background-image: none !important;
		transition: min-height 0.4s cubic-bezier(0.4,0,0.2,1), background 0.3s;
	}

	.links li {
		margin: 18px 0;
		width: 100%;
		text-align: center;
		transition: opacity 0.3s, transform 0.3s;
		opacity: 1;
		transform: translateY(0);
	}

	.links:not(.mobile-active) li {
		opacity: 0;
		transform: translateY(-20px);
		pointer-events: none;
	}

	a {
		font-size: 1.5em;
		padding: 12px 0;
		width: 100%;
		display: block;
	}

	img {
		height: 40px;
		width: 40px;
		margin-left: 8px;
	}
}
</style>