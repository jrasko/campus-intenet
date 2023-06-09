import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi'
import '@mdi/font/css/materialdesignicons.css'
import axios from 'axios'
import { isLoggedIn, noAuthPages } from '@/utils'

const vuetify = createVuetify({
  components: components,
  directives: directives,
  theme: {
    defaultTheme: 'dark'
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi
    }
  }
})

axios.defaults.baseURL = 'http://localhost'
if (import.meta.env.MODE === 'development') {
  axios.defaults.baseURL += ':8080'
}

router.beforeEach((to) => {
  if (!isLoggedIn() && !noAuthPages.includes(to.name)) {
    return { name: 'login' }
  }
})
const app = createApp(App)
app.use(router)
app.use(vuetify)

app.mount('#app')
