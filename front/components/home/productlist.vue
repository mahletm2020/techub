<template>
  <div class="product-list-page">
    <!-- <div class="product-list">
      <div v-if="loading">Loading...</div>
      <div v-if="error">Error loading products</div>
      <div v-for="product in products" :key="product.id" class="product-card">
        <img :src="product.image_url" :alt="product.name" />
        <div class="product-info">
          <h3>{{ product.name }}</h3>
          <p>${{ product.price }}</p>
        </div>
      </div>
    </div> -->
    hi
  </div>
</template>

<script setup>
// import { ref, watchEffect } from 'vue';
// import { useQuery } from '@vue/apollo-composable';
// import GET_PRODUCTS from '~/queries/getProduct.gql';

// const { result, loading, error } = useQuery(GET_PRODUCTS);

// const products = ref([]);

// watchEffect(() => {
//   if (result.value) {
//     products.value = result.value.product_list;
//   }
// });
</script>




<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap');

body {
  font-family: 'Poppins', sans-serif; 
}

.product-list-page {
  background-color: #f5f5f5;
  padding: 2rem;
}

.product-list {
  display: flex;
  flex-direction: row;
  gap: 2rem;
  justify-content: center;
  flex-wrap: wrap;
}

.product-card {
  background-color: #ffffff;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  width: 200px;
  margin: 1rem;
  transition: transform 0.2s;
}

.product-card:hover {
  transform: scale(1.05);
}

.product-card img {
  width: 100%;
  height: auto;
  object-fit: cover;
}

.product-info {
  padding: 1rem;
  text-align: center;
}

.product-info h3 {
  font-size: 1.2rem;
  margin: 0.5rem 0;
}

.product-info p {
  font-size: 1rem;
  color: #888888;
  margin: 0;
}
</style>
