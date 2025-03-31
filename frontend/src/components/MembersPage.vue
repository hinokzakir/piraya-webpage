<template>
	<div class="hello">
		<h1>Aktiva Medlemmar</h1>
		<div class="members-container">
			<div v-for="(member, index) in members" :key="index" class="member-card">
				<h2>{{ member.Smeknamn }}</h2>
				<img :src="getImage(member.bild)" alt="Member Image" class="member-image" />
				<p>{{ member.Namn }}</p>
				<p>{{ member.Roll }}</p>
				<p>Kända ord: "{{ member.quote }}"</p>
			</div>
		</div>
	</div>
</template>

<script>
import membersData from '../assets/members.json'; // Import JSON file

export default {
	name: 'MembersPage',
	data() {
		return {
			members: membersData // Assign imported JSON to `members`
		};
	},

	methods: {
		getImage(imagePath) {
			try {
				// Använd require för att dynamiskt ladda bilden baserat på sökvägen från JSON
				return require(`@/assets/${imagePath}`);
			} catch (e) {
				console.error(`Kunde inte ladda bild: ${imagePath}`, e);
				return ''; // Returnera en tom sträng eller en standardbild om det misslyckas
			}
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.members-container {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); /* Auto-adjusting grid */
	gap: 20px; /* Space between cards */
	padding: 20px;
	background: rgba(116, 218, 255, 0.5);
	border-radius: 16px;
	box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
	backdrop-filter: blur(9.5px);
	-webkit-backdrop-filter: blur(9.5px);
	border: 1px solid rgba(11, 240, 255, 0.31);
}

.member-card {
	background: rgb(121, 179, 212);
	padding: 15px;
	border-radius: 8px;
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
	text-align: center;
}

.member-image {
	width: 100px; /* Fast bredd för cirkeln */
	height: 100px; /* Fast höjd för cirkeln */
	border-radius: 50%; /* Gör bilden cirkulär */
	object-fit: cover; /* Beskär bilden för att passa utan att förvrängas */
	display: block; /* Förhindrar extra utrymme under bilden */
	margin: 0 auto; /* Centrera bilden horisontellt i kortet */
}

h1 {
	font-size: 50px;
	text-align: center;
	color: white;
	text-shadow: 2px 2px 4px #000000;
}

h2 {
	font-size: 24px;
	color: white;
}
</style>
