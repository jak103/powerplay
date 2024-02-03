// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["nuxt-quasar-ui"],
  quasar: {
    components: {
      defaults: {
        QBtn: {
          color: 'primary'
        }
      }
    },
    extras: {
      svgIcons: ['mdi-v7'],
    },
    iconSet: "mdi-v7",
    plugins: ['Dark'],
    sassVariables: "~/assets/pp-variables.scss",
  },
});
