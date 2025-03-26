<template>
	<div class="landing-container">
		<h1>VEM ÄR VI?</h1>
		<p>Piraya är Umeås Datavetares festeri. Vi verkar främst för att underhålla linjens medlemmar, men vi arrangerar även fester och evenemang för andra. Ett bra exempel på detta är Otympliga Spelen, en årlig tradition fylld med glädje, tävlingar och gemenskap! Vi hoppas verkligen att ni alla har lust att delta i dessa underbara tävlingar och skapa minnen för livet.

Förutom Otympliga Spelen anordnar vi även andra uppskattade evenemang som chilisittning, gröt och glögg-kvällar, samt den klassiska ärtsoppslagningen. Dessa aktiviteter är till för att stärka gemenskapen och skapa en trivsam atmosfär för alla deltagare. Vi är också stolta över att vara en del av mottagningen, där vi hjälper till att välkomna nya studenter och introducera dem till vår fantastiska gemenskap.</p>
	</div>
</template>

<script>
export default {
  name: 'AboutPage'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.landing-container {
	margin: 20px;
	display: flex;
	justify-content: center;
	align-items: center;
	padding: 30px;
	flex-direction: column;
	border-style: solid;
	background: rgba(116, 218, 255, 0.5);
	border-radius: 16px;
	box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
	backdrop-filter: blur(9.5px);
	-webkit-backdrop-filter: blur(9.5px);
	border: 1px solid rgba(11, 240, 255, 0.31);
}
h1 {
  color: white;
  text-shadow: 2px 2px 4px #000000;
}

p {
  color: white;
  /*black outline */
  text-shadow: 2px 2px 4px #000000;
  /*make the text only be able to be 50% of screen size in width */
  max-width: 40%;
  /*center the text */
  margin: auto;
  /*add padding */
}

@media screen and (max-width: 600px) {
  p {
	max-width: 100%;
  }
	
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
