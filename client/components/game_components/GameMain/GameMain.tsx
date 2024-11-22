"use client";

import { useState, useEffect, useRef } from "react";

import { useAppSelector } from "@/lib/hooks";

import type { DamageDone } from "@/lib/game/gameSlice";

import "./GameMain.scss";
import DamageText from "./DamageText";
import Planet from "./Planet";

export interface Health {
  currentHealth: string;
  maxHealth: string;
}

interface GameMainProps {
  planetClick(data: string): void;
}

interface DmgText {
  damageDone: DamageDone;
  x: number;
  y: number;
}

export default function GameMain({ planetClick }: GameMainProps) {
  const gameData = useAppSelector((state) => state.game);
  const planetData = useAppSelector((state) => state.planet);

  const [dmgTextArr, setDmgTextArr] = useState<DmgText[]>([]);
  const planetRef = useRef<HTMLCanvasElement>(null);

  const duration: number = 700;

  useEffect(() => {
    const timeout = setTimeout(() => {
      if (dmgTextArr.length >= 1 && dmgTextArr)
        setDmgTextArr((prev) => prev.slice(0, -1));
    }, duration);

    return () => clearTimeout(timeout);
  }, [dmgTextArr]);

  const handlePlanetClick = (e: React.MouseEvent<HTMLCanvasElement>) => {
    planetClick("click");
    if (e.target === planetRef.current) {
      setDmgTextArr([
        ...dmgTextArr,
        {
          damageDone: {
            damage: gameData?.damageDone.damage,
            critical: gameData?.damageDone.critical,
          },
          x: e.clientX,
          y: e.clientY,
        },
      ]);
    }
  };

  const healthbarStyle = {
    width: `${planetData?.healthPercent}%`,
    background: gameData?.isBoss
      ? "radial-gradient(ellipse at bottom, rgba(124, 29, 179, 0.6), rgba(188, 0, 255, 0.5))"
      : "radial-gradient(ellipse at bottom, rgba(0, 106, 192, 0.6), rgba(42, 127, 202, 0.5))",
  };

  return (
    <main className="game-main">
      <div className="main-spacing"></div>
      <div className="main-current-level">Level {gameData?.currentLevel}</div>
      <div className="main-current-stage">{gameData?.currentStage}/10</div>
      <div className="main-planet-name-title">Planet name:</div>
      <div className="main-planet-name">
        {planetData?.diamondPlanet?.isDiamondPlanet
          ? "Diamond Planet"
          : planetData?.planetName}
      </div>
      <div className="main-planet-image-wrapper">
        <Planet
          planetRef={planetRef}
          click={(e) => handlePlanetClick(e)}
        ></Planet>
      </div>
      <div className="main-planet-health-title">Health</div>
      <div className="main-planet-healthbar">
        <div className="main-planet-healthbar-amount">
          {planetData?.currentHealth}/{planetData?.maxHealth}
        </div>
        <div className="main-planet-healthbar-wrapper">
          <div
            className="main-planet-healthbar-above"
            style={healthbarStyle}
          ></div>
          <div className="main-planet-healthbar-under"></div>
        </div>
      </div>
      <>
        {dmgTextArr?.map((el, i) => {
          return (
            <DamageText
              key={i}
              damageDone={el.damageDone}
              pos={{
                x: el.x,
                y: el.y,
              }}
              duration={duration}
            ></DamageText>
          );
        })}
      </>
    </main>
  );
}
