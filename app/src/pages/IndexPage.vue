<template>
  <p>Index page</p>
  <q-btn color="white" text-color="black" label="Subscribe" @click="subscribeToNotifications" />
  <br />
  <q-btn color="white" text-color="black" label="Send notification" @click="postNotification" />

</template>

<script setup lang="ts">
  import { api } from 'boot/axios';


  async function postNotification() {
    console.log("Send request for notification");
    const res = await api.get('/api/v1/notifications/send');
    console.log('res', res);
  }

    // Function to request user permission and subscribe to notifications
  async function subscribeToNotifications() {
    let perm = await Notification.requestPermission();
    console.log('perm', perm);
    if (perm == "granted") {
      let sw = await navigator.serviceWorker.ready;
      let sub = await sw.pushManager.subscribe({ userVisibleOnly: true, applicationServerKey: "BMPQhGq2KuP92WTzRK7S5UgLk5v8H0ZoNXXJji0J5wO3ufLm24AgelUfpe0BvasoupYfSagpGFZvwRTSBS-KYzY" });
      const res = await api.post('/api/v1/notifications/subscribe', sub)
      console.log('res', res);
    }
  }
</script>
