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
      <StatsElement
        title="Current damage:"
        data={data?.currentDamage}
      ></StatsElement>
      <StatsElement
        title="Highest level ever:"
        data={data?.maxLevel}
      ></StatsElement>
      <StatsElement
        title="Planets destroyed:"
        data={data?.planetsDestroyed}
      ></StatsElement>
    </div>
  );
}
