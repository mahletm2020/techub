<template>
  <div class="cart">
    <h2>Your Cart</h2>
    <div class="cart-item">
      <img src="" alt="Product Name" class="product-image">
      <div class="item-details">
        <h3>Product Name</h3>
        <p>Quantity: 1</p>
        <p>Price: $100.00</p>
      </div>
    </div>
    <div class="cart-item">
      <img src="" alt="Product Name" class="product-image">
      <div class="item-details">
        <h3>Product Name</h3>
        <p>Quantity: 1</p>
        <p>Price: $100.00</p>
      </div>
    </div>
    <div class="cart-total">
      <h3>Total: $200.00</h3>
     </div>
       <button @click="goToCheckout">Proceed to Checkout</button>
     </div>
</template>

<script>
export default {
  props: {
    cartItems: {
      type: Array,
      default: () => []
    },
    cartTotal: {
      type: Number,
      default: 0
    }
  },
  methods: {
    goToCheckout() {
      this.$router.push('/checkout');
    }
  }
};
</script>

<style scoped>
.cart {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  background-color: #fff;
}

.cart h2 {
  font-size: 2rem;
  margin-bottom: 1.5rem;
  text-align: center;
  color: #333;
}

.cart-item {
  display: flex;
  align-items: center;
  margin-bottom: 1.5rem;
  border-bottom: 1px solid #ddd;
  padding-bottom: 1.5rem;
}

.product-image {
  width: 100px;
  height: 100px;
  margin-right: 1.5rem;
  border-radius: 8px;
  object-fit: cover;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.item-details {
  flex: 1;
}

.item-details h3 {
  font-size: 1.2rem;
  margin-bottom: 0.5rem;
  color: #333;
}

.item-details p {
  margin: 0.3rem 0;
  font-size: 1rem;
  color: #666;
}

.cart-total {
  margin-top: 2rem;
  text-align: right;
}

.cart-total h3 {
  font-size: 1.8rem;
  color: #333;
}

button {
  width: 100%;
  padding: 1rem;
  font-size: 1.2rem;
  font-weight: bold;
  color: #fff;
  background-color: #007BFF;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.3s ease;
  margin-top: 2rem;
}

button:hover {
  background-color: #0056b3;
  transform: scale(1.05);
}
</style>
