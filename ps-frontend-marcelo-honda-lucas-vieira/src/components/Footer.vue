<template>
    <footer class="app-footer">
      <div class="footer-content">
        <p>&copy; 2024 IMGEN2K. All rights reserved.</p>
        <button @click="goToInsert" class="login-button">Register Product</button>
      </div>
    </footer>
  </template>
  
  <script>
  export default {
    name: 'AppFooter',
    methods: {
      goToInsert() {
        this.$router.push('/produto');
      }
    }
  };
  </script>

<style scoped>
.app-footer {
  background-color: #1b1b1b; /* Mudança na cor de fundo para cinza */
  color: #ffa31a; /* Cor do texto alterada para #ffa31a */
  padding: 20px;
  text-align: center;
  position: fixed;
  width: 100%;
  bottom: 0;
}

.footer-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.login-button {
  background-color: white;
  color: #1b1b1b;
  border: none;
  padding: 10px 20px;
  cursor: pointer;
  border-radius: 5px;
}

.login-button:hover {
 background-color: hsla(0, 0%, 100%, 0.1);
}
</style>