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

router.beforeEach((to) => {
  if (!localStorage.getItem('jwt') && to.name !== 'login') {
    return { name: 'login' }
  }
})
const app = createApp(App)
app.use(router)
app.use(vuetify)

app.mount('#app')
