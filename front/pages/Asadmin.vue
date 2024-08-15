<template>
  <div class="admin-container">
    <h1 class="title">Admin Page</h1>
    <form @submit.prevent="submitForm" class="form-container">
      <div class="form-group">
        <label for="productName">Product Name</label>
        <input
          id="productName"
          v-model="productName"
          placeholder="Product Name"
          required
          class="input-field"
        />
      </div>
      <div class="form-group">
        <label for="productPrice">Product Price</label>
        <input
          id="productPrice"
          v-model="productPrice"
          placeholder="Product Price"
          required
          type="number"
          class="input-field"
        />
      </div>
      <div class="form-group">
        <label for="productImage">Product Image</label>
        <input
          id="productImage"
          type="file"
          @change="onFileChange"
          required
          class="input-file"
        />
      </div>
      <button type="submit" class="submit-button">Add Product</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      productName: '',
      productPrice: '',
      productImage: null,
    };
  },
  methods: {
    onFileChange(event) {
      this.productImage = event.target.files[0];
    },
    async submitForm() {
      const formData = new FormData();
      formData.append('file', this.productImage);
      formData.append('name', this.productName);
      formData.append('price', this.productPrice);

      try {
        // Upload image to Go backend
        const uploadResponse = await this.$axios.post('/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });

        const imageUrl = uploadResponse.data.url;

        // Insert product data into Hasura
        const mutation = `
          mutation insertProduct($name: String!, $price: Float!, $image_url: String!) {
            insert_products(objects: {name: $name, price: $price, image_url: $image_url}) {
              returning {
                id
              }
            }
          }
        `;

        const variables = {
          name: this.productName,
          price: parseFloat(this.productPrice),
          image_url: imageUrl,
        };

        const client = this.$apollo.provider.defaultClient;
        await client.mutate({ mutation, variables });

        alert('Product added successfully!');
      } catch (error) {
        alert('Error: ' + error.message);
      }
    },
  },
};
</script>

<style scoped>
.admin-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 40px;
  background-color: #f7f9fc;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 2.5rem;
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

.input-file {
  font-size: 1rem;
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
