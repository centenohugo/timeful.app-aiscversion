<template>
  <!-- Fills the router view so the footer sits at the bottom of the viewport
       even when the dashboard is short -->
  <div class="tw-flex tw-min-h-full tw-flex-col">
    <div
      class="tw-mx-auto tw-mt-4 tw-w-full tw-max-w-6xl tw-space-y-4 sm:tw-mt-7"
    >
      <div
        v-if="loading && !eventsNotEmpty"
        class="tw-flex tw-h-[calc(100vh-10rem)] tw-w-full tw-items-center tw-justify-center"
      >
        <v-progress-circular
          indeterminate
          color="primary"
          :size="20"
          :width="2"
        ></v-progress-circular>
      </div>

      <v-fade-transition>
        <Dashboard v-if="!loading || eventsNotEmpty" />
      </v-fade-transition>

    </div>

    <!-- Footer, pushed to the bottom on short pages -->
    <div
      class="tw-mt-auto tw-flex tw-flex-col tw-items-center tw-justify-between tw-pb-24 tw-pt-8 sm:tw-pb-12"
    >
      <router-link
        class="tw-text-xs tw-font-medium tw-text-gray"
        :to="{ name: 'privacy-policy' }"
      >
        Privacy Policy
      </router-link>
    </div>

    <!-- FAB -->
    <BottomFab v-if="isPhone" id="create-event-btn" @click="() => _createNew()">
      <v-icon>mdi-plus</v-icon>
    </BottomFab>
  </div>
</template>

<script>
import EventType from "@/components/EventType.vue"
import BottomFab from "@/components/BottomFab.vue"
import CreateSpeedDial from "@/components/CreateSpeedDial.vue"
import Dashboard from "@/components/home/Dashboard.vue"
import { mapState, mapActions, mapMutations } from "vuex"
import { eventTypes } from "@/constants"
import { isPhone, get } from "@/utils"

export default {
  name: "Home",

  metaInfo: {
    title: "Home - Schedule meetings",
  },

  components: {
    EventType,
    BottomFab,
    CreateSpeedDial,
    Dashboard,
  },

  props: {
    openNewGroup: { type: Boolean, default: false },
  },

  data: () => ({
    loading: true,
  }),

  mounted() {
    this.setNewDialogOptions({
      show: this.openNewGroup,
      openNewGroup: this.openNewGroup,
      eventOnly: false,
    })
  },

  computed: {
    ...mapState(["events", "authUser"]),
    eventsNotEmpty() {
      return this.events.length > 0
    },
    isPhone() {
      return isPhone(this.$vuetify)
    },
  },

  methods: {
    ...mapMutations(["setAuthUser", "setNewDialogOptions"]),
    ...mapActions(["getEvents", "createNew"]),
    userRespondedToEvent(event) {
      return event.hasResponded ?? false
    },
    _createNew() {
      this.createNew({ eventOnly: false })
    },
    createFolder() {},
  },

  created() {
    this.getEvents().then(() => {
      this.loading = false
    })
    get("/user/profile")
      .then((authUser) => {
        this.setAuthUser(authUser)
      })
      .catch(() => {
        this.setAuthUser(null)
      })
  },
}
</script>
