/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./TailsOfOld/static/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("tw-elements/dist/plugin.cjs")],
  darkMode: "class"
}