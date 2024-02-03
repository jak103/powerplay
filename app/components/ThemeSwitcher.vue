<script lang="ts" setup>
const $q = useQuasar();
const theme = ref<boolean | "auto">("auto");

onMounted(() => {
    let themeVal: boolean | string = localStorage.getItem('powerPlayTheme') || 'auto';
    if(themeVal == 'true') themeVal = true;
    else if (themeVal == 'false') themeVal = false;
    else themeVal = 'auto'
    setAll(themeVal as boolean | 'auto');
});

const setAll = (val: boolean | "auto") => {
  $q.dark.set(val);
  theme.value = val;
  localStorage.setItem("powerPlayTheme", `${val}`);
};

const setTheme = () => {
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
    flat
    fab-mini
  />
  <q-btn
    v-else-if="theme === true"
    color="black"
    :icon="mdiWeatherNight"
    @click="setTheme"
    title="Dark Theme"
    flat
    fab-mini
  />
  <q-btn
    v-if="theme === false"
    color="black"
    :icon="mdiWeatherSunny"
    @click="setTheme"
    title="Light Theme"
    flat
    fab-mini
  />
</template>
