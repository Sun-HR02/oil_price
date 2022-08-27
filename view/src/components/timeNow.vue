<template>

  <h1>{{Year}}年{{Month}}月{{D}}日</h1>
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
      Second: "",
      D: ""
    }
  },
  mounted() {
    //
    setInterval(this.calculateTime,1000)
    this.calculateTime()
  },
  methods: {
    calculateTime() {
      var date = new Date()
      this.Day =  this.translateDay(date.getDay())
      this.Year = date.getFullYear()
      this.Month = this.translateMonth(date.getMonth())
      this.Hour = date.getHours()
      this.Minute = date.getMinutes()
      this.Second = date.getSeconds()
      this.D = date.getDate()
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
      },
      translateMonth(month) {
        if(month==0){
          return "一"
        }else if(month ==1 ){
          return "二"
        }else if(month ==2 ){
          return "三"
        }else if(month ==3 ){
          return "四"
        }else if(month ==4 ){
          return "五"
        }else if(month ==5 ){
          return "六"
        }else if(month ==6 ){
          return "七"
        }else if(month ==7 ){
          return "八"
        }else if(month ==8 ){
          return "九"
        }else if(month ==9 ){
          return "十"
        }else if(month ==10 ){
          return "十一"
        }else if(month ==11 ){
          return "十二"
        }
        },
      }
}
</script>

<style scoped>

</style>