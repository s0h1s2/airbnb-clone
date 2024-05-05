import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AuthView from '@/views/AuthView.vue'
import ListingView from "@/views/ListingView.vue"
import TripsView from "@/views/TripsView.vue"
import ReservationView from "@/views/ReservationView.vue"
import FavoritesView from "@/views/FavoritesView.vue"
import PropertiesView from "@/views/PropertiesView.vue"
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/auth',
      name: 'auth',
      component: AuthView
    },
    {
      path: '/favorites',
      name: 'favorites',
      component: FavoritesView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: "/trips",
      name: "trips",
      component: TripsView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: "/reservations",
      name: "reservations",
      component: ReservationView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: "/listing/:id",
      name: "listing",
      component: ListingView
    },
    {
      path: '/properties',
      name: 'properties',
      component: PropertiesView,
      meta: {
        requiresAuth: true
      },

    }
  ]
})
declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth: boolean
  }
}
export default router
