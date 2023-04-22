import { createApp } from 'vue'
import App from './UserIndex.vue'

import 'simple-keyboard/build/css/index.css';
import axios from 'axios'
import VueAxios from 'vue-axios'

import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"

// createApp(App).use(router).mount('#app')
createApp(App).use(VueAxios, axios).mount('#app')
