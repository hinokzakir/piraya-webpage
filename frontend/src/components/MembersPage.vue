<template>
	<div class="hello">
		<h1>Aktiva Medlemmar</h1>
		<div class="members-container">
			<div v-for="(member, index) in members" :key="index" class="member-card">
				<h2>{{ member.Smeknamn }}</h2>
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
