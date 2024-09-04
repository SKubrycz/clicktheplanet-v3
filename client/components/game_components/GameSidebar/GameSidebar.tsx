"use client";

import { useState } from "react";

import Tabs from "./Tabs/Tabs";

import "./GameSidebar.scss";

export default function GameSidebar() {
  const tabs: Array<string> = ["Store", "Ship", "Stats"];

  const [tabTitle, setTabTitle] = useState<string>(tabs[0]);

  const handleTabTitle = (el: string) => {
    setTabTitle(el);
  };

  return (
    <aside className="game-sidebar">
      <div className="game-sidebar-currency">
        <div style={{ display: "flex" }}>
          <div className="test-yellow">Gold"icon"</div>: 842 678K
        </div>
        <div style={{ display: "flex" }}>
          <div className="test-blue">Diamond"icon"</div>: 450
        </div>
      </div>
      <div className="game-sidebar-options">
        <div className="game-sidebar-tabs">
          {tabs.map((el) => {
            return (
              <div
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
