<template>
 <h1> 今日油价</h1>

  <li v-for="item in msg" :key="item"> {{item}}</li>

</template>


<script>
import axios from "axios";

export default {
  name: "index-oil",
  data() {
    return {
      msg: []
    }

  },
  mounted() {
    axios
        .get('http://localhost/prices')
        .then(response => (this.msg = response))
        .catch(function (error) { // 请求失败处理
          console.log(error);
        });
  }
}
</script>

<style scoped>

</style>