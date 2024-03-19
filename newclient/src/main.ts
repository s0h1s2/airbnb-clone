import './assets/base.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { OhVueIcon, addIcons } from 'oh-vue-icons'
import { OiSearch, PrBars, FaUserAlt,MdClose } from "oh-vue-icons/icons"
const app = createApp(App)
addIcons(OiSearch, PrBars, FaUserAlt,MdClose)
app.use(createPinia())
app.use(router)
app.component("v-icon", OhVueIcon)
app.mount('#app')
