"use client";

import { useState, useEffect } from "react";

import "./GameMain.scss";

export default function GameMain() {
  const [width, setWidth] = useState<number>(100);

  useEffect(() => {
    const interval = setInterval(() => {
      setWidth((width) => width - 10);
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    console.log(width);
    if (width <= 0) setWidth(100);
  }, [width]);

  return (
    <main className="game-main">
      <div className="main-spacing"></div>
      <div className="main-current-level">Level 14</div>
      <div className="main-current-stage">3/10</div>
      <div className="main-planet-name">Planet name</div>
      <div className="main-planet-image"></div>
      {/* Temporary */}
      <div className="main-planet-health-title">Health</div>
      <div className="main-planet-healthbar">
        <div className="main-planet-healthbar-amount">466/1000</div>
        <div className="main-planet-healthbar-wrapper">
          <div
            className="main-planet-healthbar-above"
            style={{ width: `${width}%` }}
          ></div>
          <div className="main-planet-healthbar-under"></div>
        </div>
      </div>
    </main>
  );
}
