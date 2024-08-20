<template>
  <div class="wrapper">
    <form @submit.prevent="handleSubmit">
      <h1>Atualizar Produto</h1>
      <div class="input-box">
        <input v-model="id" type="text" :placeholder="product ? product.id : 'ID do Produto'" readonly required class="highlight">
        <i class='bx bxs-id-card'></i>
      </div>
      <div class="input-box">
        <input v-model="formattedName" type="text" :placeholder="product ? `Name: ${product.name}` : 'Name'" @input="updateName" required>
        <i class='bx bxs-box'></i>
      </div>
      <div class="input-box">
        <input v-model="formattedPrice" type="text" :placeholder="product ? `Price: R$ ${product.price.toFixed(2)}` : 'Price'" @input="updatePrice" required>
        <i class='bx bxs-dollar-circle'></i>
      </div>
      <div class="input-box">
        <input v-model="formattedDescription" type="text" :placeholder="product ? `Description: ${product.description}` : 'Description'" @input="updateDescription" required>
        <i class='bx bxs-detail'></i>
      </div>
      <div class="input-box">
        <input v-model="formattedAmount" type="text" :placeholder="product ? `Amount: ${product.amount}` : 'Amount'" @input="updateAmount" required>
        <i class='bx bxs-cart'></i>
      </div>
      <div class="input-box">
        <input v-model="formattedWeight" @input="updateWeight" type="text" :placeholder="product ? `Weight (g): ${product.weight}` : 'Weight (g)'" required>
        <i class='bx bxs-weight'></i>
      </div>
      <div class="button-container">
        <button type="button" class="btn back-button" @click="goBack">Voltar</button>
        <button type="submit" class="btn">Atualizar</button>
      </div>
    </form>
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
      name: '',
      price: '',
      description: '',
      amount: '',
      weight: ''
    };
  },

  async created() {
    const productName = this.$route.params.name;
    try {
      const response = await axios.get(`http://localhost:8080/produto?name=${productName}`);
      this.product = response.data;
      if (!this.product) {
        console.error(`Produto com nome ${productName} não encontrado`);
        alert(`Produto com nome ${productName} não encontrado`);
      } else {
        // Preencher os campos do formulário com os dados do produto
        this.id = this.product.id;
        this.name = this.product.name;
        this.price = this.product.price.toFixed(2);
        this.description = this.product.description;
        this.amount = this.product.amount;
        this.weight = this.product.weight;
      }
    } catch (error) {
      console.error('Erro ao carregar produto:', error);
      alert('Erro ao carregar produto. Tente novamente.');
    } finally {
      this.loading = false;
    }
  },

  computed: {
    formattedName: {
      get() {
        return this.name ? `Name: ${this.name}` : '';
      },
      set(value) {
        this.name = value.replace('Name: ', '');
      }
    },
    formattedPrice: {
      get() {
        return this.price ? `Price: R$ ${this.price}` : '';
      },
      set(value) {
        // Remove "R$ " prefix before updating the price value
        this.price = value.replace('Price: R$ ', '');
      }
    },
    formattedDescription: {
      get() {
        return this.description ? `Description: ${this.description}` : '';
      },
      set(value) {
        this.description = value.replace('Description: ', '');
      }
    },
    formattedAmount: {
      get() {
        return this.amount ? `Amount: ${this.amount}` : '';
      },
      set(value) {
        this.amount = value.replace('Amount: ', '');
      }
    },
    formattedWeight: {
      get() {
        return this.weight ? `Weight (g): ${this.weight}` : '';
      },
      set(value) {
        // Remove "Weight (g): " prefix before updating the weight value
        this.weight = value.replace('Weight (g): ', '');
      }
    }
  },

  methods: {
    updateName(event) {
      this.name = event.target.value.replace('Name: ', '');
    },
    updatePrice(event) {
      this.price = event.target.value.replace('Price: R$ ', '');
    },
    updateDescription(event) {
      this.description = event.target.value.replace('Description: ', '');
    },
    updateAmount(event) {
      this.amount = event.target.value.replace('Amount: ', '');
    },
    updateWeight(event) {
      this.weight = event.target.value.replace('Weight (g): ', '');
    },

    async handleSubmit() {
      const updatedProduct = {
        name: this.name,
        price: parseFloat(this.price),
        description: this.description,
        amount: parseInt(this.amount),
        weight: this.weight
      };

      try {
        await axios.put(`http://localhost:8080/produto/alterar?id=${this.id}`, updatedProduct);
        alert('Produto atualizado com sucesso!');
        this.$router.push('/produtos');
      } catch (error) {
        console.error('Erro ao atualizar produto:', error);
        alert('Erro ao atualizar produto. Tente novamente.');
      }
    },

    goBack() {
      this.$router.push('/produtos');
    }
  }
};
</script>

<style scoped>
.highlight {
  border-color: #f0ad4e; /* Cor de destaque */
  color: #f0ad4e; /* Cor do texto */
  font-weight: bold; /* Texto em negrito */
  padding: 20px 45px 20px 20px; /* Espaçamento interno */
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: "Roboto Mono", sans-serif;
}
body {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: url(bg.jpg) no-repeat;
  background-size: cover;
  background-position: center;
}
.wrapper {
  width: 420px;
  background: transparent;
  border: 2px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(9px);
  color: #fff;
  border-radius: 12px;
  padding: 30px 40px;
}
.wrapper h1 {
  font-size: 36px;
  text-align: center;
}
.input-box {
  position: relative;
  width: 100%;
  height: 50px;
  margin: 30px 0;
}
.input-box input {
  width: 100%;
  height: 100%;
  background: transparent;
  border: none;
  outline: none;
  border: 2px solid rgba(255, 255, 255, 0.2);
  border-radius: 40px;
  font-size: 16px;
  color: #fff;
  padding: 20px 45px 20px 20px;
}
.input-box input::placeholder {
  color: #fff;
}
.btn {
  width: 100%;
  height: 45px;
  background: #fff;
  border: none;
  outline: none;
  border-radius: 40px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  font-size: 16px;
  color: #333;
  font-weight: 600;
}
.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.back-button {
  background-color: #008CBA;
}

.back-button:hover {
  background-color: #005f73;
}

.btn:not(.back-button) {
  background-color: #4CAF50;
}

.btn:not(.back-button):hover {
  background-color: #3e8e41;
}
</style>
