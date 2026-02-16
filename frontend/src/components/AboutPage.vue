<template>
	<div class="landing-container">
		<h1>VEM ÄR VI?</h1>
		<p>Piraya är Umeås Datavetares festeri. Vi verkar främst för att underhålla linjens medlemmar, men vi arrangerar även fester och evenemang för andra. Ett bra exempel på detta är Otympliga Spelen, en årlig tradition fylld med glädje, tävlingar och gemenskap! Vi hoppas verkligen att ni alla har lust att delta i dessa underbara tävlingar och skapa minnen för livet.

Förutom Otympliga Spelen anordnar vi även andra uppskattade evenemang som chilisittning, gröt och glögg-kvällar, samt den klassiska ärtsoppslagningen. Dessa aktiviteter är till för att stärka gemenskapen och skapa en trivsam atmosfär för alla deltagare. Vi är också stolta över att vara en del av mottagningen, där vi hjälper till att välkomna nya studenter och introducera dem till vår fantastiska gemenskap.</p>
		<img src="../assets/landing.jpeg" alt="image" width="500px" height="400px">
	</div>
</template>

<script>
export default {
  name: 'AboutPage'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');

.landing-container {
  min-height: 100vh;
  background: #000;
  margin: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 48px 8vw 48px 8vw;
}
h1 {
  font-family: 'IM Fell English SC', serif;
  color: #fff;
  font-size: clamp(2.2rem, 6vw, 4.5rem);
  text-align: center;
  text-shadow: 0 2px 16px #a8180c, 0 1px 0 #000;
  margin-top: 100px;
  margin-bottom: 32px;
  letter-spacing: 0.04em;
}
p {
  font-family: 'IM Fell English SC', serif;
  color: #fff;
  text-shadow: 0 2px 8px #000;
  max-width: 600px;
  width: 100%;
  text-align: center;
  margin: 0 auto 32px auto;
  font-size: 1.25rem;
  line-height: 1.6;
}
img {
  margin: 24px 0 0 0;
  border-radius: 18px;
  border: 3px solid #a8180c;
  box-shadow: 0 2px 24px #000a;
  width: 420px;
  height: 320px;
  object-fit: cover;
  max-width: 90vw;
}
@media screen and (max-width: 600px) {
  .landing-container {
    padding: 24px 2vw 24px 2vw;
  }
  h1 {
    font-size: clamp(1.2rem, 10vw, 2.2rem);
    margin-bottom: 18px;
  }
  p {
    font-size: 1rem;
    margin-bottom: 18px;
    max-width: 100%;
  }
  img {
    width: 90vw;
    height: 180px;
    margin-top: 16px;
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
