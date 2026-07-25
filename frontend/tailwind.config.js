const colors = require("tailwindcss/colors")

module.exports = {
  content: [
    "./index.html",
    "./public/**/*.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  important: true,
  theme: {
    extend: {
      fontSize: {
        xs: ["0.813rem", "1rem"],
      },
    },
    /**
     * AISC Madrid brand palette (see frontend/AESTHETIC.md).
     *
     * IMPORTANT: the `*-green` and `*-blue` token names below are HISTORICAL —
     * they date from the original Timeful green/blue branding and are referenced
     * by `tw-bg-green`, `tw-text-green`, `tw-bg-blue`, ... across ~27 files.
     * They were deliberately NOT renamed; only their hex values were remapped so
     * the whole UI re-themes from this one place. Read them as:
     *
     *   green family  -> AISC pink   (#EB178E, PRIMARY action color) + its shades
     *   blue family   -> AISC cyan   (#20CCF1, SECONDARY) + a darkened,
     *                    text-readable variant used where white text sits on it
     *
     * `avail-green`, `red`, `orange`, `yellow` are SEMANTIC (availability heat
     * scale, errors, warnings) and are intentionally left un-branded.
     */
    colors: {
      transparent: "transparent",
      current: "currentColor",
      // --- AISC pink family (historically "green") ---
      "pale-green": "#FAD3EA", // pastel pink (card/row tint)
      "light-green": "#F45BAE", // light pink (accent borders)
      "ligher-green": "#FCE4F3", // pink tint background
      green: "#EB178E", // AISC pink — PRIMARY action color
      "dark-green": "#D43089", // pink hover/pressed
      "darkest-green": "#C4106F", // deepest pink
      // --- AISC cyan family (historically "blue") ---
      "light-blue": "#20CCF1", // AISC cyan — SECONDARY
      blue: "#0B7E99", // darkened cyan (readable as text on white / white text on it)
      orange: "#E5A800",
      yellow: "#FFE8B8",
      "dark-yellow": "#997700",
      white: "#FFFFFF",
      "off-white": "#F2F2F2",
      black: "#000000",
      gray: "#BDBDBD",
      "dark-gray": "#6B6B6B",
      "very-dark-gray": "#4F4F4F",
      "light-gray": "#f3f4f6",
      "light-gray-stroke": "#dfdfdf",
      // Semantic, NOT branded: encodes how many people are available, so it must
      // stay a distinct, legible light->dark scale independent of the brand pink.
      "avail-green": colors.emerald, // The green used for marking availability
      red: "#DB1616",
    },
    screens: {
      sm: "640px",
      md: "768px",
      mdlg: "896px",
      lg: "1024px",
      xl: "1280px",
      "2xl": "1536px",
    },
  },
  plugins: [],
  prefix: "tw-",
  safelist: [],
}
