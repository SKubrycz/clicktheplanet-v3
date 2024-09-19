"use client";

import { useContext } from "react";
import { GameContext } from "@/app/game/page";

import StatsElement from "./StatsElement";

import "./Stats.scss";

import { Data } from "@/app/game/page";

export default function Stats() {
  let data: Data | undefined = useContext(GameContext);

  return (
    <div className="stats-content">
      <StatsElement title="Gold:" data={data?.gold}></StatsElement>
      <StatsElement title="Diamonds:" data={data?.diamonds}></StatsElement>
      <StatsElement
        title="Current damage:"
        data={data?.currentDamage}
      ></StatsElement>
      <StatsElement
        title="Highest damage ever:"
        data={data?.maxDamage}
      ></StatsElement>
      <StatsElement
        title="Current level:"
        data={data?.currentLevel}
      ></StatsElement>
      <StatsElement
        title="Highest level ever:"
        data={data?.maxLevel}
      ></StatsElement>
      <StatsElement
        title="Current stage:"
        data={data?.currentStage}
      ></StatsElement>
      <StatsElement
        title="Highest level max stage:"
        data={data?.maxStage}
      ></StatsElement>
      <StatsElement
        title="Planets destroyed:"
        data={data?.planetsDestroyed}
      ></StatsElement>
    </div>
  );
}
