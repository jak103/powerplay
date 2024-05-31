import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'teaminfo/:teamName', name: 'TeamInfo', component: () => import('pages/TeamInfo.vue') },
      { path: 'leagueinfo/:id', name: 'LeagueInfo', component: () => import('pages/LeagueInfo.vue') },
      { path: 'chat', name: 'ChatPage', component: () => import('pages/chat/chatPage.vue')},
      { path: 'profile', name: 'ProfilePage', component: () => import('pages/profile/profilePage.vue')},
      { path: 'schedule', name: 'SchedulePage', component: () => import('pages/schedule/schedulePage.vue')},
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
