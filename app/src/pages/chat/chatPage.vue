<template>
  <q-page class="row">
    <h5 class="q-mt-lg q-mb-sm q-ml-sm">Teams</h5>
    <CardComponent
      v-for="channel in teams"
      :key="channel.channel_id"
      :type="channel.type"
      :name="channel.channel_name"
      :unread="channel.missed_chats"
      :image="channel.channel_image"
      :link="link"
    />
    <h5 class="q-mt-lg q-mb-sm q-ml-sm">Leagues</h5>
    <CardComponent
      v-for="channel in leagues"
      :key="channel.channel_id"
      :type="channel.type"
      :name="channel.channel_name"
      :unread="channel.missed_chats"
      :image="channel.channel_image"
      :link="link"
    />
    <h5 class="q-mt-lg q-mb-sm q-ml-sm">Channels</h5>
    <CardComponent
      v-for="channel in channelType"
      :key="channel.channel_id"
      :type="channel.type"
      :name="channel.channel_name"
      :unread="channel.missed_chats"
      :image="channel.channel_image"
      :link="link"
    />
    <h5 class="q-mt-lg q-mb-sm q-ml-sm">Direct Messages</h5>
    <CardComponent
      v-for="channel in dms"
      :key="channel.channel_id"
      :type="channel.type"
      :name="channel.channel_name"
      :unread="channel.missed_chats"
      :image="channel.channel_image"
      :link="link"
    />
  </q-page>
  <q-page-sticky position="bottom-right" :offset="[25, 25]">
    <q-btn fab icon="add" color="blue" @click="handleAddChannnel" />
  </q-page-sticky>
</template>

<script setup lang="ts">
import { useChannelStore } from 'app/src/stores/channelStore';
import { onMounted } from 'vue';
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import CardComponent from '../chat/CardComponent.vue';
import router from 'src/router';
import { useRouter } from 'vue-router';
const vrouter = useRouter();

const channelStore = useChannelStore();
const { channels } = channelStore;
const link = '/';
const teams = channels.filter((channel) => channel.type === 'team');
const leagues = channels.filter((channel) => channel.type === 'league');
const channelType = channels.filter((channel) => channel.type === 'channel');
const dms = channels.filter((channel) => channel.type === 'dm');

function handleAddChannnel() {
  vrouter.push({ name: 'CreateChannel' });
}

defineComponent({
  name: 'ChatPage',
  components: {
    CardComponent,
  },
});

onMounted(() => {
  console.log('Channels:', channels);
});
</script>
