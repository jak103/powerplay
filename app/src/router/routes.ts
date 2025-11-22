import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue'), meta: { title: 'Power Play' } },
      { path: 'teaminfo/:teamName', name: 'TeamInfo', component: () => import('pages/TeamInfo.vue'), meta: { title: 'Team Info' } },
      { path: 'leagueinfo/:id', name: 'LeagueInfo', component: () => import('pages/LeagueInfo.vue'), meta: { title: 'League Info' } },
      { path: 'roster', name: 'RosterPage', component: () => import('pages/RosterPage.vue'), meta: { title: 'Roster' } },
      { path: 'standings', name: 'StandingsPage', component: () => import('pages/StandingsPage.vue'), meta: { title: 'Standings' } },
      { path: 'statistics', name: 'StatsPage', component: () => import('pages/StatsPage.vue'), meta: { title: 'Statistics' } },
      { path: 'substitues', name: 'SubPage', component: () => import('pages/SubPage.vue'), meta: { title: 'Substitutes' } },
      { path: 'chat', name: 'ChatPage', component: () => import('pages/chat/chatPage.vue'), meta: { title: 'Chat' } },
      { path: 'chat/createChannel', name: 'CreateChannel', component: () => import('pages/chat/CreateChannel.vue'), meta: { title: 'Create Channel' } },
      { path: 'profile', name: 'ProfilePage', component: () => import('pages/profile/profilePage.vue'), meta: { title: 'Profile' } },
      { path: 'profile/edit-profile', name: 'EditProfile', component: () => import('pages/profile/edit-profile/editProfile.vue'), meta: { title: 'Edit Profile' }},
      { path: 'profile/replace-image', name: 'ReaplaceImage', component: () => import('pages/profile/replace-image/replaceImage.vue'), meta: { title: 'Replace Image' } },
      { path: 'schedule', name: 'SchedulePage', component: () => import('pages/schedule/schedulePage.vue'), meta: { title: 'Schedule' } },
      { path: 'game-details/:gameId/:teamId', name: 'GameDetailsPage', component: () => import('pages/schedule/gameDetailsPage.vue'), meta: { title: 'Game Details' } },
    ],
  },
  {
  path: '/chatpage/',
  component: () => import('layouts/ChatLayout.vue'),
    children: [
      { path: '/:chatId', name: 'Chat', component: () => import('pages/chat/ChatUi.vue')},
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
