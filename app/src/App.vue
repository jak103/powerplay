<template>
  <router-view />

  <q-dialog v-model="showPrompt">
    <q-card>
      <q-card-section>
        <div class="text-h6">Install our App</div>
      </q-card-section>
      <q-card-section>
        <div v-if="osName === 'iOS'">
          To install this app on your iOS device, tap the share icon and select "Add to Home Screen".
        </div>
        <div v-if="osName === 'Android'">
          To install this app on your Android device, tap the menu button and select "Add to Home Screen".
        </div>
        <div v-if="osName === 'Windows' || osName === 'macOS'">
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
import { ref, onMounted } from 'vue';
import { UAParser } from 'ua-parser-js';

defineOptions({
  name: 'App',
});

// Get the OS name
let parser = new UAParser();
const osObject = parser.getResult();
const osName = ref(osObject.os.name);

// Installation prompt below
const showPrompt = ref(false);
onMounted(() => {
  // Only show the install prompt once
  if (localStorage.getItem('shownInstallPrompt') !== 'true') {
      showPrompt.value = true;
    }
    localStorage.setItem('shownInstallPrompt', 'true');
  });
</script>
