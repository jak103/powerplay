<template>
  <div>
    <h1>My Profile</h1>
    <div class="vstack d-inline-flex gap-2">
      <NuxtLink class="btn btn-link mt-auto" to="/">Website Home</NuxtLink>
      <NuxtLink class="btn btn-primary" to="/app/sign-in">Sign Out</NuxtLink>
      <ThemeToggle />
    </div>
    <div class="vstack gap-2">
      <p></p>
      <p></p>
      <h3>Notification Preferences</h3>
      <div class="notification-preferences">
        <div v-for="(item, index) in items" :key="index" class="form-check">
          <input class="form-check-input" type="checkbox" :id="'check-' + index" v-model="item.enabled">
          <label class="form-check-label" :for="'check-' + index">
            {{ item.name }}
          </label>
        </div>
        <button class="btn btn-primary" @click="sendPreferences">Save Preferences</button>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  data() {
    return {
      items: [
        // TEST Example items, replace with dynamic data as needed
        { name: 'Group 1', enabled: false },
        { name: 'Person A', enabled: true },
        // Add more items here
      ],
    };
  },
  methods: {
    sendPreferences() {
      // Send the preferences to the Service Worker
      if (navigator.serviceWorker && navigator.serviceWorker.controller) {
        const preferences = this.items.filter(item => item.enabled).map(item => item.name);
        navigator.serviceWorker.controller.postMessage({
          type: 'update-preferences',
          preferences: preferences,
        });
      }
    },
  },
};
</script>