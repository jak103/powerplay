<script lang="ts" setup>
const theme = ref<null | string>(null);
onMounted(() => {
  theme.value = localStorage.getItem("theme") || "light";
});
const toggleTheme = () => {
  theme.value = theme.value === "light" ? "dark" : "light";
  localStorage.setItem("theme", theme.value);
  if (theme.value !== "dark") {
    document.body.removeAttribute("data-bs-theme");
  } else {
    document.body.setAttribute("data-bs-theme", "dark");
  }
};
</script>

<template>
  <button v-if="theme" class="btn btn-primary" @click="toggleTheme">
    <Icon v-if="theme === 'light'" name="mdiWhiteBalanceSunny" size="1.6em" />
    <Icon v-else name="mdiWeatherNight" size="1.6em" />
    Theme
  </button>
</template>
