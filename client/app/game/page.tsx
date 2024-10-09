"use client";

import { useState, useEffect, useRef } from "react";
import { useRouter } from "next/navigation";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";
import { Init, Click, Upgrade, DealDps } from "@/lib/game/gameSlice";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetErrorMessage, SetError } from "@/lib/game/errorSlice";
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
  const levelData = useAppSelector((state) => state.level);
  const dispatch = useAppDispatch();
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
        router.push("/"); //if token is not valid
      });

    console.log(res);
  };

  const handlePlanetClickData = (data: string) => {
    if (socket.current) socket.current.send(data);
    console.log("click");
  };

  const handleBulkUpgrade = (e: KeyboardEvent, pressed: boolean) => {
    if (pressed) {
      if (e.key === "z") {
        dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 10 }));
      }
      if (e.key === "x") {
        dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 100 }));
      }
      if (e.key === "c") {
        dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 1000 }));
      }
    } else if (!pressed) {
      dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 1 }));
    }
  };

  useEffect(() => {
    if (upgradeData && upgradeData.index != -1) {
      if (socket.current) socket.current.send(JSON.stringify(upgradeData));
    }
  }, [upgradeData]);

  useEffect(() => {
    console.log(levelData);
    if (levelData.action === "previous" || levelData.action === "next") {
      if (socket.current) socket.current.send(levelData.action);
    }
  }, [levelData]);

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

      if (message.action === "error") {
        dispatch(SetErrorMessage(message.data.error));
        console.log(message.data);
      }
      if (message.action === "dps") {
        dispatch(DealDps(message.data));
      }
    };

    socket.current.onclose = () => {
      console.log("Disconnected");
    };

    document.addEventListener("keydown", (e) => handleBulkUpgrade(e, true));
    document.addEventListener("keyup", (e) => handleBulkUpgrade(e, false));

    return () => {
      document.removeEventListener("keydown", (e) =>
        handleBulkUpgrade(e, true)
      );
      document.removeEventListener("keyup", (e) => handleBulkUpgrade(e, false));
    };
  }, []);

  const handleSocketClose = (): Promise<void> => {
    if (socket.current) socket.current.close(1000);
    return Promise.resolve();
  };

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
          <GameNavbar
            handleSocketClose={() => handleSocketClose()}
          ></GameNavbar>
          <div className="game-content-wrapper">
            <GameSidebar></GameSidebar>
            <GameMain planetClick={handlePlanetClickData}></GameMain>
          </div>
        </>
      )}
    </div>
  );
}
