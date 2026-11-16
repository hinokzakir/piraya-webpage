<template>
	<div class="hello">
		<h1>Aktiva Medlemmar</h1>
		<p v-if="loading">Laddar medlemmar...</p>
		<p v-else-if="members.length === 0">Inga medlemmar tillagda ännu.</p>
		<div v-else class="members-container">
			<div v-for="member in members" :key="member.id || member.Smeknamn" class="member-card">
				<h2>{{ member.Smeknamn }}</h2>
				<img v-if="member.bild" :src="getImage(member.bild)" alt="Member Image" class="member-image" />
				<p>{{ member.Namn }}</p>
				<p>{{ member.Roll }}</p>
				<p v-if="member.quote">Kända ord: "{{ member.quote }}"</p>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: 'MembersPage',
	data() {
		return {
			members: [],
			loading: true
		};
	},
	async mounted() {
		try {
			const res = await fetch('/api/members');
			if (res.ok) {
				this.members = await res.json();
			}
		} catch (e) {
			console.error('Kunde inte hämta medlemmar:', e);
		} finally {
			this.loading = false;
		}
	},
	methods: {
		getImage(imagePath) {
			if (!imagePath) return '';
			if (imagePath.startsWith('http://') || imagePath.startsWith('https://') || imagePath.startsWith('/')) {
				return imagePath;
			}
			try {
				return require(`@/assets/${imagePath}`);
			} catch (e) {
				console.error(`Kunde inte ladda bild: ${imagePath}`, e);
				return '';
			}
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');

.hello {
	min-height: 100vh;
	background: #000;
	padding: 0;
}

.members-container {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
	gap: 32px;
	padding: 32px 8vw 32px 8vw;
	background: none;
	border-radius: 0;
	box-shadow: none;
	backdrop-filter: none;
	border: none;
}

.member-card {
	background: rgba(20, 20, 20, 0.92);
	padding: 32px 16px 24px 16px;
	border-radius: 18px;
	box-shadow: 0 2px 16px rgba(0,0,0,0.45);
	text-align: center;
	color: #fff;
	font-family: 'IM Fell English SC', serif;
	transition: transform 0.2s;
}
.member-card:hover {
	transform: translateY(-4px) scale(1.03);
	box-shadow: 0 8px 32px #a8180c55;
}

.member-image {
	width: 120px;
	height: 120px;
	border-radius: 50%;
	object-fit: cover;
	display: block;
	margin: 0 auto 18px auto;
	border: 3px solid #a8180c;
	box-shadow: 0 2px 12px #000a;
}

h1 {
	font-family: 'IM Fell English SC', serif;
	font-size: clamp(2.2rem, 6vw, 4.5rem);
	text-align: center;
	color: #fff;
	text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
	margin-top: 100px;
	margin-bottom: 32px;
	letter-spacing: 0.04em;
}

h2 {
	font-family: 'IM Fell English SC', serif;
	font-size: 1.5rem;
	color: #fff;
	margin-bottom: 0.2em;
}

p {
	color: #fff;
	font-size: 1.05rem;
	margin: 0.2em 0;
	font-family: 'IM Fell English SC', serif;
}

@media (max-width: 600px) {
	.members-container {
		padding: 4px 0 4px 0;
		gap: 12px;
	}
	.member-card {
		max-width: 80vw;
		width: 80vw;
		margin: 0 auto;
		box-sizing: border-box;
		padding: 12px 2px 10px 2px;
	}
	.member-image {
		width: 60px;
		height: 60px;
	}
	h1 {
		font-size: clamp(1rem, 8vw, 1.5rem);
		margin-bottom: 10px;
	}
	h2 {
		font-size: 1rem;
	}
	p {
		font-size: 0.95rem;
	}
}
</style>
