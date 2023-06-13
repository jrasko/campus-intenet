import { createRouter, createWebHashHistory } from 'vue-router'
import ManageView from '@/views/ManageView.vue'
import AddEditView from '@/views/AddView.vue'
import EditView from '@/views/EditView.vue'
import LoginView from '@/views/LoginView.vue'
import ShameView from '@/views/ShameView.vue'

const router = createRouter({
  history: createWebHashHistory(),
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
      path: '/edit/:id',
      name: 'edit',
      component: EditView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/shame',
      name: 'shame',
      component: ShameView
    }
  ]
})

export default router
