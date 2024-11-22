import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface DiamondPlanet {
  diamonds: number;
  chance: number;
  isDiamondPlanet: boolean;
}

export interface Planet {
  planetName: string;
  currentHealth: string;
  healthPercent: number;
  maxHealth: string;
  diamondPlanet: DiamondPlanet;
}

const planetObject: Planet = {
  planetName: "Planet_name",
  currentHealth: "10",
  healthPercent: 100,
  maxHealth: "10",
  diamondPlanet: {
    diamonds: 1,
    chance: 0.001,
    isDiamondPlanet: false,
  },
};

export const planetSlice = createSlice({
  name: "planet",
  initialState: planetObject,
  reducers: {
    UpdatePlanet: (state, action: PayloadAction<Planet>) => {
      return { ...action.payload };
    },
  },
});

export const { UpdatePlanet } = planetSlice.actions;

export default planetSlice.reducer;
