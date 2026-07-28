<template>
  <footer class="aisc-footer tw-p-8 tw-text-very-dark-gray">
    <div class="tw-mx-auto tw-max-w-7xl">
      <div class="tw-mb-3 tw-flex tw-items-center">
        <v-img
          alt="Meet AISC Logo"
          class="shrink"
          contain
          src="@/assets/aisc-logo.png"
          transition="fade-transition"
          width="48"
          height="48"
          max-width="48"
        />
      </div>
      <router-link :to="{ name: 'privacy-policy' }" class="tw-text-sm">
        Privacy Policy
      </router-link>
    </div>
  </footer>
</template>

<style scoped>
/* AISC footer gradient: near-white -> pink tint -> full brand pink (left to
   right). All footer content is left-aligned, so it sits on the light end of
   the gradient and stays dark-on-light / readable. */
.aisc-footer {
  background-image: linear-gradient(
    to right,
    #fafafa 0%,
    #fce4f3 45%,
    #eb178e 100%
  );
}
a {
  @apply tw-text-darkest-green !important;
}
a:hover {
  @apply tw-text-black !important;
}
</style>

<script>
export default {
  name: "Footer",
}
</script>
