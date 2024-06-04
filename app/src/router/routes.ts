import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'teaminfo/:teamName', name: 'TeamInfo', component: () => import('pages/TeamInfo.vue') },
      { path: 'leagueinfo/:id', name: 'LeagueInfo', component: () => import('pages/LeagueInfo.vue') },
      { path: 'roster', name: 'RosterPage', component: () => import('pages/RosterPage.vue') },
      { path: 'standings', name: 'StandingsPage', component: () => import('pages/StandingsPage.vue') },
      { path: 'statistics', name: 'StatsPage', component: () => import('pages/StatsPage.vue') },
      { path: 'substitues', name: 'SubPage', component: () => import('pages/SubPage.vue') },
      { path: 'chat', name: 'ChatPage', component: () => import('pages/chat/chatPage.vue')},
      { path: 'chat/createchannel', name: 'CreateChannel', component: () => import('pages/chat/CreateChannel.vue')},
      { path: 'profile', name: 'ProfilePage', component: () => import('pages/profile/profilePage.vue')},
      { path: 'schedule', name: 'SchedulePage', component: () => import('pages/schedule/schedulePage.vue')},
      { path: 'game-details', name: 'GameDetailsPage', component: () => import('pages/schedule/gameDetailsPage.vue')},
    ],
  },
  {
  path: '/chat/chatId',
  component: () => import('layouts/ChatLayout.vue'),
    children: [
      { path: '', name: 'Chat', component: () => import('pages/chat/ChatUi.vue')},
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
