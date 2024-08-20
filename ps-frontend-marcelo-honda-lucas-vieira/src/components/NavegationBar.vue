<template>
    <nav class="navbar">
      <div class="navbar-brand">
        <router-link to="/" class="navbar-item">TESTE</router-link>
      </div>
      <div class="navbar-menu">
        <div class="navbar-start">
          <a href="https://gitlab.com/ltp2-c-megalodonte" class="navbar-item" target="_blank">About</a>
        </div>
        <div class="navbar-end">
          <router-link to="/produto" class="navbar-item" v-if="!isUserRoute && !isAdminRoute">Insert</router-link>
          <router-link to="/produtos" class="navbar-item" v-if="!isUserRoute && !isAdminRoute">Dashboard</router-link>
          <LogoutButton v-if="isUserRoute || isAdminRoute" />
        </div>
      </div>
    </nav>
  </template>

<script>
</script>
  
<style scoped>
.navbar {
  width: 100%;
  background-color: #1b1b1b; /* Mudança na cor de fundo para cinza */
  padding: 1rem;
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar-brand, .navbar-start, .navbar-end {
  display: flex;
  align-items: center;
}

.navbar-menu {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.navbar-start {
  flex-grow: 1;
}

.navbar-item {
  padding: 0 1rem;
  color: white;
  text-decoration: none;
  border-radius: 5px; /* Bordas arredondadas nos botões */
}

.navbar-item:hover {
  background-color: hsla(0, 0%, 100%, 0.1);
}
</style>

