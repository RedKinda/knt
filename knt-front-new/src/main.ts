import { createApp } from 'vue'
import App from '/src/App.vue'
import router from '@/router/index'

import 'simple-keyboard/build/css/index.css';

import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"

createApp(App).use(router).mount('#app')
