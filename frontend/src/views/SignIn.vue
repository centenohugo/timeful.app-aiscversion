<template>
  <div
    class="tw-flex tw-min-h-screen tw-items-center tw-justify-center tw-bg-light-gray tw-px-4"
  >
    <div class="tw-w-full tw-max-w-[420px]">
      <!-- Logo -->
      <div class="tw-mb-8 tw-flex tw-justify-center">
        <router-link :to="{ name: 'landing' }">
          <v-img
            alt="Meet AISC Logo"
            class="shrink tw-cursor-pointer"
            contain
            src="@/assets/aisc-logo.png"
            transition="fade-transition"
            width="96"
            height="96"
          />
        </router-link>
      </div>

      <v-card class="tw-rounded-xl tw-px-2 tw-py-4">
        <v-card-title class="tw-flex tw-flex-col tw-items-center tw-pb-0">
          <div class="tw-text-2xl tw-font-medium">
            {{ isSignUp ? "Create an account" : "Welcome back" }}
          </div>
          <div class="tw-mt-1 tw-text-sm tw-font-normal tw-text-dark-gray">
            {{ isSignUp ? "Sign up to get started" : "Sign in to your account" }}
          </div>
        </v-card-title>
        <v-card-text class="tw-flex tw-flex-col tw-items-center tw-pt-6">
          <div class="tw-mb-4 tw-flex tw-w-full tw-flex-col tw-gap-y-2">
            <v-btn
              block
              @click="signIn(calendarTypes.GOOGLE)"
              class="tw-bg-white"
            >
              <div class="tw-flex tw-w-full tw-items-center tw-gap-2">
                <v-img
                  class="tw-flex-initial"
                  width="20"
                  height="20"
                  src="@/assets/google_logo.svg"
                />
                <v-spacer />
                {{ isSignUp ? "Sign up with" : "Continue with" }} Google
                <v-spacer />
              </div>
            </v-btn>
          </div>
          <div class="tw-text-center tw-text-xs">
            By continuing, you agree to our
            <router-link class="tw-text-blue" :to="{ name: 'privacy-policy' }">
              privacy policy
            </router-link>
          </div>
        </v-card-text>
      </v-card>

      <div
        class="tw-mt-4 tw-rounded-xl tw-bg-light-gray-stroke/50 tw-py-4 tw-text-center tw-text-sm tw-text-dark-gray"
      >
        <template v-if="isSignUp">
          Already have an account?
          <router-link
            class="tw-font-medium tw-text-green"
            :to="{ name: 'sign-in', query: $route.query }"
            >Log in</router-link
          >
        </template>
        <template v-else>
          Don't have an account?
          <router-link
            class="tw-font-medium tw-text-green"
            :to="{ name: 'sign-up', query: $route.query }"
            >Sign up</router-link
          >
        </template>
      </div>
    </div>
  </div>
</template>

<script>
import { calendarTypes } from "@/constants"
import { signInGoogle } from "@/utils"
import { mapMutations } from "vuex"
import Logo from "@/components/Logo.vue"

export default {
  name: "SignIn",

  props: {
    initialIsSignUp: { type: Boolean, default: false },
  },

  metaInfo() {
    return {
      title: this.isSignUp
        ? "Sign Up - Schedule meetings"
        : "Sign In - Schedule meetings",
    }
  },

  components: {
    Logo,
  },

  data() {
    return {
      calendarTypes,
      isSignUp: this.initialIsSignUp,
    }
  },

  methods: {
    ...mapMutations(["setAuthUser"]),
    signIn(provider) {
      signInGoogle({ state: null, selectAccount: true })
    },
  },
}
</script>
