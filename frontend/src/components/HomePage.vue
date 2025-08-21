<template>
  <div class="homepage-root">
    <div v-if="!imagesLoaded" class="loader-wrapper">
      <div class="loader"></div>
    </div>
    <div v-else class="carousel-wrapper">
      <transition-group name="fade" tag="div">
        <img
          v-for="(img, idx) in images"
          v-show="idx === currentImage"
          :key="img"
          :src="img"
          class="carousel-image"
          alt="carousel"
        />
      </transition-group>
      <div class="background-gradient"></div>
      <div class="bottom-vignette"></div>
    </div>
    <div class="hero-title-wrapper" v-if="imagesLoaded">
      <h1 class="hero-title">PIRAYA PARTY PROVIDERS</h1>
      <div class="hero-subtitle">FÖRSER FESTER TILL UMEÅS DATAVETARE M.FL. SEDAN 1985</div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      images: [
        require('../assets/homepage/SAM_0419.webp'),
        require('../assets/homepage/DSC04425.webp'),
        require('../assets/homepage/SAM_0452.webp'),
        require('../assets/homepage/DSCF0112.webp')
      ],
      currentImage: 0,
      intervalId: null,
      imagesLoaded: false
    };
  },
  mounted() {
    // Preload all images
    this.preloadImages(this.images).then(() => {
      this.imagesLoaded = true;
      // Start carousel if more than one image
      if (this.images.length > 1) {
        this.intervalId = setInterval(() => {
          this.currentImage = (this.currentImage + 1) % this.images.length;
        }, 3000);
      }
    });
  },
  beforeUnmount() {
    if (this.intervalId) clearInterval(this.intervalId);
  },
  methods: {
    preloadImages(imageArray) {
      // Returns a promise that resolves when all images are loaded
      return Promise.all(
        imageArray.map(src => {
          return new Promise(resolve => {
            const img = new window.Image();
            img.src = src;
            img.onload = resolve;
            img.onerror = resolve;
          });
        })
      );
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Orbitron:wght@900&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@900&family=Protest+Revolution&display=swap');
@import url('https://fonts.googleapis.com/css2?family=Pirata+One&display=swap');
@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English+SC&display=swap');

.homepage-root {
  position: relative;
  min-height: 100vh;
  width: 100vw;
  overflow: hidden;
  background: #000;
  margin: 0;
  padding: 0;
  padding-top: 80px;
}
.carousel-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 100vw;
  height: 100vh;
  z-index: 1;
}
.carousel-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  object-fit: cover;
  opacity: 0.6;
  transition: opacity 0.5s cubic-bezier(0.4,0,0.2,1);
}
.background-gradient {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 100vw;
  height: 100vh;
  background: linear-gradient(180deg, rgba(0,0,0,0.7) 0%, rgba(0,0,0,0.4) 60%, rgba(0,0,0,0.85) 100%);
  z-index: 2;
  pointer-events: none;
}
.hero-title-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  z-index: 3;
  pointer-events: none;
}
.hero-title {
  font-family: 'IM Fell English SC', serif;
  font-size: clamp(2.5rem, 10vw, 6rem);
  font-weight: 900;
  color: #fff;
  letter-spacing: 0.01em;
  text-align: center;
  white-space: nowrap;
  text-shadow:
    0 0 8px #fff,
    0 0 16px #a8180c,
    0 0 32px #a8180c,
    0 0 64px #fff,
    0 2px 8px #000;
  user-select: none;
}
.hero-subtitle {
  font-family: 'IM Fell English SC', serif;
  font-size: clamp(1.4rem, 4vw, 2.7rem);
  color: #fff;
  text-align: center;
  margin-top: 0.4rem;
  letter-spacing: 0.02em;
  text-shadow:
    0 0 4px #fff,
    0 0 8px #a8180c,
    0 1px 4px #000;
  user-select: none;
}
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s cubic-bezier(0.4,0,0.2,1);
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
@media (max-width: 768px) {
  .carousel-wrapper,
  .carousel-image,
  .background-gradient,
  .hero-title-wrapper {
    width: 100vw;
    height: 100vh;
    min-height: 100svh;
    min-width: 100vw;
  }
  .hero-title {
    font-size: clamp(1rem, 6vw, 1.5rem);
    padding: 0 4px;
    word-break: break-word;
  }
  .hero-subtitle {
    font-size: clamp(0.8rem, 4vw, 1.1rem);
    margin-top: 0.3rem;
    padding: 0 10px;
  }
}
.loader-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #000;
  z-index: 10;
}
.loader {
  border: 8px solid #fff;
  border-top: 8px solid #a8180c;
  border-radius: 50%;
  width: 64px;
  height: 64px;
  animation: spin 1s linear infinite;
}
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
.bottom-vignette {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 120px;
  width: 100vw;
  pointer-events: none;
  z-index: 4;
  background: linear-gradient(180deg, rgba(0,0,0,0) 0%, rgba(0,0,0,0.7) 70%, #000 100%);
}
</style>
