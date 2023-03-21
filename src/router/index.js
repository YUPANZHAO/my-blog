import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "../views/HomeView.vue"
import PaperView from "../views/PaperView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
      meta: { 
        keepAlive: true,
      }, 
    },
    {
      path: '/paper',
      name: 'paper',
      component: PaperView,
      meta: { 
        keepAlive: false,
      }, 
    },
  ],
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  },
})

export default router
