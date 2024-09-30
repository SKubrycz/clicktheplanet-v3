"use client";

import { DamageDone } from "@/lib/game/gameSlice";

interface Coords {
  x: number;
  y: number;
}

interface DamageTextProps {
  damageDone: DamageDone;
  pos: Coords;
  duration: number;
}

export default function DamageText({
  damageDone,
  pos,
  duration,
}: DamageTextProps) {
  return (
    <div
      className={
        damageDone.critical
          ? "damage-text damage-text-critical"
          : "damage-text damage-text-normal"
      }
      style={{
        top: `${pos.y}px`,
        left: `${pos.x}px`,
        animation: `${duration}ms fadeAway 1 forwards`,
      }}
    >
      {damageDone.damage}
    </div>
  );
}
