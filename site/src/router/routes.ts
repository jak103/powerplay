import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/IndexPage.vue') }],
  },
  {
    path: '/information',
    component: () => import('layouts/InformationLayout.vue'),
    children: [
      { path: '', component: () => import('pages/information/InformationPage.vue') }
      // { path: 'how-to-join', component: () => import('pages/information/HowToJoin.vue') },
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
