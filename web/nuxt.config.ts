// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: {enabled: true},
  ssr: false,
  app: {
    head: {
      meta: [
        {name: 'viewport', content: 'width=device-width, initial-scale=1'}
      ],
      script: [
        {src: '//embed.typeform.com/next/embed.js'}
      ],
    }
  },
  css: [
    '~/assets/css/main.scss'
  ]
})
