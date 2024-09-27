"use client";

import { useState, useEffect, useRef } from "react";
import { useRouter } from "next/navigation";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";
import {
  Init,
  Click,
  Upgrade,
  UpdateStore,
  DealDps,
  gameObject,
  UpdateShip,
} from "@/lib/game/gameSlice";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { CircularProgress } from "@mui/material";

interface ActionMessage {
  action: string;
  data: any;
}

export default function Game() {
  const [loading, setLoading] = useState<boolean>(true);
  const gameData = useAppSelector((state) => state.game);
  const upgradeData = useAppSelector((state) => state.upgrade);
  const dispatch = useAppDispatch();
  //const [data, setData] = useState<Data | undefined>(gameObject);
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
    if (upgradeData && upgradeData.index != -1) {
      if (socket.current) socket.current.send(JSON.stringify(upgradeData));
    }
  }, [upgradeData]);

  useEffect(() => {
    handleGetGame();

    socket.current = new WebSocket("ws://localhost:8000/ws_game");

    socket.current.onopen = (e: Event) => {
      setLoading(false);
      console.log("WebSocket connection established");
      if (socket.current) {
        socket.current.send("init");
        socket.current.send("dps");
      }
    };

    socket.current.onmessage = (e: MessageEvent) => {
      let message: ActionMessage = JSON.parse(e.data);
      //console.log("From the server: ", message);
      if (message.action === "init") {
        dispatch(Init(message.data));
        console.log(message.data);
      }
      if (message.action === "click") {
        dispatch(Click(message.data));
        console.log(message.data);
      }
      if (message.action === "upgrade") {
        dispatch(Upgrade(message.data));
      }
      // if (message.action === "store") {
      //   dispatch(UpdateStore(message.data));
      //   console.log(gameData);
      // }
      // if (message.action === "ship") {
      //   dispatch(UpdateShip(message.data));
      // }
      if (message.action === "dps") {
        dispatch(DealDps(message.data));
      }
    };

    socket.current.onclose = () => {
      console.log("Disconnected");
    };
  }, []);

  return (
    <div className="game-wrapper">
      {loading ? (
        <CircularProgress
          sx={{
            position: "absolute",
            transform: "translate(-50%, -50%)",
            top: "50%",
            left: "50%",
          }}
        ></CircularProgress>
      ) : (
        <>
          <GameNavbar></GameNavbar>
          <div className="game-content-wrapper">
            <GameSidebar></GameSidebar>
            <GameMain planetClick={handlePlanetClickData}></GameMain>
          </div>
        </>
      )}
    </div>
  );
}
