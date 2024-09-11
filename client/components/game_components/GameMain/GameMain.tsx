"use client";

import { useState, useEffect } from "react";

import "./GameMain.scss";

export interface Health {
  currentHealth: number;
  maxHealth: number;
}

interface GameMainProps {
  planetClick(data: string): void;
}

export default function GameMain({ planetClick }: GameMainProps) {
  const [health, setHealth] = useState<Health>({
    currentHealth: 10,
    maxHealth: 10,
  });
  const [width, setWidth] = useState<number>(100);

  useEffect(() => {
    /*     const timeout = setTimeout(() => {
      handlePlanetClick();
    }, 1000);
    return () => clearTimeout(timeout); */
    setWidth((health.currentHealth * 100) / health.maxHealth);
  }, [health.currentHealth]);

  const handlePlanetClick = () => {
    // after WebSocket connection -> send click message to the server
    // for now the click mechanic is as below:
    setHealth({
      ...health,
      currentHealth: health.currentHealth - 1,
    });
    planetClick("click");
    if (health.currentHealth <= 0) {
      setHealth({
        ...health,
        currentHealth: health.maxHealth,
      });
    }
    //console.log(width);
    //console.log(health.currentHealth);
  };

  const healthbarStyle = {
    width: `${width}%`,
  };

  return (
    <main className="game-main">
      <div className="main-spacing"></div>
      <div className="main-current-level">Level 14</div>
      <div className="main-current-stage">3/10</div>
      <div className="main-planet-name">Planet name</div>
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
