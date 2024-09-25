import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Store {
  index: number;
  level: number;
  cost: string;
  damage: string;
}

interface StoreMessage {
  gold: string;
  diamonds: number;
  store: Store;
}

export interface Ship {
  index: number;
  level: number;
  cost: string;
  multiplier: number;
  damage: string;
}

interface ShipMessage {
  gold: string;
  diamonds: number;
  ship: Ship;
}

export interface Data {
  gold: string;
  diamonds: number;
  currentDamage: string;
  maxDamage: string;
  planetName: string;
  currentHealth: string;
  healthPercent: number;
  maxHealth: string;
  currentLevel: number;
  maxLevel: number;
  currentStage: number;
  maxStage: number;
  planetsDestroyed: string;
  store: Store[];
  ship: Ship[];
}

export interface UpgradeMessage {
  upgrade: "store" | "ship";
  index: number;
}

export const gameObject: Data = {
  gold: "100",
  diamonds: 0,
  currentDamage: "1",
  maxDamage: "1",
  planetName: "Planet_name",
  currentHealth: "10",
  healthPercent: 100,
  maxHealth: "10",
  currentLevel: 1,
  maxLevel: 1,
  currentStage: 1,
  maxStage: 1,
  planetsDestroyed: "0",
  store: [
    {
      index: 1,
      level: 0,
      cost: "0",
      damage: "0",
    },
  ],
  ship: [
    {
      index: 1,
      level: 0,
      cost: "0",
      multiplier: 0.0,
      damage: "0",
    },
  ],
};

export const gameSlice = createSlice({
  name: "game",
  initialState: gameObject,
  reducers: {
    Init: (state, action: PayloadAction<Data>) => {
      console.log("Init reducer");
      console.log(state);
      return { ...action.payload };
    },
    Click: (state, action: PayloadAction<Data>) => {
      console.log("Click reducer");
      console.log(state);
      return { ...action.payload };
    },
    UpdateStore: (state, action: PayloadAction<StoreMessage>) => {
      return {
        ...state,
        gold: action.payload.gold,
        diamonds: action.payload.diamonds,
        store: {
          ...state.store,
          [action.payload.store.index]: {
            index: action.payload.store.index,
            level: action.payload.store.level,
            cost: action.payload.store.cost,
            damage: action.payload.store.damage,
          },
        },
      };
    },
    DealDps: (state, action: PayloadAction<Data>) => {
      return {
        ...state,
        gold: action.payload.gold,
        diamonds: action.payload.diamonds,
        currentDamage: action.payload.currentDamage,
        maxDamage: action.payload.maxDamage,
        planetName: action.payload.planetName,
        currentHealth: action.payload.currentHealth,
        healthPercent: action.payload.healthPercent,
        maxHealth: action.payload.maxHealth,
        currentLevel: action.payload.currentLevel,
        maxLevel: action.payload.maxLevel,
        currentStage: action.payload.currentStage,
        maxStage: action.payload.maxStage,
        planetsDestroyed: action.payload.planetsDestroyed,
      };
    },
  },
});

export const { Init, Click, UpdateStore, DealDps } = gameSlice.actions;

export default gameSlice.reducer;
