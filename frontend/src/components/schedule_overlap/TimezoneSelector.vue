<!-- Allows user to change timezone -->
<template>
  <div
    class="tw-flex tw-items-center tw-justify-center"
    id="timezone-select-container"
  >
    <div :class="`tw-mr-2 tw-mt-px ${labelColor}`">{{ label }}</div>
    <v-select
      id="timezone-select"
      :value="value"
      @input="onChange"
      :items="timezones"
      :menu-props="{ auto: true }"
      class="tw-z-20 -tw-mt-px tw-w-52 tw-text-sm"
      dense
      color="#219653"
      item-color="green"
      hide-details
      item-text="label"
      return-object
    >
      <template v-slot:item="{ item, on, attrs }">
        <v-list-item v-bind="attrs" v-on="on">
          <v-list-item-content>
            <v-list-item-title>
              {{ item.gmtString }} {{ item.label }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
      <template v-slot:selection="{ item }">
        <div class="v-select__selection v-select__selection--comma">
          {{ item.gmtString }} {{ item.label }}
        </div>
      </template>
    </v-select>
    <v-btn v-if="timezoneModified" @click="resetTimezone" icon color="primary"
      ><v-icon>mdi-refresh</v-icon></v-btn
    >
  </div>
</template>

<script>
import { getLocalTimezoneOption, getTimezoneOptions } from "@/utils"

export default {
  name: "TimezoneSelector",

  props: {
    value: { type: Object, required: true },
    label: { type: String, default: "Shown in" },
    labelColor: { type: String, default: "" },
    referenceDate: { type: Date, default: null },
    // Whether the selection is remembered in localStorage (the viewing preference).
    // Turned off where the timezone belongs to something else, e.g. the event form
    persist: { type: Boolean, default: true },
  },

  created() {
    if (this.value.value) return // Timezone has already been set

    // Set timezone to localstorage timezone if localstorage is set
    if (this.persist && localStorage["timezone"]) {
      this.$emit("input", JSON.parse(localStorage["timezone"]))
      return
    }

    // Otherwise, set timezone to local timezone
    this.$emit("input", this.getLocalTimezone())
  },

  data() {
    return {
      // Whether a timezone has been saved to localStorage
      hasStoredTimezone: this.persist && !!localStorage["timezone"],
    }
  },

  computed: {
    effectiveReferenceDate() {
      return this.referenceDate ?? new Date()
    },
    /** Returns an array of all supported timezones */
    timezones() {
      return getTimezoneOptions(this.effectiveReferenceDate)
    },
    localTimezone() {
      return getLocalTimezoneOption(this.effectiveReferenceDate, this.timezones)
    },
    /** Whether the timezone has been modified from the local timezone */
    timezoneModified() {
      return (
        this.hasStoredTimezone ||
        (!!this.value?.value &&
          !!this.localTimezone &&
          this.value.offset !== this.localTimezone.offset)
      )
    },
  },

  methods: {
    /** Updates local storage and emits the new timezone */
    onChange(val) {
      if (this.persist) {
        localStorage["timezone"] = JSON.stringify(val)
        this.hasStoredTimezone = true
      }
      this.$emit("input", val)
    },
    /** Returns a timezone object for the local timezone */
    getLocalTimezone() {
      return this.localTimezone
    },
    /** Resets timezone to the local timezone and clears localstorage as well */
    resetTimezone() {
      this.$emit("input", this.getLocalTimezone())
      if (this.persist) {
        localStorage.removeItem("timezone")
      }
      this.hasStoredTimezone = false
    },
  },

  watch: {
    referenceDate() {
      if (!this.value?.value) {
        return
      }

      const refreshedTimezone = this.timezones.find(
        (timezone) => timezone.value === this.value.value
      )

      if (!refreshedTimezone || refreshedTimezone.offset === this.value.offset) {
        return
      }

      if (this.persist && localStorage["timezone"]) {
        localStorage["timezone"] = JSON.stringify(refreshedTimezone)
      }

      this.$emit("input", refreshedTimezone)
    },
  },
}
</script>
