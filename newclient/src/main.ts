import './assets/base.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { OhVueIcon, addIcons } from 'oh-vue-icons'
import { OiSearch, PrBars, FaUserAlt, MdClose, FaUmbrellaBeach, GiWindmill,MdVilla} from "oh-vue-icons/icons"
import Toast from "vue-toastification"
import "vue-toastification/dist/index.css";
import { useUserStore } from './stores/userStore'

const app = createApp(App)
addIcons(OiSearch, PrBars, FaUserAlt, MdClose, FaUmbrellaBeach, GiWindmill, MdVilla)
app.use(createPinia())
app.use(router)
app.use(Toast)
const userStore = useUserStore()

router.beforeEach((route) => {
    if (userStore.isAuth && route.name == 'auth') {
        return { name: 'home' }
    }
    if (!userStore.isAuth && route.meta.requiresAuth) {
        return { name: 'auth' }
    }
})

app.component("v-icon", OhVueIcon)
app.mount('#app')
