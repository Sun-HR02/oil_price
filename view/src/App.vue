<template style="margin: 0">

  <div class="root">
  <div class="common-layout">
    <el-container>
      <el-header class="component" style="background-color: #E6A23C">
        <h1 class="title">Sun Haoran的个人主页</h1>
      </el-header>
      <el-main>

        <div class="Carousel">
          <carousel-show></carousel-show>
        </div>

        <div class="oil" style="width: 28%; float: left">
          <el-card class="box-card" shadow="always">
            <index-oil class="component"></index-oil>
          </el-card>
        </div>

        <div class="time" style="float: left; width: 50%">
        <el-card class="box-card" shadow="always">
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
import CarouselShow from "@/components/carouselShow";

export default {
  name: 'App',
  components: {
    CarouselShow,
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
  background-color:#CFD3DC;
}
.title{
  margin: 0;
  top: 20%;
  left: 0;
  right: 0;
  bottom: 0;
}

</style>
