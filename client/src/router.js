import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    { path: '/', redirect: '/plants/start'},
    { path: '/plants/start', name: 'Start', component: () => import('./components/pages/StartPage.vue')},
    { path: '/plants/sale', name: 'Sale', component: () => import('./components/pages/SalePlantPage.vue')},
    { path: '/plants/sale/:species', name: 'SpecificPlant', component: () => import('./components/pages/PlantForSale.vue')},
    { path: '/plants/care', name: 'Care', component: () => import('./components/pages/CarePage.vue')},
    { path: '/plants/user', name: 'User', component: () => import('./components/pages/UserPage.vue')},
    { path: '/plants/ads', name: 'Ads', component: () => import('./components/pages/MyAdsPage.vue')}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router