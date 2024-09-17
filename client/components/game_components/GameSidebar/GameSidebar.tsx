"use client";

import { useState, useContext } from "react";
import { GameContext } from "@/app/game/page";

import Tabs from "./Tabs/Tabs";

import "./GameSidebar.scss";

import { Data } from "@/app/game/page";

export default function GameSidebar() {
  let data: Data | undefined = useContext(GameContext);
  const tabs: Array<string> = ["Store", "Ship", "Stats"];

  const [tabTitle, setTabTitle] = useState<string>(tabs[0]);

  const handleTabTitle = (el: string) => {
    setTabTitle(el);
  };

  return (
    <aside className="game-sidebar">
      <div className="game-sidebar-currency">
        <div style={{ display: "flex" }}>
          <div className="test-yellow">Gold"icon"</div>: {data?.gold}
        </div>
        <div style={{ display: "flex" }}>
          <div className="test-blue">Diamond"icon"</div>: {data?.diamonds}
        </div>
      </div>
      <div className="game-sidebar-options">
        <div className="game-sidebar-tabs">
          {tabs.map((el, i) => {
            return (
              <div
                key={i}
                className="game-sidebar-tab-el"
                onClick={() => handleTabTitle(el)}
              >
                {el}
              </div>
            );
          })}
        </div>
        <div className="game-sidebar-options-content">
          <div className="game-sidebar-options-content-title">{tabTitle}</div>
          <Tabs tabTitle={tabTitle}></Tabs>
        </div>
      </div>
    </aside>
  );
}
