import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/mcp-market'
  },
  {
    path: '/mcp-market',
    name: 'McpMarket',
    component: () => import('@/views/McpMarket.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
