import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue'), meta: { title: 'Power Play' } },
      { path: 'chat', component: () => import('pages/chat/chatPage.vue'), meta: { title: 'Chat' }},
      { path: 'profile', component: () => import('pages/profile/profilePage.vue'), meta: { title: 'Profile' }},
      { path: 'schedule', component: () => import('pages/schedule/schedulePage.vue'), meta: { title: 'Schedule' }},
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
