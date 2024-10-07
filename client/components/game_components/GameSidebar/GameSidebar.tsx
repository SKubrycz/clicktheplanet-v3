"use client";

import { useState, useRef } from "react";
import { useAppSelector } from "@/lib/hooks";
import ErrorMessage from "@/components/game_components/ErrorMessage/ErrorMessage";

import Tabs from "./Tabs/Tabs";

import "./GameSidebar.scss";

export default function GameSidebar() {
  const gameData = useAppSelector((state) => state.game);
  const errorData = useAppSelector((state) => state.error);
  const tabs: Array<string> = ["Store", "Ship", "Stats"];

  const [tabTitle, setTabTitle] = useState<string>(tabs[0]);

  const handleTabTitle = (el: string) => {
    setTabTitle(el);
  };

  return (
    <aside className="game-sidebar">
      {errorData?.isVisible && (
        <ErrorMessage
          message={errorData?.message}
          pos={errorData?.pos}
        ></ErrorMessage>
      )}
      <div className="game-sidebar-currency">
        <div style={{ display: "flex" }}>
          <div className="test-yellow">Gold"icon"</div>: {gameData?.gold}
        </div>
        <div style={{ display: "flex" }}>
          <div className="test-blue">Diamond"icon"</div>: {gameData?.diamonds}
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
