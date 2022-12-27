/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,jsx,tsx,js,ts}"],
  theme: {
    extend: {},
  },
  plugins: [require("tw-elements/dist/plugin")],
  darkMode: "class",
};
