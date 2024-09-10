"use client";

import { useEffect } from "react";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";

export default function Game() {
  const handleGetGame = async () => {
    try {
      let res = await fetch("http://localhost:8000/game", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        mode: "cors",
        credentials: "include",
      });
      let data = await res.json();
      console.log(data);
    } catch (err) {}
  };

  const socket = new WebSocket("ws://localhost:8000/ws_game");

  socket.addEventListener("open", (e) => {
    console.log("WebSocket connection established");
    socket.send("Message to server");
  });

  socket.addEventListener("message", (e) => {
    console.log("From the server: ", e.data);
  });

  socket.addEventListener("close", (e) => {
    console.log("Disconnected");
  });

  useEffect(() => {
    handleGetGame();
  }, []);

  return (
    <div className="game-wrapper">
      <GameNavbar></GameNavbar>
      <div className="game-content-wrapper">
        <GameSidebar></GameSidebar>
        <GameMain></GameMain>
      </div>
    </div>
  );
}
