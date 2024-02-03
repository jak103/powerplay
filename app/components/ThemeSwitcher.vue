<script lang="ts" setup>
const $q = useQuasar();
const theme = ref<boolean | "auto">("auto");

onMounted(() => {
    theme.value = JSON.parse(localStorage.getItem('powerPlayTheme') || 'auto');
    $q.dark.set(theme.value);
});

const setTheme = () => {
  const setAll = (val: boolean | "auto") => {
    $q.dark.set(val);
    theme.value = val;
    localStorage.setItem("powerPlayTheme", `${val}`);
  };
  if ($q.dark.mode === "auto") {
    setAll(true);
  } else if ($q.dark.mode === true) {
    setAll(false);
  } else {
    setAll("auto");
  }
};
</script>

<template>
  <q-btn
    v-if="theme === 'auto'"
    color="black"
    :icon="mdiThemeLightDark"
    @click="setTheme"
    title="System Theme"
  />
  <q-btn
    v-else-if="theme === true"
    color="black"
    :icon="mdiWeatherNight"
    @click="setTheme"
    title="Dark Theme"
  />
  <q-btn
    v-if="theme === false"
    color="black"
    :icon="mdiWeatherSunny"
    @click="setTheme"
    title="Light Theme"
  />
</template>
