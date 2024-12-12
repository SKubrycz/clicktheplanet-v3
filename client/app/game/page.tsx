"use client";

import { useState, useEffect, useRef } from "react";
import { useRouter } from "next/navigation";

import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";
import { GameInit, Click, Upgrade, DealDps, Data } from "@/lib/game/gameSlice";
import { UpdatePlanet } from "@/lib/game/planetSlice";
import { UpdateStore } from "@/lib/game/storeSlice";
import { UpdateShip } from "@/lib/game/shipSlice";
import { UpdateDiamondUpgrade } from "@/lib/game/diamondUpgradeSlice";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetErrorMessage } from "@/lib/game/errorSlice";
import { SetGlobalError } from "@/lib/game/globalErrorSlice";
import { SetSettings } from "@/lib/game/settingsSlice";
import type { SettingsState } from "@/lib/game/settingsSlice";
import { SetUser } from "@/lib/game/userSlice";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import {
  Alert,
  CircularProgress,
  Snackbar,
  ThemeProvider,
} from "@mui/material";
import { defaultTheme } from "@/assets/defaultTheme";

interface ActionMessage {
  action: string;
  data: Data;
}

function settingsExist(settings: SettingsState) {
  for (let option in settings) {
    if (settings[option] == null) {
      return false;
    }
  }
  return true;
}

export default function Game() {
  const [loading, setLoading] = useState<boolean>(true);
  const [open, setOpen] = useState<boolean>(false);
  const gameData = useAppSelector((state) => state.game);
  const upgradeData = useAppSelector((state) => state.upgrade);
  const levelData = useAppSelector((state) => state.level);
  const globalErrorData = useAppSelector((state) => state.globalError);
  const settingsData = useAppSelector((state) => state.settings);
  const dispatch = useAppDispatch();
  const socket = useRef<WebSocket | null>(null);
  const pressedKeys = useRef<boolean[]>([false, false, false]);
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
      dispatch(SetUser(data));
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
        router.push("/");
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
      if (e.key === "z" && !pressedKeys.current[0]) {
        dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 10 }));
        pressedKeys.current[0] = true;
      }
      if (e.key === "x" && !pressedKeys.current[1]) {
        dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 100 }));
        pressedKeys.current[1] = true;
      }
      if (e.key === "c" && !pressedKeys.current[2]) {
        dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 1000 }));
        pressedKeys.current[2] = true;
      }
    } else if (!pressed) {
      dispatch(UpgradeElement({ upgrade: "", index: -1, levels: 1 }));
      pressedKeys.current[0] = false;
      pressedKeys.current[1] = false;
      pressedKeys.current[2] = false;
    }
  };

  useEffect(() => {
    if (upgradeData && upgradeData.index != -1) {
      if (socket.current) socket.current.send(JSON.stringify(upgradeData));
    }
  }, [upgradeData]);

  useEffect(() => {
    if (
      levelData.action === "previous" ||
      levelData.action === "next" ||
      levelData.action === "maxlevel"
    ) {
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

        const {
          gold,
          diamonds,
          currentDamage,
          maxDamage,
          damageDone,
          diamondUpgradesUnlocked,
          planetGold,
          isBoss,
          diamondPlanet,
          currentLevel,
          maxLevel,
          currentStage,
          maxStage,
          planetsDestroyed,
          planetName,
          currentHealth,
          healthPercent,
          maxHealth,
          store,
          ship,
          diamondUpgrade,
        } = message.data;

        // Initialize:
        // gameSlice
        // planetSlice
        // storeSlice
        // shipSlice
        // diamondUpgradeSlice
        // (?) settingsSlice

        dispatch(
          GameInit({
            gold: gold,
            diamonds: diamonds,
            currentDamage: currentDamage,
            maxDamage: maxDamage,
            damageDone: damageDone,
            diamondUpgradesUnlocked: diamondUpgradesUnlocked,
            planetGold: planetGold,
            isBoss: isBoss,
            currentLevel: currentLevel,
            maxLevel: maxLevel,
            currentStage: currentStage,
            maxStage: maxStage,
            planetsDestroyed: planetsDestroyed,
          })
        );

        dispatch(
          UpdatePlanet({
            planetName: planetName,
            currentHealth: currentHealth,
            healthPercent: healthPercent,
            maxHealth: maxHealth,
            diamondPlanet: diamondPlanet,
          })
        );

        dispatch(UpdateStore(store));
        dispatch(UpdateShip(ship));
        dispatch(UpdateDiamondUpgrade(diamondUpgrade));
      }
      if (message.action === "click") {
        console.log(message.data);

        const {
          gold,
          diamonds,
          currentDamage,
          maxDamage,
          damageDone,
          diamondUpgradesUnlocked,
          planetGold,
          isBoss,
          diamondPlanet,
          currentLevel,
          maxLevel,
          currentStage,
          maxStage,
          planetsDestroyed,
          planetName,
          currentHealth,
          healthPercent,
          maxHealth,
        } = message.data;

        // Update:
        // gameSlice
        // planetSlice

        dispatch(
          Click({
            gold: gold,
            diamonds: diamonds,
            currentDamage: currentDamage,
            maxDamage: maxDamage,
            damageDone: damageDone,
            diamondUpgradesUnlocked: diamondUpgradesUnlocked,
            planetGold: planetGold,
            isBoss: isBoss,
            currentLevel: currentLevel,
            maxLevel: maxLevel,
            currentStage: currentStage,
            maxStage: maxStage,
            planetsDestroyed: planetsDestroyed,
          })
        );

        dispatch(
          UpdatePlanet({
            planetName: planetName,
            currentHealth: currentHealth,
            healthPercent: healthPercent,
            maxHealth: maxHealth,
            diamondPlanet: diamondPlanet,
          })
        );
      }
      if (
        message.action === "upgrade" ||
        message.action === "diamond-upgrade"
      ) {
        const {
          gold,
          diamonds,
          currentDamage,
          maxDamage,
          store,
          ship,
          diamondUpgrade,
        } = message.data;

        // Update
        // gameSlice
        // storeSlice
        // shipSlice
        // diamondUpgradeSlice

        dispatch(
          Upgrade({
            gold: gold,
            diamonds: diamonds,
            currentDamage: currentDamage,
            maxDamage: maxDamage,
          })
        );

        dispatch(UpdateStore(store));
        dispatch(UpdateShip(ship));
        dispatch(UpdateDiamondUpgrade(diamondUpgrade));
      }

      if (message.action === "error") {
        dispatch(SetErrorMessage(message.data.error));
        console.log(message.data);
      }
      if (message.action === "dps") {
        const {
          gold,
          diamonds,
          currentDamage,
          maxDamage,
          diamondUpgradesUnlocked,
          planetGold,
          isBoss,
          diamondPlanet,
          currentLevel,
          maxLevel,
          currentStage,
          maxStage,
          planetsDestroyed,
          planetName,
          currentHealth,
          healthPercent,
          maxHealth,
          store,
          ship,
        } = message.data;

        // Update
        // gameSlice
        // planetSlice
        // storeSlice
        // shipSlice

        dispatch(
          DealDps({
            gold: gold,
            diamonds: diamonds,
            currentDamage: currentDamage,
            maxDamage: maxDamage,
            planetGold: planetGold,
            isBoss: isBoss,
            currentLevel: currentLevel,
            maxLevel: maxLevel,
            diamondUpgradesUnlocked: diamondUpgradesUnlocked,
            currentStage: currentStage,
            maxStage: maxStage,
            planetsDestroyed: planetsDestroyed,
          })
        );

        dispatch(
          UpdatePlanet({
            planetName: planetName,
            currentHealth: currentHealth,
            healthPercent: healthPercent,
            maxHealth: maxHealth,
            diamondPlanet: diamondPlanet,
          })
        );

        dispatch(UpdateStore(store));
        dispatch(UpdateShip(ship));
      }
    };

    socket.current.onclose = () => {
      console.log("Disconnected");
      if (socket.current && socket.current.readyState === 3) {
        dispatch(
          SetGlobalError({
            message: "You've been disconnected from Websockets",
            statusCode: `(${socket.current.readyState}) CLOSED`,
          })
        );
        setOpen(true);
      }
    };

    let settings = localStorage.getItem("settings");
    if (settings) {
      const parsedSettings: SettingsState = JSON.parse(settings);
      if (settingsExist(parsedSettings)) dispatch(SetSettings(parsedSettings));
    }

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
    <ThemeProvider theme={defaultTheme}>
      <div
        className={
          gameData?.isBoss
            ? `game-wrapper game-boss-background`
            : `game-wrapper`
        }
      >
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
    </ThemeProvider>
  );
}
