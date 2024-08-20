<template>
  <div id="app">
    <Navbar />
    <Titulo />
    <div class="content">
      <router-view />
      <!-- <AppFooter /> -->
    </div>
  </div>
</template>


<script setup>
import AppFooter from './components/Footer.vue'
import NavegationBar from './components/NavegationBar.vue'
import Titulo from './components/Titulo.vue'
</script>

<style scoped>
#app {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* Garantir que os elementos filhos comecem do topo */
  align-items: center;
  min-height: 100vh;
  box-sizing: border-box;
  overflow: auto; /* Permitir rolagem */
  font-family: "Roboto Mono", -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
}

#app::before {
  content: "";
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(to bottom, #34C759, #007bff);
  background-size: 250% 400%;
  animation: gradientAnimation 5s ease infinite;
  overflow: visible; /* Removido overflow: visible; para o background */
  z-index: -1;
}

@keyframes gradientAnimation {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.content {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%; /* Certifique-se de que o conteúdo ocupe toda a largura */
}

h1 {
  margin-top: 0;
}

/* Estilos adicionais para garantir que a Navbar ocupe toda a largura */
/* nav {
  width: 100%;
} */
</style>

