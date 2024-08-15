<template>
  <nav>
    <ul>
      <li><NuxtLink to="/">Home</NuxtLink></li>
      <li><NuxtLink to="/cart">cart</NuxtLink></li>
      <li><NuxtLink to="/login">login</NuxtLink></li>
  
    </ul>
  </nav>
  
</template>

<script>
export default {
  name: 'NavBar',
}
</script>

<style scoped>
nav {
  width: 100%;
  height: 5rem;
  background-color: #ede6e6; /* Add background color here */
  color: rgb(28, 27, 27); /* Ensure text color contrasts with background */
  display: flex;
  align-items: center;
  padding: 0 2rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* Optional: Add shadow for better visual */
}

ul {
  margin-right: 12px;
  display: flex;
  flex-direction: row;
  gap: 2rem;
  list-style-type: none;
  margin: 0;
  padding: 0;
  justify-content: center;
  width: 100%;
}

li {
  margin-top: 20px;
}

a {
  text-decoration: none;
  color: rgb(61, 58, 58); /* Ensure link color contrasts with background */
  font-weight: 600;
  font-size: 1rem;
  transition: color 0.3s ease, transform 0.3s ease;
}

a:hover {
  color: #007BFF;
  transform: scale(1.05);
}
</style>
