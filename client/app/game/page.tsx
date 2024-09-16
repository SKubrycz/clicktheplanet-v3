"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";

export default function Game() {
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

  const handlePlanetClickData = (data: string) => {
    socket.send(data);
    console.log("click");
  };

  useEffect(() => {
    handleGetGame();
  }, []);

  return (
    <div className="game-wrapper">
      <GameNavbar></GameNavbar>
      <div className="game-content-wrapper">
        <GameSidebar></GameSidebar>
        <GameMain planetClick={handlePlanetClickData}></GameMain>
      </div>
    </div>
  );
}
