<template>
  <div>
    <div class="PaperItem" v-for="paper in papers" :key="paper">
      <div style="line-height:0;font-size:0;">
        <img class="PictureCover" :src="paper.cover_url">
      </div>
      <p class="Title" @click="onClickedTitle(paper)"> {{ paper.title }} </p>
      <p class="Introduction" v-for="content in paper.introduction" :key="content"> 
        {{ content }} 
      </p>
      <div class="InfoGroup">
        <div style="line-height:0;font-size:0;">
          <img class="DateIcon" src="https://img1.imgtp.com/2023/03/19/1x4xkvM5.png" >
        </div>
        <p class="Date"> {{ paper.date }} </p>
        <div style="line-height:0;font-size:0;">
          <img class="CategoryIcon" src="https://img1.imgtp.com/2023/03/19/DIF6J9jb.png" >
        </div>
        <p class="Category"> {{ paper.category }} </p>
      </div>
    </div>
  </div>
</template>

<script>
import { getArticleList } from '../api/getArticleList.js'
export default {
  data() {
    return {
      papers: [],
    }
  },
  methods: {
    onClickedTitle(data) {
      this.$router.push( { 
        name: 'paper', 
        query: {
          paperName: data.name,
        },
      });
    },
    async initData() {
      var reply = await getArticleList({
        "category": '',
      })
      if(reply.status == 200) {
        this.papers = reply.data.articles
      }
    },
  },
  mounted() {    
    this.initData()
  },
}
</script>

<style scoped>
.PaperItem {
  width: 100%;
  height: fit-content;
  background-color: rgb(255, 255, 255);
  margin-top: 30px;
  border-radius: 8px;
  box-shadow: 0px 0px 20px #b9b9b9;
  padding-bottom: 15px;
}
.PictureCover {
  width: 100%;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  vertical-align: middle;
}
.Title {
  margin-top: 15px;
  margin-left: 15px;
  margin-right: 15px;
  margin-bottom: 0px;
  font-size: 24px;
  color: black;
}
.Title:hover {
  color: #4800ce;
}
.Introduction {
  margin: 0;
  margin-left: 15px;
  margin-right: 15px;
  margin-top: 15px;
  font-size: 18px;
  color: #5a5a5a;
}
.InfoGroup {
  display: flex;
  margin-left: 15px;
  margin-top: 10px;
  height: 26px;
}
.DateIcon {
  width: 26px;
  height: 26px;
  opacity: 50%;
}
.CategoryIcon {
  width: 26px;
  height: 26px;
  margin-left: 15px;
  opacity: 50%;
  position: relative;
  top: -1px;
}
.Date {
  margin: 0;
  padding: 0;
  position: relative;
  top: 2.5px;
  margin-left: 5px;
  font-size: 16px;
  color: #5a5a5a;
}
.Category {
  margin: 0;
  padding: 0;
  position: relative;
  top: 2px;
  margin-left: 5px;
  font-size: 16px;
  color: #5a5a5a;
}
</style>