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
		<div class="info-container">
			<div class="landing-events">
				<div class="title">NYHETER</div>
				<div class="fb-page" data-href="https://www.facebook.com/piraya.party.providers" data-tabs="timeline" data-width="500" data-height="500" data-small-header="false" data-adapt-container-width="true" data-hide-cover="false" data-show-facepile="true"><blockquote cite="https://www.facebook.com/piraya.party.providers" class="fb-xfbml-parse-ignore"><a href="https://www.facebook.com/piraya.party.providers">Piraya Party Providers</a></blockquote></div>
			</div>
			<div class="landing-events">
				<div class="title"> KOMMANDE EVENTS</div>
				<div class="events-container">
					<template v-for="(event, index) in event" :key="index">
						<div class="event-card">
							<h3>{{ event.name }}</h3>
							<p>{{ event.description }}</p>
							<p>{{ event.dateStart }}</p>
							<p>{{ event.dateEnd}}</p>
						</div>
					</template>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import eventsData from '../assets/events.json';

export default {
	name: 'HomePage',
	data() {
		return {
			event: eventsData // Assign imported JSON to `members`
		};
	}
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
	text-shadow: #000 1px 0 10px;
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

.info-container {
	display: flex;
	justify-content: center;
	align-items: center;
	padding: 30px;
}

.landing-events {
	padding: 30px;
	border-style: solid;
	background: rgba(116, 218, 255, 0.5);
	border-radius: 16px;
	box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
	backdrop-filter: blur(9.5px);
	-webkit-backdrop-filter: blur(9.5px);
	border: 1px solid rgba(11, 240, 255, 0.31);
	margin: 50px;
}

.event-container {
	display: flex;
	flex-direction: column;
}

.event-card {
	margin: 10px;
	padding: 10px;
	border-style: solid;
	border-color: rgb(73, 114, 142);
	background: rgba(116, 218, 255, 0.5);
	box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
	backdrop-filter: blur(9.5px);
	-webkit-backdrop-filter: blur(9.5px);
	border: 1px solid rgba(11, 240, 255, 0.31);
	border-radius: 16px;
}

.landing-image {
	border-style: solid;
	border-color: rgb(73, 114, 142);
	margin-right: 50px;
	align-items: center;
}

@media (max-width: 768px) {
	.landing-container {
		flex-direction: column;
		/* Stack items vertically */
	}
	.landing-image {
		margin-right: 0;
	}
	img {
		width: 100%;
		height: auto;
	}
	.landing-text {
		margin-right: 0;
	}
	.info-container {
		flex-direction: column;
	}
	.event-container {
		margin-right: 0;
	}
}

h1 {
	color: white;
	text-shadow: -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000, 1px 1px 0 #000;
}

p {
	color: white;
	text-shadow: #000 1px 0 10px;
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
