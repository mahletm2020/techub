<template>
  <div class="order-summary">
    <div class="order-item" v-for="item in items" :key="item.id">
      <img :src="item.image" :alt="item.name" class="order-item-img">
      <div class="order-item-details">
        <h3>{{ item.name }}</h3>
        <p>Amount: {{ item.amount }}</p>
        <p>Price: ${{ item.price }}</p>
      </div>
    </div>
    <a href="#" class="view-more">View More</a>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          id: 1,
          name: 'Headset',
          amount: 1,
          price: 59.99,
          image: '@public/images/headset.png'
        },
        {
          id: 2,
          name: 'Mouse',
          amount: 2,
          price: 29.99,
          image: '@public/images/mouse.png'
        }
        // Add more items as needed
      ]
    }
  }
}
</script>

<style scoped>
.order-summary {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.order-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem 0;
  border-bottom: 1px solid #ddd;
}

.order-item:last-child {
  border-bottom: none;
}

.order-item-img {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 8px;
}

.order-item-details {
  flex-grow: 1;
}

.order-item-details h3 {
  margin: 0;
  font-size: 1.2rem;
}

.order-item-details p {
  margin: 0.2rem 0;
  color: #555;
}

.view-more {
  display: block;
  text-align: center;
  margin-top: 1rem;
  color: #007BFF;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.3s ease;
}

.view-more:hover {
  color: #0056b3;
}
</style>
