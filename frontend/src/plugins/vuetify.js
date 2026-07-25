import Vue from "vue"
import Vuetify from "vuetify/lib"
import tailwind from "../../tailwind.config"

Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    themes: {
      light: {
        // `colors.green` / `colors.blue` are historical token names that now hold
        // the AISC pink / cyan brand values — see tailwind.config.js.
        primary: tailwind.theme.colors.green, // AISC pink #EB178E
        secondary: tailwind.theme.colors["light-blue"], // AISC cyan #20CCF1
        error: tailwind.theme.colors.red,
      },
    },
  },
  breakpoint: {
    thresholds: {
      xs: 640,
      sm: 768,
      md: 1024,
      lg: 1280,
    },
    scrollBarWidth: 0,
  },
})
