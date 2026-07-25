<template>
  <v-avatar v-if="user" :size="size">
    <img v-if="user.picture" :src="user.picture" referrerpolicy="no-referrer" />
    <div
      v-else
      :class="`tw-flex tw-size-full tw-items-center tw-justify-center tw-bg-[linear-gradient(-25deg,#2b6cb0,#63b3ed,#2b6cb0)] tw-text-${textSize} tw-text-white`"
    >
      {{ user.firstName?.charAt(0) ?? user.email?.charAt(0) ?? "" }}
    </div>
  </v-avatar>
</template>

<script>
import { calendarTypes } from "@/constants"

export default {
  name: "UserAvatarContent",
  props: {
    user: Object,
    size: { type: Number, default: 48 },
  },

  computed: {
    calendarTypes() {
      return calendarTypes
    },
    textSize() {
      return this.size <= 24 ? "xs" : "lg"
    },
  },
}
</script>
