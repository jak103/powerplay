<script lang="ts" setup>
const $q = useQuasar();
const theme = ref<boolean | "auto">("auto");
const showOverlay = ref(true);

onMounted(() => {
  let themeVal: boolean | string =
    localStorage.getItem("powerPlayTheme") || "auto";
  if (themeVal == "true") themeVal = true;
  else if (themeVal == "false") themeVal = false;
  else themeVal = "auto";
  setAll(themeVal as boolean | "auto");
  showOverlay.value = false;
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
  <Teleport to="body">
    <div v-if="showOverlay" class="theme-cover">
      <q-circular-progress
        indeterminate
        rounded
        size="50px"
        color="grey-8"
        class="q-ma-md fade-in"
      />
    </div>
  </Teleport>
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

<style lang="scss" scoped>
.theme-cover {
  position: absolute;
  background: #ffffff;
  top: 0;
  right: 0;
  left: 0;
  bottom: 0;
  z-index: 100000;
  display: flex;
  justify-content: center;
  align-items: center;
  .fade-in {
    opacity: 1;
    animation-name: fadeInOpacity;
    animation-iteration-count: 1;
    animation-timing-function: ease-in;
    animation-duration: 3s;
  }
}
@media (prefers-color-scheme: dark) {
  .theme-cover {
    background: #000000;
  }
}
@keyframes fadeInOpacity {
  0% {
    opacity: 0;
  }
  30% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>
