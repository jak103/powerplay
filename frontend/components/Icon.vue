<script lang="ts" setup>
import {
  mdiAbacus,
  mdiAccountCircle,
  mdiAccountMultiple,
  mdiCalendarMonth,
  mdiChartBarStacked,
  mdiCheckCircle,
  mdiCheckCircleOutline,
  mdiChevronLeft,
  mdiCircleOutline,
  mdiCloseCircle,
  mdiCloseCircleOutline,
  mdiDotsHorizontalCircleOutline,
  mdiHome,
  mdiMedalOutline,
  mdiMessageOutline,
  mdiPencilCircle,
  mdiWeatherNight,
  mdiWhiteBalanceSunny,
} from "@mdi/js";
const MDI: Record<string, string> = {
  mdiAbacus,
  mdiAccountCircle,
  mdiAccountMultiple,
  mdiCalendarMonth,
  mdiChartBarStacked,
  mdiCheckCircle,
  mdiCheckCircleOutline,
  mdiChevronLeft,
  mdiCircleOutline,
  mdiCloseCircle,
  mdiCloseCircleOutline,
  mdiDotsHorizontalCircleOutline,
  mdiHome,
  mdiMedalOutline,
  mdiMessageOutline,
  mdiPencilCircle,
  mdiWeatherNight,
  mdiWhiteBalanceSunny,
};

interface Props {
  size?: `${number}px` | `${number}em`;
  name: string;
}
const props = withDefaults(defineProps<Props>(), {
  size: "1em",
});

const toPascalCase = (input: string) => {
  return input
    .replace(/-([a-z])/g, (_, char) => char.toUpperCase())
    .replace(/^\w/, (char) => char.toUpperCase());
};
const getIcon = computed(() => {
  let name = props.name;
  if (!name.startsWith("mdi")) {
    name = `mdi${toPascalCase(name)}`;
  }
  return MDI[name];
});
</script>

<template>
  <span :style="`height: ${size}; width: ${size};`">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
      <path :d="getIcon" />
    </svg>
  </span>
</template>

<style lang="scss" scoped>
span {
  display: inline-flex;
  vertical-align: middle;
}
:deep(svg) {
  fill: currentColor;
}
</style>
