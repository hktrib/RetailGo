"use client";

import { useState, useEffect } from "react";
import { useTheme } from "next-themes";
import { Button } from "@/components/ui/button";
import { Sun, Moon } from "lucide-react";

export default function ThemeSwitcher() {
  const [mounted, setMounted] = useState(false);
  const { theme, setTheme } = useTheme();

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return null;
  }

  return (
    <li className="px-1">
      {theme === "light" ? (
        <Button
          type="button"
          onClick={() => setTheme("dark")}
          variant="ghost"
          className="px-2"
        >
          <Sun className="h-4 w-4" />
        </Button>
      ) : (
        <Button
          type="button"
          onClick={() => setTheme("light")}
          variant="ghost"
          className="px-2 hover:bg-zinc-900"
        >
          <Moon className="h-4 w-4" />
        </Button>
      )}
    </li>
  );
}
