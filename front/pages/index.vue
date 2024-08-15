<template>
  <div>
    <Nav/>
    <Home/>

    <!-- <Productlist/> -->
    <Productdetail/>
    
  </div>
</template>

<script>
import Home from '~/components/home/home.vue';
import Nav from '~/components/nav.vue';
import Productdetail from '../components/home/productdetail.vue';
import Productlist from '~/components/home/productlist.vue';

  export default {
    components: {
      Nav,
      Home,
      Productlist,
      Productdetail
    }
  }
</script>

<style  scoped>
.nav {
  margin-top: -14px;
}

</style>