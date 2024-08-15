<template>
  <div class="container">
    <h1 class="title">Login as</h1>
    <div class="button-container">
      <NuxtLink to="/adminlogin" class="button admin-button">
        <span class="button-text">As Admin</span>
      </NuxtLink>
      <NuxtLink to="/asuser" class="button user-button">
        <span class="button-text">As User</span>
      </NuxtLink>
    </div>
  </div>
</template>

<script>
export default {
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #f5f7fa, #c3cfe2);
  padding: 20px;
}

.title {
  font-size: 2.5rem;
  color: #333;
  margin-bottom: 40px;
}

.button-container {
  display: flex;
  gap: 20px;
}

.button {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 15px 25px;
  border-radius: 8px;
  text-decoration: none;
  font-size: 1.2rem;
  font-weight: bold;
  color: #fff;
  transition: background-color 0.3s ease, transform 0.2s ease;
  cursor: pointer;
}

.button:hover {
  transform: scale(1.05);
}

.admin-button {
  background-color: #007bff;
}

.admin-button:hover {
  background-color: #0056b3;
}

.user-button {
  background-color: #28a745;
}

.user-button:hover {
  background-color: #218838;
}
</style>
