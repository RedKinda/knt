import { createApp } from 'vue'
import App from './UserIndex.vue'
import router from './router'

import Keyboard from 'simple-keyboard';
import 'simple-keyboard/build/css/index.css';

import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"


createApp(App).use(router).mount('#app')
