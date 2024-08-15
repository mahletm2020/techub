<template>
  <div class="admin-login-container">
    <h1 class="title">Admin Login</h1>
    <form @submit.prevent="submitForm" class="form-container">
      <div class="form-group">
        <label for="username">Username</label>
        <input
          id="username"
          v-model="username"
          placeholder="Enter your username"
          required
          class="input-field"
        />
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input
          id="password"
          v-model="password"
          type="password"
          placeholder="Enter your password"
          required
          class="input-field"
        />
      </div>
      <button type="submit" class="submit-button">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await this.$axios.post('/admin/login', {
          username: this.username,
          password: this.password,
        });

        if (response.data.success) {
          alert('Login successful!');
          this.$router.push('/admin/dashboard');
        } else {
          alert('Invalid credentials, please try again.');
        }
      } catch (error) {
        alert('Error: ' + error.message);
      }
    },
  },
};
</script>

<style scoped>
.admin-login-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 40px;
  background-color: #f7f9fc;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 2rem;
  color: #333;
  text-align: center;
  margin-bottom: 20px;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

label {
  font-size: 1rem;
  color: #555;
  margin-bottom: 5px;
}

.input-field {
  padding: 10px;
  font-size: 1rem;
  border: 1px solid #ddd;
  border-radius: 5px;
  transition: border-color 0.3s;
}

.input-field:focus {
  border-color: #007bff;
  outline: none;
}

.submit-button {
  padding: 15px;
  font-size: 1.2rem;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.3s;
}

.submit-button:hover {
  background-color: #0056b3;
  transform: scale(1.05);
}
</style>
