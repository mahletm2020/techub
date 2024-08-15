<template>
  <div class="payment-method">
    <h2>Payment Method</h2>
    <div class="payment-option">
      <input type="radio" id="credit-card" name="payment" value="credit-card">
      <label for="credit-card">Credit Card</label>
    </div>
    <div class="payment-option">
      <input type="radio" id="paypal" name="payment" value="paypal">
      <label for="paypal">PayPal</label>
    </div>
    <div class="payment-option">
      <input type="radio" id="bank-transfer" name="payment" value="bank-transfer">
      <label for="bank-transfer">Bank Transfer</label>
    </div>
  </div>
</template>


<script>
  export default {
    
  }
</script>
<style scoped>
.payment-method {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  background-color: #fff;
}

.payment-method h2 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  text-align: center;
}

.payment-option {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}

.payment-option input[type="radio"] {
  margin-right: 0.5rem;
}

.payment-option label {
  font-size: 1rem;
  color: #555;
}
</style>
