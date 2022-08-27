<template>
  <el-carousel :interval="4000" type="card" height="400px" >
    <el-carousel-item>
      <a :href="urlG">
         <el-image  :src=picG :fit="fit" />
      </a>
    </el-carousel-item>

    <el-carousel-item>
      <a :href="urlB">
        <el-image  :src=picB :fit="fit" />
      </a>
    </el-carousel-item>

    <el-carousel-item>
      <a :href="urlA">
        <el-image  :src=picA :fit="fit" />
      </a>
    </el-carousel-item>

    <el-carousel-item>
      <a :href="urlJ">
        <el-image  :src=picJ :fit="fit" />
      </a>
    </el-carousel-item>

  </el-carousel>
</template>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>

<script>
export default {
  name: "carouselShow",
  data(){
    return {
      urlB: "http://bilibili.com",
      picB: "https://img1.baidu.com/it/u=3149333398,404439120&fm=253&fmt=auto&app=138&f=JPEG?w=800&h=500",
      urlG:"http://github.com",
      picG: "https://img0.baidu.com/it/u=749251901,1179736680&fm=253&fmt=auto&app=138&f=JPEG?w=889&h=500",
      urlJ:"https://juejin.cn/",
      picJ: "https://img1.baidu.com/it/u=1569830584,2078216289&fm=253&fmt=auto&app=138&f=JPEG?w=778&h=500",
      urlA:"http://apple.com",
      picA: "https://img1.baidu.com/it/u=2051511161,2811101998&fm=253&fmt=auto&app=138&f=JPEG?w=889&h=500",
    }
  }
}
</script>

