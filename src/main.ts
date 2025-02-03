import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { VueQueryPlugin } from '@tanstack/vue-query'

import PrimeVue from 'primevue/config'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import Avatar from 'primevue/avatar'
import Dropdown from 'primevue/dropdown'
import ProgressBar from 'primevue/progressbar'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'

// import 'primevue/resources/themes/aura-light-green/theme.css'
import Aura from '@primevue/themes/aura'
import 'primeflex/primeflex.css'
import 'primeicons/primeicons.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

// app.use(PrimeVue)
app.use(PrimeVue, {
  // Default theme configuration
  theme: {
    preset: Aura,
    options: {
      //   // prefix: 'p',
      // darkModeSelector: 'system',
      darkModeSelector: '.my-app-dark',
      //   // cssLayer: false,
    },
  },
})
app.use(createPinia())
app.use(router)
app.use(VueQueryPlugin)

app.component('Button', Button)
app.component('Textarea', Textarea)
app.component('Avatar', Avatar)
app.component('Dropdown', Dropdown)
app.component('ProgressBar', ProgressBar)
app.component('Dialog', Dialog)
app.component('InputText', InputText)

app.mount('#app')
