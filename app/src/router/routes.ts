import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
    { path: '', component: () => import('pages/HomePage.vue') },
    { path: 'teaminfo/:teamName', name: 'TeamInfo', component: () => import('pages/TeamInfo.vue') },
    { path: 'leagueinfo/:id', name: 'LeagueInfo', component: () => import('pages/LeagueInfo.vue') },
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
