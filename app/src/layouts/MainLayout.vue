<template>
  <q-layout view="lHh Lpr fff" class="bg-grey-1">
    <q-header elevated class="bg-white text-grey-8" height-hint="64">
      <q-toolbar class="GPL__toolbar" style="height: 64px">

        <q-toolbar-title shrink class="row items-center no-wrap">
          <span class="q-ml-sm">League Manager</span>
        </q-toolbar-title>

        <q-space />

        <q-select v-model="roleModel" :options="roleOptions" label="Role" style="width: 100px" label-color="red" />


        <q-space />
        <!-- TODO Only show this in dev -->
        <div class="q-gutter-sm row items-center no-wrap">
          <q-btn round flat>
            <q-avatar size="26px">
              <img src="https://cdn.quasar.dev/img/boy-avatar.png">
            </q-avatar>
            <q-tooltip>Account</q-tooltip>
          </q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container class="GPL__page-container">
      <router-view />
      <template v-if="$q.screen.lt.md">
        <q-footer style="background-color: white;" class="">
          <menu-buttons />
        </q-footer>
      </template>
      <template v-if="$q.screen.gt.sm">
        <q-page-sticky expand position="left">
          <div class="fit q-pt-xl q-px-sm column">
            <menu-buttons />
          </div>
        </q-page-sticky>
      </template>

    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">

import { ref } from 'vue';

import MenuButtons from 'src/components/MenuButtons.vue';

const roleModel = ref(null);
const roleOptions = ['Player', 'Captain', 'Staff', 'Manager']

</script>

<style lang="sass">
.GPL

  &__toolbar
    height: 64px

  &__toolbar-input
    width: 35%

  &__drawer-item
    line-height: 24px
    border-radius: 0 24px 24px 0
    margin-right: 12px

    .q-item__section--avatar
      padding-left: 12px
      .q-icon
        color: #5f6368

    .q-item__label:not(.q-item__label--caption)
      color: #3c4043
      letter-spacing: .01785714em
      font-size: .875rem
      font-weight: 500
      line-height: 1.25remns

    &--storage
      border-radius: 0
      margin-right: 0
      padding-top: 24px
      padding-bottom: 24px

  &__side-btn
    &__label
      font-size: 12px
      line-height: 24px
      letter-spacing: .01785714em
      font-weight: 500

  @media (min-width: 1024px)
    &__page-container
      padding-left: 94px
  </style>