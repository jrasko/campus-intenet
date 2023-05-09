import { createRouter, createWebHistory } from 'vue-router'
import ManageView from '@/views/ManageView.vue'
import AddEditView from '@/views/AddView.vue'
import EditView from '@/views/EditView.vue'
import LoginView from '@/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: ManageView
    },
    {
      path: '/add',
      name: 'add',
      component: AddEditView
    },
    {
      path: '/edit/:mac',
      name: 'edit',
      component: EditView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    }
  ]
})

export default router
