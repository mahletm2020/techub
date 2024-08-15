<template>
  <div>
 <Nav/>
  <div class="login-container">

    <h2>Login</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="text"  :v-model="username">username</label>
        <input type="text" id="text" required>
      </div>
      <div class="form-group">
        <label for="email" :v-model="email">Email</label>
        <input type="email" id="email" required>
      </div>
      <div class="form-group">
        <label for="password" :v-model="password">Password</label>
        <input type="password" id="password" required>
      </div>
      <button type="submit">Login</button>
    </form>
    <p class="signup-link">Don't have an account? <a href="/signup">Sign up</a></p>
  </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      username:'',
      email:'',
      password:'',
    }
  },
  methods: {
 
  }
}
</script>
<style scoped>
.login-container {
  max-width: 400px;
  margin: auto;
  padding: 30px;
  background-color: #f9f9f9;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  margin-bottom: 20px;
}

form {
  display: flex;
  flex-direction: column;
}

.form-group {
  margin-bottom: 15px;
}

label {
  font-weight: bold;
  font-family: 'Arial Narrow', Arial, sans-serif;
  
}

input[type="email"],
input[type="password"] ,
input[type="text"]
{
  margin-left: 30px;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

button {
  padding: 10px;
  font-size: 16px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

.signup-link {
  text-align: center;
  margin-top: 15px;
}

.signup-link a {
  color: #007bff;
  text-decoration: none;
}

.signup-link a:hover {
  text-decoration: underline;
}
</style>
//done with backend auth  now u need to connect ui with hasura ntn is done only configeration