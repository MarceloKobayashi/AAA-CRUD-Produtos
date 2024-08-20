<template>
  <div class="admin-dashboard">
    <div v-if="loading" class="loading">Carregando...</div>
    <div v-else>
      <div v-if="product">
        <h1>Deletar Produto</h1>
        <div class="input-box">
          <input v-model="id" type="text" :placeholder="product ? product.id : 'ID do Produto'" readonly required class="highlight">
          <i class='bx bxs-id-card'></i>
        </div>
        <div class="input-box">
          <input v-model="formattedName" type="text" :placeholder="product ? `Name: ${product.name}` : 'Name'" readonly required>
          <i class='bx bxs-box'></i>
        </div>
        <div class="input-box">
          <input v-model="formattedPrice" type="text" :placeholder="product ? `Price: R$ ${product.price.toFixed(2)}` : 'Price'" readonly required>
          <i class='bx bxs-dollar-circle'></i>
        </div>
        <div class="input-box">
          <input v-model="formattedDescription" type="text" :placeholder="product ? `Description: ${product.description}` : 'Description'" readonly required>
          <i class='bx bxs-detail'></i>
        </div>
        <div class="input-box">
          <input v-model="formattedAmount" type="text" :placeholder="product ? `Amount: ${product.amount}` : 'Amount'" readonly required>
          <i class='bx bxs-cart'></i>
        </div>
        <div class="input-box">
          <input v-model="formattedWeight" type="text" :placeholder="product ? `Weight (g): ${product.weight}` : 'Weight (g)'" readonly required>
          <i class='bx bxs-weight'></i>
        </div>
        <div class="button-container">
          <button class="btn back-button" @click="goBack">Voltar</button>
          <button class="btn confirm-button" @click="confirmDelete">Confirmar Exclusão</button>
        </div>
      </div>
      <div v-else>
        <p>Carregando produto...</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      product: null,
      loading: true,
      id: '',
      formattedName: '',
      formattedPrice: '',
      formattedDescription: '',
      formattedAmount: '',
      formattedWeight: ''
    };
  },

  created() {
    this.fetchProduct(this.$route.params.name);
  },

  methods: {
    async fetchProduct(productName) {
      try {
        const response = await axios.get(`http://localhost:8080/produto?name=${productName}`);
        this.product = response.data;
        if (!this.product) {
          console.error(`Produto com nome ${productName} não encontrado`);
          alert(`Produto com nome ${productName} não encontrado`);
        } else {
          this.id = this.product.id;
          this.formattedName = `Name: ${this.product.name}`;
          this.formattedPrice = `Price: R$ ${this.product.price.toFixed(2)}`;
          this.formattedDescription = `Description: ${this.product.description}`;
          this.formattedAmount = `Amount: ${this.product.amount}`;
          this.formattedWeight = `Weight (g): ${this.product.weight}`;
        }
      } catch (error) {
        console.error('Erro ao carregar produto:', error);
        alert('Erro ao carregar produto. Tente novamente.');
      } finally {
        this.loading = false;
      }
    },

    async confirmDelete() {
      if (this.product) {
        try {
          await axios.delete(`http://localhost:8080/produto?name=${this.product.name}`);
          alert('Produto deletado com sucesso!');
          this.$router.push('/produtos');
        } catch (error) {
          console.error('Erro ao deletar produto:', error);
          alert('Erro ao deletar produto. Tente novamente.');
        }
      } else {
        alert('Nenhum produto selecionado para deleção.');
      }
    },

    goBack() {
      this.$router.push('/produtos');
    }
  }
};
</script>

<style scoped>
.admin-dashboard {
  width: 420px;
  background: transparent;
  border: 2px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(9px);
  color: #fff;
  border-radius: 12px;
  padding: 30px 40px;
}

.loading {
  text-align: center;
  margin-top: 20px;
}

.button-container {
  margin-top: 20px;
}

input[readonly] {
  background-color: #f5f5f5; /* Cor de fundo para campos somente leitura */
  color: #333; /* Cor do texto */
  border: 2px solid #000; /* Borda preta */
  border-radius: 4px; /* Borda arredondada */
  padding: 8px; /* Espaçamento interno */
  font-size: 16px; /* Tamanho da fonte */
}

button {
  background-color: #4CAF50;
  color: black;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 10px;
  margin-right: 10px;
}

button:hover {
  background-color: #3e8e41;
}

.back-button {
  background-color: #008CBA;
}

.back-button:hover {
  background-color: #005f73;
}

.confirm-button {
  background-color: #f44336;
}

.confirm-button:hover {
  background-color: #d32f2f;
}

.input-box {
  position: relative;
  width: 100%;
  height: 50px;
  margin: 15px 0;
}

.input-box input {
  width: calc(100% - 40px);
  height: 100%;
  background: transparent;
  border: none;
  outline: none;
  border: 2px solid rgba(255, 255, 255, 0.2);
  border-radius: 40px;
  font-size: 16px;
  color: #fff;
  padding: 0 20px;
}

.input-box i {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 20px;
}

.highlight {
  border-color: #f0ad4e; /* Cor de destaque */
  color: #f0ad4e; /* Cor do texto */
  font-weight: bold; /* Texto em negrito */
}
</style>
