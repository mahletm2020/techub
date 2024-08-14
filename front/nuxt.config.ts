// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  plugins: ['~/plugins/apollo-client.ts'],
  runtimeConfig: {
    public: {
      HASURA_ADMIN_SECRET: process.env.HASURA_ADMIN_SECRET,
    },
  },
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true }
})
