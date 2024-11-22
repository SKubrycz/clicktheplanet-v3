import { createSlice, PayloadAction, current } from "@reduxjs/toolkit";
import { DiamondPlanet } from "./planetSlice";

export interface Store {
  index: number;
  level: number;
  cost: string;
  damage: string;
  locked: boolean;
}

export interface StoreWrapper {
  [key: number]: Store;
}

interface StoreMessage {
  gold: string;
  diamonds: number;
  store: Store;
}

interface DiamondUpgrade {
  index: number;
  level: number;
  multiplier: string;
  cost: number;
}

interface DiamondUpgradeWrapper {
  [key: number]: DiamondUpgrade;
}

interface DiamondUpgradeMessage {
  gold: string;
  diamonds: number;
  diamondUpgrade: DiamondUpgrade;
}

export interface Ship {
  index: number;
  level: number;
  cost: string;
  multiplier: number;
  damage: string;
  locked: boolean;
}

export interface ShipWrapper {
  [key: number]: Ship;
}

interface ShipMessage {
  gold: string;
  diamonds: number;
  ship: Ship;
}

interface UpgradeMessage {
  gold: string;
  diamonds: number;
  currentDamage: string;
  maxDamage: string;
}

export interface DamageDone {
  damage: string;
  critical: boolean;
}

interface GameDps {
  gold: string;
  diamonds: number;
  currentDamage: string;
  maxDamage: string;
  diamondUpgradesUnlocked: boolean;
  planetGold: string;
  isBoss: boolean;
  currentLevel: number;
  maxLevel: number;
  currentStage: number;
  maxStage: number;
  planetsDestroyed: string;
}

export interface Data {
  gold: string;
  diamonds: number;
  currentDamage: string;
  maxDamage: string;
  damageDone: DamageDone;
  planetName: string;
  currentHealth: string;
  healthPercent: number;
  maxHealth: string;
  diamondUpgradesUnlocked: boolean;
  planetGold: string;
  isBoss: boolean;
  diamondPlanet: DiamondPlanet;
  currentLevel: number;
  maxLevel: number;
  currentStage: number;
  maxStage: number;
  planetsDestroyed: string;
  store: StoreWrapper;
  ship: ShipWrapper;
  diamondUpgrade: DiamondUpgradeWrapper;
  error: string;
}

export interface Game {
  gold: string;
  diamonds: number;
  currentDamage: string;
  maxDamage: string;
  damageDone: DamageDone;
  diamondUpgradesUnlocked: boolean;
  planetGold: string;
  isBoss: boolean;
  currentLevel: number;
  maxLevel: number;
  currentStage: number;
  maxStage: number;
  planetsDestroyed: string;
}

export const gameObject: Game = {
  gold: "100",
  diamonds: 0,
  currentDamage: "1",
  maxDamage: "1",
  damageDone: {
    damage: "0",
    critical: false,
  },
  planetGold: "1",
  isBoss: false,
  currentLevel: 1,
  maxLevel: 1,
  currentStage: 1,
  maxStage: 1,
  diamondUpgradesUnlocked: false,
  planetsDestroyed: "0",
};

export const gameSlice = createSlice({
  name: "game",
  initialState: gameObject,
  reducers: {
    GameInit: (state, action: PayloadAction<Game>) => {
      console.log(state);
      return { ...action.payload };
    },
    Click: (state, action: PayloadAction<Game>) => {
      //console.log(current(state));
      return { ...action.payload };
    },
    Upgrade: (state, action: PayloadAction<UpgradeMessage>) => {
      return {
        ...state,
        gold: action.payload.gold,
        diamonds: action.payload.diamonds,
        currentDamage: action.payload.currentDamage,
        maxDamage: action.payload.maxDamage,
      };
    },
    DealDps: (state, action: PayloadAction<GameDps>) => {
      return {
        ...state,
        gold: action.payload.gold,
        diamonds: action.payload.diamonds,
        currentDamage: action.payload.currentDamage,
        maxDamage: action.payload.maxDamage,
        planetGold: action.payload.planetGold,
        isBoss: action.payload.isBoss,
        currentLevel: action.payload.currentLevel,
        maxLevel: action.payload.maxLevel,
        diamondUpgradesUnlocked: action.payload.diamondUpgradesUnlocked,
        currentStage: action.payload.currentStage,
        maxStage: action.payload.maxStage,
        planetsDestroyed: action.payload.planetsDestroyed,
      };
    },
  },
});

export const { GameInit, Click, Upgrade, DealDps } = gameSlice.actions;

export default gameSlice.reducer;
