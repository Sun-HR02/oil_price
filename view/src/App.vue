<template style="margin: 0">

  <div class="root">
  <div class="common-layout">
    <el-container>
      <el-header class="component">
        <h1 class="title">Sun Haoran的个人主页</h1>
      </el-header>
      <el-main>
        <div class="oil" style="width: 28%; float: left">
          <el-card class="box-card">
            <index-oil class="component"></index-oil>
          </el-card>
        </div>

        <div class="time" style="float: left; width: 50%">
        <el-card class="box-card">
        <time-now></time-now>
        </el-card>
        </div>


      </el-main>
    </el-container>
  </div>
  </div>
</template>


<script>
import IndexOil from "@/components/oilPrice";
import TimeNow from "@/components/timeNow";

export default {
  name: 'App',
  components: {
    TimeNow,
    IndexOil
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ffffff;
  margin: 0;
  bottom: 0px;
  top: 0px;
  left: 0px;
  right: 0px;
  position: absolute;
  background: antiquewhite;
}
.title{
  background: rgba(63, 60, 60, 0.5);
  margin: 0;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

</style>
