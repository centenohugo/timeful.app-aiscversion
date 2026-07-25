<template>
  <v-img
    :alt="alt"
    class="shrink tw-cursor-pointer"
    contain
    :src="src"
    transition="fade-transition"
    :width="width"
    :height="width"
  />
</template>

<script>
import { isPhone } from "@/utils"

export default {
  name: "Logo",

  props: {
    // Kept for backwards compatibility with existing call sites. Every type now
    // resolves to the AISC Madrid logo; only the rendered size differs.
    type: {
      type: String,
      default: "timeful",
    },
  },

  computed: {
    isPhone() {
      return isPhone(this.$vuetify)
    },
    alt() {
      if (this.type === "betterwhen2meet") {
        return "Betterwhen2meet Logo"
      }

      return "Timeful Logo"
    },
    src() {
      // The AISC logo is a roughly square mark (676x680), unlike the old wide
      // Timeful wordmark, so widths below are much smaller.
      return require("@/assets/aisc-logo.png")
    },
    width() {
      switch (this.type) {
        case "betterwhen2meet":
        case "aprilfools":
          return this.isPhone ? 96 : 128
        case "timeful":
        default:
          return this.isPhone ? 36 : 44
      }
    },
  },
}
</script>
