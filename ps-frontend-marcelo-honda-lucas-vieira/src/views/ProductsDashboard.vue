<template>
  <div class="admin-dashboard">
    <div v-if="loading" class="loading">Carregando...</div>
    <div v-else class="product-list-container">
      <h2>Lista de Produtos</h2>
      <button @click="goBack" class="back-button">Home</button>
      <div class="product-list">
        <div v-for="product in products" :key="product.id" class="product-item">
          <p>ID: {{ product.id }}</p>
          <p>Nome: {{ product.name }}</p>
          <p>Preço: R$ {{ product.price.toFixed(2) }}</p>
          <p>Descrição: {{ product.description }}</p>
          <p>Quantidade: {{ product.amount }}</p>
          <p>Disponibilidade: {{ product.disponibility ? 'Disponível' : 'Indisponível' }}</p>
          <p>Peso: {{ product.weight }}</p>
          <div class="buttons">
            <button @click="updateProduct(product.name)">Atualizar</button>
            <button @click="deleteProduct(product.name)">Excluir</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      products: [],
      loading: true
    };
  },

  async created() {
    try {
      const response = await axios.get('http://localhost:8080/produtos');
      console.log('Dados recebidos do backend:', response.data); // Verificar a resposta do backend
      if (Array.isArray(response.data)) {
        this.products = response.data;
        console.log('Produtos atribuídos ao data:', this.products); // Verificar os produtos no data
      } else {
        console.error('Estrutura inesperada na resposta do backend:', response.data);
      }
    } catch (error) {
      console.error('Erro ao carregar dados:', error);
    } finally {
      this.loading = false;
    }
  },

  methods: {
    updateProduct(productName) {
      this.$router.push({ name: 'updateProduct', params: { name: productName } });
    },

    deleteProduct(productName) {
      this.$router.push({ name: 'deleteProduct', params: { name: productName } });
    },

    goBack() {
      this.$router.push('/');
    }
  }
};
</script>

<style scoped>
.admin-dashboard {
  padding: 20px;
}

.product-list-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.product-list {
  display: flex;
  flex-wrap: wrap;
  gap: 20px; /* Espaço entre os itens */
}

.product-item {
  border: 1px solid black;
  padding: 10px;
  width: calc(25% - 20px); /* Largura para exibir três itens por linha com espaçamento */
  box-sizing: border-box;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.buttons {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
}

button {
  background-color: #4CAF50;
  color: black;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 10px;
}

button:hover {
  background-color: #3e8e41;
}

.back-button {
  background-color: #008CBA;
  align-self: flex-start; /* Ensures the button aligns to the start of the container */
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 10px;
  margin-top: 10px;
}

.back-button:hover {
  background-color: #005f73;
}

h2 {
  color: #333;
  margin-bottom: 10px;
}

.back-button, h2 {
  display: block;
  width: 100%; /* Ensures both the heading and button take the full width of the container */
  text-align: center;
}
</style>
