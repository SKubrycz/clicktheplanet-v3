"use client";

import { useState, useEffect, useRef } from "react";
import { useRouter } from "next/navigation";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";
import { Init, Click, Upgrade, DealDps } from "@/lib/game/gameSlice";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetErrorMessage } from "@/lib/game/errorSlice";
import { SetGlobalError } from "@/lib/game/globalErrorSlice";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { Alert, CircularProgress, Snackbar } from "@mui/material";

interface ActionMessage {
  action: string;
  data: any;
}

export default function Game() {
  const [loading, setLoading] = useState<boolean>(true);
  const [open, setOpen] = useState<boolean>(false);
  const gameData = useAppSelector((state) => state.game);
  const upgradeData = useAppSelector((state) => state.upgrade);
  const levelData = useAppSelector((state) => state.level);
  const globalErrorData = useAppSelector((state) => state.globalError);
  const dispatch = useAppDispatch();
  const socket = useRef<WebSocket | null>(null);
  const router = useRouter();

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
      if (!res.ok) {
        throw res;
      }
      let data = await res.json();
      console.log(data);
    } catch (err: unknown) {
      if (err instanceof Response) {
        console.error("Error: ", err);
        dispatch(
          SetGlobalError({
            message: String(err.statusText),
            statusCode: err.status,
          })
        );
        setOpen(true);
      }
    }
  };

  const handleClose = () => {
    setOpen(false);
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
        console.log(message.data);
        dispatch(Init(message.data));
      }
      if (message.action === "click") {
        console.log(message.data);
        dispatch(Click(message.data));
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
          <Snackbar open={open} autoHideDuration={5000} onClose={handleClose}>
            <Alert
              onClose={handleClose}
              severity="error"
              variant="filled"
              sx={{ width: "100%" }}
            >
              <b>{globalErrorData.message}</b> - Status code:{" "}
              {globalErrorData.statusCode}
            </Alert>
          </Snackbar>
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
