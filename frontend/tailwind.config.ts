import type { Config } from "tailwindcss";

export default {
  darkMode: ["class"],
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    container: {
      center: true,
      padding: "1rem",
      screens: {
        "2xl": "1280px",
      },
    },
    extend: {
      colors: {
        border: "hsl(144 14% 82%)",
        input: "hsl(145 22% 96%)",
        ring: "hsl(142 52% 37%)",
        background: "hsl(45 35% 98%)",
        foreground: "hsl(145 22% 16%)",
        primary: {
          DEFAULT: "hsl(142 52% 37%)",
          foreground: "hsl(47 100% 97%)",
        },
        secondary: {
          DEFAULT: "hsl(44 58% 90%)",
          foreground: "hsl(145 27% 22%)",
        },
        muted: {
          DEFAULT: "hsl(142 24% 93%)",
          foreground: "hsl(150 8% 38%)",
        },
        accent: {
          DEFAULT: "hsl(31 86% 88%)",
          foreground: "hsl(20 42% 26%)",
        },
        destructive: {
          DEFAULT: "hsl(0 72% 51%)",
          foreground: "hsl(0 0% 100%)",
        },
        card: {
          DEFAULT: "hsl(48 100% 99%)",
          foreground: "hsl(145 22% 16%)",
        },
      },
      borderRadius: {
        lg: "0.9rem",
        md: "0.7rem",
        sm: "0.5rem",
      },
      boxShadow: {
        soft: "0 10px 30px -16px rgba(15, 64, 36, 0.28)",
      },
      keyframes: {
        "fade-up": {
          from: { opacity: "0", transform: "translateY(10px)" },
          to: { opacity: "1", transform: "translateY(0)" },
        },
      },
      animation: {
        "fade-up": "fade-up 240ms ease-out",
      },
    },
  },
  plugins: [],
} satisfies Config;

