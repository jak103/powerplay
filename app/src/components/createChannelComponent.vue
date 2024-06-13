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
          val="dm"
          label="Private - Only specific people"
        />
      </div>
      <div class="q-gutter-sm">
        <q-radio
          v-model="channelType"
          val="channel"
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
import { useChannelStore } from '../stores/channelStore';

const store = useChannelStore();

interface channel {
  channel_name: string;
  channel_id: string;
  channel_image: string | null;
  missed_chats: number;
  type: string;
}

const router = useRouter();

const channelName = ref('');
const channelType = ref('');

function handleCreateChannel() {
  const input: channel = {
    channel_name: channelName.value,
    channel_id: '',
    channel_image: null,
    missed_chats: 0,
    type: channelType.value,
  };

  //TODO post request to create new channel. Await Chat ID. Foward to new chat.
  debugger;
  const createdChannel: channel = store.addChannel(input);
  console.log(createdChannel);
  router.push({ name: 'Chat', params: { chatId: createdChannel.channel_id } });
}
</script>
