import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "../views/HomeView.vue"
import PaperView from "../views/PaperView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/paper',
      name: 'paper',
      component: PaperView,
    },
  ]
})

export default router
