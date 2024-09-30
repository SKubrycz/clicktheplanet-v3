"use client";

import { useState, useRef, useEffect } from "react";
import type { Coords } from "@/components/game_components/GameMain/GameMain";

interface DamageTextProps {
  dmg: string;
  pos: Coords;
  duration: number;
}

export default function DamageText({ dmg, pos, duration }: DamageTextProps) {
  //   const [visible, setVisible] = useState<boolean>(true);
  //   const timeoutRef = useRef<NodeJS.Timeout>();

  //   useEffect(() => {
  //     console.log("DamageText initialized");

  //     timeoutRef.current = setTimeout(() => {
  //       setVisible(false);
  //     }, duration);

  //     return () => clearTimeout(timeoutRef.current);
  //   }, []);

  return (
    <div
      className="damage-text"
      style={{
        top: `${pos.y}px`,
        left: `${pos.x}px`,
        animation: `${duration}ms fadeAway 1 forwards`,
      }}
    >
      {dmg}
    </div>
  );
}
