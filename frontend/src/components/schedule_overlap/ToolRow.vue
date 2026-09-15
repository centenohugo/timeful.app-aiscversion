<template>
  <div>
    <div
      class="tw-flex tw-min-h-[5rem] tw-flex-1 tw-items-center tw-justify-center tw-text-sm sm:tw-mt-0 sm:tw-justify-between"
    >
      <div
        :class="
          state === states.EDIT_AVAILABILITY
            ? 'tw-justify-center'
            : 'tw-justify-between'
        "
        class="tw-flex tw-flex-1 tw-flex-wrap tw-gap-x-4 tw-gap-y-2 tw-py-4 sm:tw-justify-start sm:tw-gap-x-4"
      >
        <!-- Select timezone -->
        <div v-if="!event.daysOnly" class="tw-flex tw-items-center tw-gap-2">
          <TimezoneSelector
            ref="timezoneSelector"
            class="tw-w-full sm:tw-w-[unset]"
            label="Times shown in"
            :value="curTimezone"
            :reference-date="timezoneReferenceDate"
            @input="(val) => $emit('update:curTimezone', val)"
          />
          <v-select
            :value="timeType"
            @input="$emit('update:timeType', $event)"
            :items="timeTypeOptions"
            :menu-props="{ auto: true }"
            item-text="label"
            item-value="value"
            class="tw-z-20 -tw-mt-px tw-w-16 tw-text-sm"
            dense
            hide-details
          />
        </div>
        <div
          v-if="isPhone && !event.daysOnly"
          class="tw-flex tw-basis-full tw-items-center tw-gap-x-2 tw-py-4"
        >
          Show
          <v-select
            :value="mobileNumDays"
            @input="$emit('update:mobileNumDays', $event)"
            :items="mobileNumDaysOptions"
            :menu-props="{ auto: true }"
            item-text="label"
            item-value="value"
            class="-tw-mt-px tw-flex-none tw-shrink tw-basis-24 tw-text-sm"
            dense
            hide-details
          />
          at a time
        </div>

        <template v-if="state !== states.EDIT_AVAILABILITY && isPhone">
          <EventOptions
            class="tw-mt-2 tw-w-full"
            :event="event"
            :showBestTimes="showBestTimes"
            @update:showBestTimes="(val) => $emit('update:showBestTimes', val)"
            :hideIfNeeded="hideIfNeeded"
            @update:hideIfNeeded="(val) => $emit('update:hideIfNeeded', val)"
            :showEventOptions="showEventOptions"
            @toggleShowEventOptions="$emit('toggleShowEventOptions')"
            :startCalendarOnMonday="startCalendarOnMonday"
            @update:startCalendarOnMonday="
              (val) => $emit('update:startCalendarOnMonday', val)
            "
            :numResponses="numResponses"
          />
        </template>
        <template
          v-if="state === states.EDIT_AVAILABILITY && isWeekly && !isPhone"
        >
          <v-spacer />
          <div class="tw-min-w-fit">
            <GCalWeekSelector
              v-if="calendarPermissionGranted"
              :week-offset="weekOffset"
              :event="event"
              @update:weekOffset="(val) => $emit('update:weekOffset', val)"
              :start-on-monday="event.startOnMonday"
            />
          </div>
        </template>
      </div>

      <div
        v-if="showScheduleEventButton"
        style="width: 181.5px"
        class="tw-hidden sm:tw-flex"
      >
        <template v-if="state !== states.SCHEDULE_EVENT">
          <v-btn
            outlined
            class="tw-w-full tw-text-blue"
            @click="(e) => $emit('scheduleEvent', e)"
          >
            <v-icon small>mdi-calendar-check</v-icon>
            <span class="tw-ml-2">Schedule event</span>
          </v-btn>
        </template>
        <template v-else>
          <v-btn
            outlined
            class="tw-mr-1 tw-text-red"
            @click="(e) => $emit('cancelScheduleEvent', e)"
          >
            Cancel
          </v-btn>
          <v-btn
            :disabled="!allowScheduleEvent || schedulingEvent"
            :loading="schedulingEvent"
            class="tw-bg-blue tw-text-white"
            @click="(e) => $emit('confirmScheduleEvent', true)"
          >
            Schedule
          </v-btn>
        </template>
      </div>
    </div>

    <!-- Explains which timezone the times are in -->
    <div
      v-if="timezoneNotice"
      class="-tw-mt-2 tw-mb-3 tw-flex tw-flex-wrap tw-items-center tw-justify-center tw-gap-x-2 tw-gap-y-1 tw-text-xs tw-text-very-dark-gray sm:tw-justify-start"
    >
      <v-icon x-small>mdi-earth</v-icon>
      <span>{{ timezoneNotice.text }}</span>
      <button
        class="tw-font-medium tw-text-green hover:tw-underline"
        @click="switchTimezone(timezoneNotice.target)"
      >
        {{ timezoneNotice.action }}
      </button>
    </div>
  </div>
</template>

<script>
import TimezoneSelector from "./TimezoneSelector.vue"
import GCalWeekSelector from "./GCalWeekSelector.vue"
import {
  isPhone,
  canScheduleEvent,
  getTimezoneOption,
  getLocalTimezoneOption,
  getLocalTimezoneCity,
} from "@/utils"
import ExpandableSection from "../ExpandableSection.vue"
import EventOptions from "./EventOptions.vue"
import { timeTypes } from "@/constants"
import { mapState } from "vuex"
import dayjs from "dayjs"
import utcPlugin from "dayjs/plugin/utc"
import timezonePlugin from "dayjs/plugin/timezone"
dayjs.extend(utcPlugin)
dayjs.extend(timezonePlugin)

export default {
  name: "ToolRow",

  props: {
    event: { type: Object, required: true },
    state: { type: String, required: true },
    states: { type: Object, required: true },
    curTimezone: { type: Object, required: true },
    startCalendarOnMonday: { type: Boolean, default: false },
    showBestTimes: { type: Boolean, required: true },
    hideIfNeeded: { type: Boolean, required: true },
    isWeekly: { type: Boolean, required: true },
    calendarPermissionGranted: { type: Boolean, required: true },
    weekOffset: { type: Number, required: true },
    timezoneReferenceDate: { type: Date, required: false, default: null },
    numResponses: { type: Number, required: true },
    mobileNumDays: { type: Number, default: 3 }, // The number of days to show at a time on mobile
    allowScheduleEvent: { type: Boolean, required: true },
    schedulingEvent: { type: Boolean, default: false },
    showEventOptions: { type: Boolean, required: true },
    timeType: { type: String, required: true },
  },

  components: {
    TimezoneSelector,
    GCalWeekSelector,
    ExpandableSection,
    EventOptions,
  },

  data: () => ({
    mobileNumDaysOptions: [
      { label: "3 days", value: 3 },
      { label: "7 days", value: 7 },
    ],
    timeTypeOptions: [
      { label: "12h", value: timeTypes.HOUR12 },
      { label: "24h", value: timeTypes.HOUR24 },
    ],
  }),

  computed: {
    ...mapState(["authUser"]),
    isPhone() {
      return isPhone(this.$vuetify)
    },
    canSchedule() {
      return canScheduleEvent(this.event, this.authUser?._id)
    },
    /**
     * Tells viewers whether the times are in their own timezone and, if the organizer
     * picked the times in another one, lets them switch between both.
     * Returns { text, action, target } or null when there's nothing worth saying
     */
    timezoneNotice() {
      if (this.event.daysOnly || !this.curTimezone?.value) return null

      const referenceDate = this.timezoneReferenceDate ?? new Date()
      const offsetOf = (value) => dayjs(referenceDate).tz(value).utcOffset()

      const local = getLocalTimezoneOption(referenceDate)
      if (!local) return null
      const city = getLocalTimezoneCity()
      const curOffset = offsetOf(this.curTimezone.value)
      const viewingLocal = curOffset === local.offset

      const organizer =
        this.event.timezone &&
        getTimezoneOption(this.event.timezone, referenceDate)
      const organizerIsElsewhere = organizer && organizer.offset !== local.offset

      if (!organizerIsElsewhere) {
        if (viewingLocal) return null
        return {
          text: `These times aren't in your time zone (${city}).`,
          action: "Show in my time zone",
          target: "local",
        }
      }

      const organizerName = `${organizer.gmtString} ${organizer.label}`
      if (viewingLocal) {
        return {
          text: `Shown in your time zone (${city}). The organizer set this event in ${organizerName}.`,
          action: "See the organizer's times",
          target: organizer,
        }
      }
      if (curOffset === organizer.offset) {
        return {
          text: `Shown in the organizer's time zone, not yours (${city}).`,
          action: "Show in my time zone",
          target: "local",
        }
      }
      return {
        text: `These times aren't in your time zone (${city}). The organizer set this event in ${organizerName}.`,
        action: "Show in my time zone",
        target: "local",
      }
    },
    showScheduleEventButton() {
      return (
        !this.event.daysOnly &&
        this.numResponses > 0 &&
        this.state !== this.states.EDIT_AVAILABILITY &&
        this.canSchedule
      )
    },
  },

  methods: {
    /** Switches to "local" (also forgetting the saved preference) or to the given timezone for this visit */
    switchTimezone(target) {
      if (target === "local") {
        this.$refs.timezoneSelector.resetTimezone()
      } else {
        this.$emit("update:curTimezone", target)
      }
    },
  },
}
</script>
