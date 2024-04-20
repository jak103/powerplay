import svgLoader from 'vite-svg-loader'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: false },
  css: ['~/node_modules/bootstrap/scss/bootstrap.scss'],
  experimental: {
    defaults: {
      nuxtLink: {
        externalRelAttribute: 'noopener noreferrer',
        activeClass: 'active',
        exactActiveClass: 'exact-active',
      }
    }
  }
});
