<template>
  <div class="tw-bg-white">
    <div
      class="tw-relative tw-m-auto tw-mb-12 tw-flex tw-max-w-6xl tw-flex-col tw-px-4 tw-pt-10 sm:tw-mb-20 sm:tw-pt-16"
    >
      <div class="tw-flex tw-flex-col tw-items-center">
        <!-- Logo, sitting quietly above the hero -->
        <div class="tw-mb-6 tw-opacity-90">
          <Logo type="timeful" />
        </div>

        <div
          class="tw-mb-6 tw-flex tw-max-w-[26rem] tw-flex-col tw-items-center sm:tw-w-[35rem] sm:tw-max-w-none"
        >
          <div
            id="header"
            class="tw-mb-4 tw-text-center tw-text-2xl tw-font-medium sm:tw-text-4xl lg:tw-text-4xl xl:tw-text-5xl"
          >
            <h1>Find a time to meet</h1>
          </div>

          <div
            class="lg:tw-text-md tw-text-left tw-text-center tw-text-sm tw-text-very-dark-gray sm:tw-text-lg md:tw-text-lg xl:tw-text-lg"
          >
            Coordinate group meetings without the back and forth.
            <br class="tw-hidden sm:tw-block" />
            Integrates with your calendar.
          </div>
        </div>

        <div class="tw-space-y-2">
          <v-btn
            class="tw-block tw-self-center tw-rounded-lg tw-bg-green tw-px-10 tw-text-base sm:tw-px-10 lg:tw-px-12"
            dark
            @click="authUser ? openDashboard() : (newDialog = true)"
            large
            :x-large="$vuetify.breakpoint.mdAndUp"
          >
            {{ authUser ? "Open dashboard" : "Create event" }}
          </v-btn>
        </div>
      </div>
    </div>

    <!-- Sign in dialog -->
    <SignInDialog v-model="signInDialog" @signIn="_signIn" />

    <!-- New event dialog -->
    <NewDialog v-model="newDialog" no-tabs @signIn="signIn" />
  </div>
</template>

<style scoped>
@media screen and (min-width: 375px) and (max-width: 640px) {
  #header {
    font-size: 1.875rem !important; /* 30px */
    line-height: 2.25rem !important; /* 36px */
  }
}
</style>

<script>
import { isPhone, signInGoogle } from "@/utils"
import NewDialog from "@/components/NewDialog.vue"
import Logo from "@/components/Logo.vue"
import SignInDialog from "@/components/SignInDialog.vue"
import { calendarTypes } from "@/constants"
import { mapState, mapMutations } from "vuex"

export default {
  name: "Landing",

  metaInfo: {
    title: "Schedule meetings",
  },

  components: {
    NewDialog,
    Logo,
    SignInDialog,
  },

  data: () => ({
    signInDialog: false,
    newDialog: false,
  }),

  computed: {
    ...mapState(["authUser"]),
    isPhone() {
      return isPhone(this.$vuetify)
    },
  },

  methods: {
    ...mapMutations(["setAuthUser"]),
    _signIn() {
      signInGoogle({ state: null, selectAccount: true })
    },
    signIn() {
      this.$router.push({ name: "sign-in" })
    },
    openDashboard() {
      this.$router.push({ name: "home" })
    },
  },
}
</script>
