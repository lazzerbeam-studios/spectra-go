---
name: tailwind-nativewind
description: Tailwind and NativeWind styling for React Native/Expo. Use when styling components, editing theme or global.css, or using className in the app.
---

# Tailwind / NativeWind

**Stack:** Tailwind with **nativewind/preset**, **className** on RN components. Theme via CSS variables in **global.css** and **tailwind.config.js**.

## Setup

* **Config:** **tailwind.config.js** uses **nativewind/preset**, **darkMode: 'class'**, **content** over **src/app**, **src/components**, **src/screens**.
* **Theme:** Colors and radius come from **global.css** **:root** (HSL values without **hsl()**). **tailwind.config.js** **theme.extend** maps **background**, **foreground**, **primary**, **secondary**, **destructive**, **muted**, **accent**, **card**, **popover**, **border**, **input**, **ring** to **hsl(var(--name))**.
* **Utils:** Use **cn()** from **~/src/lib/utils** (clsx + tailwind-merge) to merge **className** values. Use **cva** for component variants.

## Design tokens

* Prefer semantic tokens over raw colors: **text-foreground**, **text-muted-foreground**, **text-primary**, **bg-background**, **bg-card**, **bg-primary**, **border-border**, **bg-secondary/25**, **bg-primary/10**.
* Radius: **rounded-md**, **rounded-lg**, **rounded-2xl**; config has **radius** (lg, md, sm).
* Shadows: **shadow-sm shadow-black/5** for subtle elevation.
* Custom fonts: **font-Quicksand300** … **font-Quicksand700** (defined in theme **fontFamily**).

## Components

* Build variants with **cva**; base classes plus **Platform.select({ web: '...' })** for web-only **hover:**, **focus-visible:**, **aria-invalid:**.
* Pass **className** into **cn(buttonVariants({ variant, size }), className)** so callers can override or extend.
* Put **className** last on the element when passing it through.

## Conventions

* Style with **className**; use **StyleSheet** only when values are dynamic or not expressible in Tailwind.
* **content** must include any directory that contains **className** usage so classes are generated.
