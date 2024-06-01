<template>
  <router-view />

  <q-dialog v-model="showPrompt">
    <q-card>
      <q-card-section>
        <div class="text-h6">Install our App</div>
      </q-card-section>
      <q-card-section>
        <div v-if="os === 'iOS'">
          To install this app on your iOS device, tap the share icon and select "Add to Home Screen".
        </div>
        <div v-if="os === 'Android'">
          To install this app on your Android device, tap the menu button and select "Add to Home Screen".
        </div>
        <div v-if="os === 'Windows' || os === 'MacOS'">
          To install this app on your computer, click on the install icon in the address bar or go to settings and select "Install".
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="Close" @click="showPrompt = false" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
defineOptions({
  name: 'App'
});

// Installation prompt below
const showPrompt = ref(false);
const os = ref('unknown');

onMounted(() => {
  os.value = detectOS();

  // Only show the install prompt once
  if (localStorage.getItem('promptShown') !== 'true') {
    if (os.value !== 'unknown') {
      showPrompt.value = true;
    }
    localStorage.setItem('promptShown', 'true');
  }
});
</script>

<script lang="ts">
import { ref, onMounted } from 'vue';

function detectOS() {
  const userAgent = navigator.userAgent || navigator.vendor;
  if (/android/i.test(userAgent)) {
    return 'Android';
  }
  else if (/iPad|iPhone|iPod/.test(userAgent)) {
    return 'iOS';
  }
  else if (/Win(dows )?/.test(userAgent)) {
    return 'Windows';
  }
  else if (/Mac/.test(userAgent)) {
    return 'MacOS';
  } 
  else {
    return 'unknown';
  } 
}
</script>
