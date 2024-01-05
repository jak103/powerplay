import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: '/teams', component: () => import('pages/TeamPage.vue') },
      { path: '/chat', component: () => import('pages/ChatPage.vue') },
      { path: '/schedule', component: () => import('pages/SchedulePage.vue') },
      { path: '/finance', component: () => import('pages/FinancePage.vue') },
      { path: '/account', component: () => import('pages/AccountPage.vue') },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
