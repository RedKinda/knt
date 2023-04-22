import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import UserIndex from '/src/views/UserIndex.vue'
import Products from '/src/views/Products.vue'
import App from '/src/App.vue'

const routes: Array<RouteRecordRaw> = [

  {
    path:'/',
    name:'UserIndex',
    component: UserIndex
  },
  {
    path:'/products/:id',
    name:'ProductsIndex',
    component: Products
  }
  
  // {
  //   path: '/',
  //   name: 'home',
  //   component: HomeView
  // },
  // {
  //   path: '/about',
  //   name: 'about',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  // }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
