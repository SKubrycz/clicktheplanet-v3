"use client";

import { useState, useEffect } from "react";

import "./GameMain.scss";
import { Data } from "@/app/game/page";

export interface Health {
  currentHealth: string;
  maxHealth: string;
}

interface GameMainProps {
  data: Data | undefined;
  planetClick(data: string): void;
}

export default function GameMain({ data, planetClick }: GameMainProps) {
  const [health, setHealth] = useState<Health>({
    currentHealth: "10",
    maxHealth: "10",
  });
  const [width, setWidth] = useState<number>(100);

  useEffect(() => {
    if (data) {
      setHealth({
        ...health,
        currentHealth: data.currentHealth,
        maxHealth: data.maxHealth,
      });
      setWidth(data.healthPercent);
    }
  }, [data]);

  const handlePlanetClick = () => {
    // after WebSocket connection -> send click message to the server
    // for now the click mechanic is as below:
    planetClick("click");
    //console.log(width);
    //console.log(health.currentHealth);
  };

  const healthbarStyle = {
    width: `${width}%`,
  };

  return (
    <main className="game-main">
      <div className="main-spacing"></div>
      <div className="main-current-level">Level {data?.currentLevel}</div>
      <div className="main-current-stage">{data?.currentStage}/10</div>
      <div className="main-planet-name">{data?.name}</div>
      <div
        className="main-planet-image"
        onClick={() => handlePlanetClick()}
      ></div>
      {/* Temporary */}
      <div className="main-planet-health-title">Health</div>
      <div className="main-planet-healthbar">
        <div className="main-planet-healthbar-amount">
          {health.currentHealth}/{health.maxHealth}
        </div>
        <div className="main-planet-healthbar-wrapper">
          <div
            className="main-planet-healthbar-above"
            style={healthbarStyle}
          ></div>
          <div className="main-planet-healthbar-under"></div>
        </div>
      </div>
    </main>
  );
}
