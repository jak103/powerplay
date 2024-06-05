<template>
  <q-card class="my-card">
    <q-card-section>
      <div class="text-h6">Create a Channel</div>
    </q-card-section>

    <q-separator />

    <q-card-section>
      <q-input filled v-model="channelName" label="# Channel Name" />
    </q-card-section>

    <q-card-section>
      <div class="q-gutter-sm">
        <q-radio
          v-model="channelType"
          val="private"
          label="Private - Only specific people"
        />
      </div>
      <div class="q-gutter-sm">
        <q-radio
          v-model="channelType"
          val="public"
          label="Public - Anyone in Power Play"
        />
      </div>
    </q-card-section>

    <q-separator />

    <q-card-actions vertical>
      <q-btn flat @click="handleCreateChannel">Create Channel</q-btn>
    </q-card-actions>
  </q-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const channelName = ref('');
const channelType = ref('');

function handleCreateChannel() {
  //TODO post request to create new channel. Await Chat ID. Foward to new chat.
  router.push('/chat');
}
</script>
