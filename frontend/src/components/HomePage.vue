<template>
	<div class="hello">
		<div class="landing-container">
			<div class="landing-text">
				<div class="title">PIRAYA PARTY PROVIDERS</div>
				<h1>DET FESTANDE FISKERIET</h1>
				<p>FÖRSER FESTER TILL UMEÅS DATAVETARE M.FL. SEDAN 1985</p>
				<p>"VISSA MÄNNISKOR ÄR BÄTTRE ÄN ANDRA" - PIRAYA</p>
			</div>
			<div class="landing-image">
				<img alt="Vue logo" width="500px" height="400px" src="../assets/landing.jpeg">
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: 'HomePage'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Protest+Revolution&display=swap');

.title {
	font-size: 50px;
	font-weight: 900;
	color: rgb(161, 20, 20);
	font-family: "Impact", sans-serif;
	text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000;
}

.landing-container {
	margin: 20px;
	display: flex;
	justify-content: center;
	align-items: center;
	padding: 30px;
	border-style: solid;
	background: rgba(116, 218, 255, 0.5);
	border-radius: 16px;
	box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
	backdrop-filter: blur(9.5px);
	-webkit-backdrop-filter: blur(9.5px);
	border: 1px solid rgba(11, 240, 255, 0.31);
}

.landing-image {
	border-style: solid;
	border-color: rgb(73, 114, 142);
}

@media (max-width: 768px) {
	.landing-container {
		flex-direction: column;
		/* Stack items vertically */
	}
}

h1 {
	color: white;
	text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000;
}

p {
	color: white;
}

.landing-text {
	margin-right: 100px;
}

h3 {
	margin: 40px 0 0;
}

ul {
	list-style-type: none;
	padding: 0;
}

li {
	display: inline-block;
	margin: 0 10px;
}

a {
	color: #42b983;
}
</style>
