<template>
  <footer class="tw-bg-green tw-p-8 tw-text-white">
    <div class="tw-mx-auto tw-max-w-7xl">
      <div class="tw-mb-3 tw-flex tw-items-center">
        <v-icon class="tw-mr-1 tw-text-white">mdi-calendar-blank</v-icon>
        <span class="tw-text-2xl tw-font-medium">timeful</span>
      </div>
      <router-link :to="{ name: 'privacy-policy' }" class="tw-text-sm">
        Privacy Policy
      </router-link>
    </div>
  </footer>
</template>

<style scoped>
a {
  @apply tw-text-white/80 !important;
}
a:hover {
  @apply tw-text-white !important;
}
</style>

<script>
export default {
  name: "Footer",
}
</script>
