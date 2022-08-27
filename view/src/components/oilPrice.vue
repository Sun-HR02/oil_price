<template>
  <div id="root" style="background-color: #E5EAF3">
   <div id="header">
     今日油价
   </div>
    <div id="content">
  <el-table :data="msg" height="250">
    <el-table-column fixed prop="Prov" label="省份" width="100" />
    <el-table-column prop="NinetyTwo" label="92号汽油" width="180" />
    <el-table-column prop="NinetyFive" label="95号汽油" width="180"/>
    <el-table-column prop="NinetyEight" label="98号汽油" width="180" />
    <el-table-column prop="Zero" label="0号柴油" width="180" />
  </el-table>
    </div>
  </div>
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
        .get('http://localhost:8081/api/prices')
        .then(response => (this.msg = response.data))
        .catch(function (error) { // 请求失败处理
          console.log(error);
        });
  }
}
</script>
