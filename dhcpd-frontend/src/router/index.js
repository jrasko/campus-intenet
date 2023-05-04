import { createRouter, createWebHistory } from 'vue-router'
import ManageView from "@/views/ManageView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: ManageView
    },
  ]
})

export default router
