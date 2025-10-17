"use client";

import { useTheme } from "@/provider/useTheme";

export default function Home() {
  const { toggleTheme } = useTheme();

  return (
    <div className="p-5">
      <div>
        <p className="text-lg font-jet font-semibold">Hello, World!</p>
        <button
          className="font-jet border border-foreground px-2 py-1 rounded-md mt-2 cursor-pointer hover:bg-foreground/10 transition-colors"
          onClick={toggleTheme}
        >
          switch theme
        </button>
      </div>
    </div>
  );
}
