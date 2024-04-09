<script lang="ts" setup>
const toggleTheme = () => {
  useTheme().value = useTheme().value === "light" ? "dark" : "light";
  localStorage.setItem("theme", useTheme().value as string);
};
</script>

<template>
  <div>
    <button class="btn btn-primary" @click="toggleTheme">
      <Icon v-if="useTheme().value === 'light'" name="mdiWhiteBalanceSunny" size="1.6em" />
      <Icon v-else name="mdiWeatherNight" size="1.6em" />
      Theme
    </button>
  </div>
</template>
