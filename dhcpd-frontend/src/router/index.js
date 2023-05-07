import { createRouter, createWebHistory } from 'vue-router'
import ManageView from '@/views/ManageView.vue'
import AddEditView from '@/views/AddView.vue'
import EditView from '@/views/EditView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: ManageView
    },
    {
      path: '/update',
      name: 'add',
      component: AddEditView
    },
    {
      path: '/update/:mac',
      name: 'edit',
      component: EditView
    }
  ]
})

export default router
