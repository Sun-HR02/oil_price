<template>
<h1>今天是<span style="font-weight: lighter">星期{{Day}}</span></h1>
  <span>{{Hour}}:{{Minute}}:{{Second}}</span>

</template>

<script>
export default {
  name: "timeNow",
  data() {
    return {
      Day : "",
      Month: "",
      Year: "",
      Hour: "",
      Minute: "",
      Second: ""
    }
  },
  mounted() {
    //
    this.calculateTime()
  },
  methods: {
    calculateTime() {
      var date = new Date()
      this.Day =  this.translateDay(date.getDay())
      this.Year = date.getFullYear()
      this.Month = date.getMonth()
      this.Hour = date.getHours()
      this.Minute = date.getMinutes()
      this.Second = date.getSeconds()
    },
    translateDay(day){
      if(day==1){
        return "一"
      }else if(day ==2 ){
        return "二"
      }else if(day ==3 ){
        return "三"
      }else if(day ==4 ){
        return "四"
      }else if(day ==5 ){
        return "五"
      }else if(day ==6 ){
        return "六"
      }else if(day ==7 ){
        return "七"
      }
    }
  }
}
</script>

<style scoped>

</style>