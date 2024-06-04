import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'chat', component: () => import('pages/chat/chatPage.vue')},
      { path: 'profile', component: () => import('pages/profile/profilePage.vue')},
      { path: 'schedule', component: () => import('pages/schedule/schedulePage.vue')},
      { path: 'profile/edit-profile', component: () => import('pages/profile/edit-profile/editProfile.vue')},
      { path: 'profile/replace-image', component: () => import('pages/profile/replace-image/replaceImage.vue')}
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
