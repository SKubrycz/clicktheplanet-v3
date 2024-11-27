"use client";

import { useAppSelector } from "@/lib/hooks";

import StatsElement from "./StatsElement";

import "./Stats.scss";

export default function Stats() {
  const gameData = useAppSelector((state) => state.game);
  const planetData = useAppSelector((state) => state.planet);

  return (
    <div className="stats-content">
      <StatsElement title="Gold:" data={gameData?.gold}></StatsElement>
      <StatsElement title="Diamonds:" data={gameData?.diamonds}></StatsElement>
      <StatsElement
        title="Current damage:"
        data={gameData?.currentDamage}
      ></StatsElement>
      <StatsElement
        title="Highest damage ever:"
        data={gameData?.maxDamage}
      ></StatsElement>
      <StatsElement
        title="Current level:"
        data={gameData?.currentLevel}
      ></StatsElement>
      <StatsElement
        title="Highest level ever:"
        data={gameData?.maxLevel}
      ></StatsElement>
      <StatsElement
        title="Current stage:"
        data={gameData?.currentStage}
      ></StatsElement>
      <StatsElement
        title="Highest level max stage:"
        data={gameData?.maxStage}
      ></StatsElement>
      <StatsElement
        title="Planets destroyed:"
        data={gameData?.planetsDestroyed}
      ></StatsElement>
      {gameData?.diamondUpgradesUnlocked && (
        <StatsElement
          title="Diamond Planet chance:"
          data={`${planetData?.diamondPlanet?.chance * 100}%`}
        ></StatsElement>
      )}
    </div>
  );
}
