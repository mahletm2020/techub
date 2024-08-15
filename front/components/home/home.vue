<template>
  <div>
    <div class="home">
      <div class="title">
        <h2>Welcome to E.com</h2>
      </div>
      <div class="imgg">
        <!-- <img src="@/public/images/headset.png" alt="Headset Image" /> -->
      </div>
    </div>

    <div class="search-bar">
      <input type="text" placeholder="Search for products..." />
      <button>Search</button>
    </div>

    <div class="categories">
      <button>All Categories</button>
      <button>Electronics</button>
      <button>Fashion</button>
      <button>Home</button>
      <button>Beauty</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomePage'
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap');

body {
  font-family: 'Poppins', sans-serif;
  background-color: #f5f5f5;
  margin: 0;
  padding: 0;
}

.home {
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 2rem;
  background-color: #fff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  margin: 2rem auto;
  max-width: 1200px;
  gap: 2rem;
}

.title {
  font-size: 2rem;
  font-weight: 700;
  text-align: left;
  color: #333;
  letter-spacing: 1.5px;
}

.imgg {
  display: flex;
  justify-content: center;
  align-items: center;
  flex: 1;
}

.imgg img {
  max-width: 100%;
  height: auto;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.imgg img:hover {
  transform: scale(1.05);
}

.search-bar {
  display: flex;
  justify-content: center;
  margin-bottom: 2rem;
}

.search-bar input {
  padding: 0.75rem;
  width: 300px;
  border: 1px solid #ccc;
  border-radius: 4px 0 0 4px;
  outline: none;
}

.search-bar button {
  padding: 0.75rem 1rem;
  border: 1px solid #ccc;
  background-color: #007bff;
  color: #fff;
  border-radius: 0 4px 4px 0;
  cursor: pointer;
}

.search-bar button:hover {
  background-color: #0056b3;
}

.categories {
  display: flex;
  justify-content: center;
  gap: 1rem;
}

.categories button {
  padding: 0.75rem 1rem;
  border: 1px solid #ccc;
  background-color: #fff;
  color: #333;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.categories button:hover {
  background-color: #007bff;
  color: #fff;
}
</style>
