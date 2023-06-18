/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{js,ts,jsx,tsx}",
    "./src/**/**/*.{js,ts,jsx,tsx}",
    "./src/**/**/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      listStyleType: {
        circle: "circle",
        square: "square",
      },
      colors: {
        "light-black": "#333333",
      },
    },
  },
  plugins: [],
};
