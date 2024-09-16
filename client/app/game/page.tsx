"use client";

import { useState, useEffect, useRef, useCallback } from "react";
import { useRouter } from "next/navigation";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";

export interface Data {
  name: string;
  currentHealth: string;
  healthPercent: number;
  maxHealth: string;
  currentLevel: number;
  currentStage: number;
}

export default function Game() {
  const [data, setData] = useState<Data | undefined>({
    name: "Planet_name",
    currentHealth: "10",
    healthPercent: 100,
    maxHealth: "10",
    currentLevel: 1,
    currentStage: 1,
  });
  const socket = useRef<WebSocket | null>(null);
  const router = useRouter();

  const handleGetGame = async () => {
    let res = await fetch("http://localhost:8000/game", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      mode: "cors",
      credentials: "include",
    })
      .then((res) => {
        if (!res.ok) {
          throw new Error(res.statusText);
        }
        return res.json();
      })
      .catch((err) => {
        if (err instanceof Error) {
          console.error("Error: ", err.message);
        } else {
          console.error("Error: ", String(err));
        }
        //router.push("/"); //if token is not valid
      });

    console.log(res);
  };

  const handlePlanetClickData = (data: string) => {
    if (socket.current) socket.current.send(data);
    console.log("click");
  };

  useEffect(() => {
    handleGetGame();

    socket.current = new WebSocket("ws://localhost:8000/ws_game");

    socket.current.onopen = (e: Event) => {
      console.log("WebSocket connection established");
      if (socket.current) {
        socket.current.send("click");
      }
    };

    socket.current.onmessage = (e: MessageEvent) => {
      let data = JSON.parse(e.data);
      console.log("From the server: ", data);
      setData(JSON.parse(e.data));
    };

    socket.current.onclose = () => {
      console.log("Disconnected");
    };
  }, []);

  return (
    <div className="game-wrapper">
      <GameNavbar></GameNavbar>
      <div className="game-content-wrapper">
        <GameSidebar></GameSidebar>
        <GameMain data={data} planetClick={handlePlanetClickData}></GameMain>
      </div>
    </div>
  );
}
