<template>
  <div class="form-container">
     <form id="shipping-form">
       <div class="form-group">
         <label for="fullName">Full Name</label>
         <input type="text" id="fullName" name="fullName" required placeholder="Enter your full name">
       </div>
 
       <div class="form-group">
         <label for="address">Address</label>
         <input type="text" id="address" name="address" required placeholder="Enter your address">
       </div>
 
       <div class="form-group">
         <label for="city">City</label>
         <input type="text" id="city" name="city" required placeholder="Enter your city">
       </div>
 
       <div class="form-group">
         <label for="state">State</label>
         <input type="text" id="state" name="state" required placeholder="Enter your state">
       </div>
 
       <div class="form-group">
         <label for="country">Country</label>
         <input type="text" id="country" name="country" required placeholder="Enter your country">
       </div>
 
       <button type="submit">Submit</button>
     </form>
   </div>
 </template>
 
 <script>
   export default {
     
   }
 </script>
 
 <style  scoped>
 /* Import Google Font */
 @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600&display=swap');
 
 body {
   font-family: 'Poppins', sans-serif;
   background-color: #f3f4f6;
   display: flex;
   justify-content: center;
   align-items: center;
   height: 100vh;
   margin: 0;
 }
 
 .form-container {
   max-width: 600px;
   width: 100%;
   padding: 2rem;
   background-color: #fff;
   border-radius: 8px;
   box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
 }
 
 .form-group {
   margin-bottom: 1.5rem;
 }
 
 label {
   display: block;
   font-weight: 600;
   margin-bottom: 0.5rem;
   color: #333;
 }
 
 input {
   width: 100%;
   padding: 0.75rem;
   font-size: 1rem;
   border: 1px solid #ccc;
   border-radius: 4px;
   transition: border-color 0.3s ease;
 }
 
 input:focus {
   border-color: #007BFF;
   outline: none;
 }
 
 button {
   padding: 0.75rem 1.5rem;
   font-size: 1rem;
   font-weight: 600;
   color: #fff;
   background-color: #007BFF;
   border: none;
   border-radius: 4px;
   cursor: pointer;
   transition: background-color 0.3s ease;
   width: 100%;
 }
 
 button:hover {
   background-color: #0056b3;
 }
 
 </style>