<template>
  <div class="wrapper">
    <form @submit.prevent="handleSubmit">
      <h1>Registrar Produto</h1>
      <div class="input-box">
        <input v-model="name" type="text" placeholder="Nome do Produto" required>
        <i class='bx bxs-box'></i>
      </div>
      <div class="input-box">
        <input v-model="price" @input="convertToFloat" type="text" :placeholder="`Preço: R$ ${formattedPrice}`" required>
        <i class='bx bxs-dollar-circle'></i>
      </div>
      <div class="input-box">
        <input v-model="description" type="text" placeholder="Descrição do produto" required>
        <i class='bx bxs-detail'></i>
      </div>
      <div class="input-box">
        <input v-model="amount" type="text" placeholder="Quantidade" required>
        <i class='bx bxs-cart'></i>
      </div>
      <div class="input-box">
        <input v-model="weight" @input="convertToString" type="text" placeholder="Peso (g)" required>
        <i class='bx bxs-weight'></i>
      </div>
      <button type="button" class="btn home-btn" @click="goHome">Home</button>
      <button type="submit" class="btn">Registrar</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      name: '',
      price: '',
      description: '',
      amount: '',
      weight: ''
    };
  },

  setup() {
    const router = useRouter();
    return { router };
  },

  computed: {
    formattedPrice() {
      // Format price to two decimal places if it's a valid number
      if (!isNaN(parseFloat(this.price))) {
        return parseFloat(this.price).toFixed(2);
      } else {
        return '';
      }
    }
  },

  methods: {
    convertToFloat(event) {
      // Remove non-numeric characters except dot (.) and format to two decimal places
      const value = event.target.value.replace(/[^\d.]/g, '');
      const parts = value.split('.');
      if (parts.length > 2) {
        // More than one dot found, keep only the first part
        event.target.value = parts[0] + '.' + parts.slice(1).join('');
      } else if (parts.length === 2) {
        // Ensure only one dot remains and format to two decimal places
        event.target.value = parts[0] + '.' + parts[1].slice(0, 2);
      } else {
        // No dot or only one integer part
        event.target.value = parts[0];
      }
      this.price = event.target.value;
    },

    convertToString(event) {
      this.weight = event.target.value.toString();
    },

    async handleSubmit() {
      console.log('Produto registrado:', {
        name: this.name,
        price: this.price,
        description: this.description,
        amount: this.amount,
        weight: this.weight
      });

      try {
        const response = await axios.post('http://localhost:8080/produto', {
          name: this.name,
          price: parseFloat(this.price),
          description: this.description,
          amount: parseInt(this.amount),
          weight: this.weight
        });

        if (response.status === 201) {
          this.$router.push('/produtos');
        } else {
          alert("Falha no registro.");
        }
      } catch (error) {
        console.error('Erro ao conectar ao servidor:', error);
        alert("Erro ao conectar ao servidor. Por favor, tente novamente mais tarde.");
      }
    },

    goHome() {
      this.$router.push('/');
    }
  }
};
</script>

<style scoped>
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
    border: 2px solid rgba(255, 255, 255, .2);
    backdrop-filter: blur(9px);
    color: #fff;
    border-radius: 12px;
    padding: 30px 40px;
  }
  .wrapper h1 {
    font-size: 36px;
    text-align: center;
  }
  .wrapper .input-box {
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
    border: 2px solid rgba(255, 255, 255, .2);
    border-radius: 40px;
    font-size: 16px;
    color: #fff;
    padding: 20px 45px 20px 20px;
  }
  .input-box input::placeholder {
    color: #fff;
  }
  .input-box i {
    position: absolute;
    right: 20px;
    top: 30%;
    transform: translate(-50%);
    font-size: 20px;
  }
  .wrapper .btn {
    width: 100%;
    height: 45px;
    background: #fff;
    border: none;
    outline: none;
    border-radius: 40px;
    box-shadow: 0 0 10px rgba(0, 0, 0, .1);
    cursor: pointer;
    font-size: 16px;
    color: #333;
    font-weight: 600;
  }
  .wrapper .home-btn {
    background: #008CBA;
    margin-bottom: 10px; /* Add margin to create space between buttons */
  }
  .wrapper .home-btn:hover {
    background: #005f73;
  }
</style>
