<template>
  <SlideToggle
    :value="value"
    @input="(e) => $emit('input', e)"
    :options="options"
  >
    <!-- Each option shows a swatch with the exact color it paints -->
    <template #option-available="{ option }">
      <span :class="[swatchClass, option.swatchClass]"></span>
      <span class="tw-line-clamp-1">{{ option.text }}</span>
    </template>
    <template #option-if_needed="{ option }">
      <span :class="[swatchClass, option.swatchClass]"></span>
      <span class="tw-line-clamp-1">{{ option.text }}</span>
    </template>
  </SlideToggle>
</template>

<script>
import { availabilityTypes } from "@/constants"
import SlideToggle from "@/components/SlideToggle.vue"

export default {
  name: "AvailabilityTypeToggle",

  props: {
    value: { type: String, required: true },
  },

  components: { SlideToggle },

  data() {
    return {
      swatchClass:
        "tw-h-3 tw-w-3 tw-flex-none tw-rounded-sm tw-border tw-border-gray",
      options: [
        {
          text: "Available",
          activeClass: "tw-text-available-dark tw-bg-available/5",
          borderClass: "tw-border-available",
          borderStyle: { boxShadow: "0px 2px 8px 0px #00994C40" },
          swatchClass: "tw-bg-available/[.47]",
          value: availabilityTypes.AVAILABLE,
        },
        {
          text: "If needed",
          activeClass: "tw-text-dark-yellow tw-bg-yellow/5",
          borderClass: "tw-border-orange",
          borderStyle: { boxShadow: "0px 2px 8px 0px #FEDB9340" },
          swatchClass: "tw-bg-yellow",
          value: availabilityTypes.IF_NEEDED,
        },
      ],
    }
  },
}
</script>
