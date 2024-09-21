import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Store {
  index: number;
  level: number;
  cost: string;
  damage: string;
}

export interface Ship {
  index: number;
  level: number;
  cost: string;
  multiplier: number;
  damage: string;
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
      state = action.payload;
      console.log("Init reducer");
      console.log(state);
    },
    Click: (state, action: PayloadAction<Data>) => {
      state = action.payload;
      console.log("Click reducer");
      console.log(state);
    },
    UpgradeStore: (state, action: PayloadAction<Data>) => {},
    UpgradeShip: (state, action: PayloadAction<Data>) => {},
  },
});

export const { Init, Click, UpgradeStore, UpgradeShip } = gameSlice.actions;

export default gameSlice.reducer;
