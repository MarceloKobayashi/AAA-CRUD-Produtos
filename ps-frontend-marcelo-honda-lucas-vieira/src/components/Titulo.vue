<template>
  <div class="title-container">
    <h1>Products Manager</h1>
    <h2>
      <span :class="{ animated: true }">The</span> Manager
    </h2>
  </div>
</template>

<script>
export default {
  data() {
    return {
      words: ['best', 'funniest', 'coolest'],
      currentWordIndex: 0,
    };
  },
  computed: {
    currentWord() {
      return this.words[this.currentWordIndex];
    },
  },
  mounted() {
    this.startWordRotation();
  },
  methods: {
    startWordRotation() {
      setInterval(() => {
        this.currentWordIndex = (this.currentWordIndex + 1) % this.words.length;
      }, 1500);
    },
  },
};
</script>

<style scoped>
.title-container {
  text-align: center;
  margin-bottom: 2rem;
  color: #1b1b1b;
}

h1 {
  font-size: 3rem; /* Aumentei o tamanho da fonte */
  margin-bottom: 0.5rem;
  font-weight: bold; /* Adicionei peso à fonte */
  text-shadow: 0px 0px 10px rgba(0, 0, 0, 0.5); /* Adicionei sombra à fonte */
}

h2 {
  font-size: 2rem; /* Aumentei o tamanho da fonte */
  font-weight: bold; /* Adicionei peso à fonte */
}

h2 span {
  color: #007bff; /* Mudei a cor para um azul mais chamativo */
  font-weight: bold;
  animation: pulse 1.5s infinite; /* Adicionei animação de pulsação */
}

.animated {
  animation: pulse 1.5s infinite;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
  }
}
</style>