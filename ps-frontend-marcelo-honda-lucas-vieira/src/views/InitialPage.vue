<template>
    <div>
        <!-- Título maior e mais acima -->
        <h1 style="font-size: 36px; margin-top: 50px;">Welcome to the initial page.</h1>

        <!-- Botões pretos grandes lado a lado, com mais espaço entre eles -->
        <div class="button-container">
            <router-link to="/produto" class="big-button">Cadastrar</router-link>
            <router-link to="/produtos" class="big-button">Listar</router-link>
        </div>
    </div>
</template>

<style scoped>
.button-container {
    display: flex;
    justify-content: space-between; /* Espaçamento maior entre os botões */
    margin-top: 40px; /* Margem superior maior para separar do título */
}

.big-button {
    display: inline-block;
    padding: 30px 60px; /* Aumente o padding para fazer os botões maiores */
    background-color: #000; /* Cor de fundo preta */
    color: white;
    text-decoration: none;
    border: none;
    border-radius: 8px;
    font-size: 1.5rem; /* Tamanho da fonte maior */
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.big-button:hover {
    background-color: #333; /* Cor de fundo mais escura no hover */
}
</style>