<template>
  <div class="root">
    <NameCard :class="classNameCard "
      AvatarUrl="https://img1.imgtp.com/2023/03/15/Yu7q4P8x.jpg"
      AvatarSize="100px"
      Name="YUPAN ZHAO"
      Description="Huizhou University"
      :IconGroup="iconGroup" />
  </div>
</template>

<script>
import NameCard from '../components/NameCard.vue'
export default {
  components: {
    NameCard
  },
  data() {
    return {
      screenWidth: window.innerWidth,
      screenHeight: window.innerHeight,
      iconGroup: [
        { 
          url: "https://img1.imgtp.com/2023/03/15/6OB2Q7Go.png",
          link: "https://github.com/YUPANZHAO",
          hint: "GitHub",
        },
        { 
          url: "https://img1.imgtp.com/2023/03/15/CCetn0ZT.png",
          hint: "yupanzhao2022@163.com",
        },
        { 
          url: "https://img1.imgtp.com/2023/03/15/PK67VI4A.png", 
          link: "http://47.115.209.240:8080",
          hint: "个人博客",
        },
      ],
    }
  },
  computed: {
    classNameCard() {
      return this.screenHeight > this.screenWidth ? 'NameCard-Phone' : 'NameCard';
    },
  },
  methods: {
    windowSizeChange() {
      this.screenWidth = window.innerWidth;
      this.screenHeight = window.innerHeight;
    },
  },
  mounted() {
    this.windowSizeChange();
    window.onresize = () => {
      if(!this.timer) {
        this.timer = true;
        let that = this;
        setTimeout(function() {
          that.windowSizeChange();
          that.timer = false;
        }, 400);
      }
    };
  },
  destroyed() {
    window.onresize = null;
  },
}
</script>

<style>
.root {
  width: 100%;
  height: 100%;
}
.NameCard-Phone {
  width: fit-content;
  margin: 0 auto;
  position: relative;
  top: 13%;
}
.NameCard {
  width: fit-content;
  position: relative;
  top: 50%;
  transform: translateY(-50%);
  left: 25%
}
</style>