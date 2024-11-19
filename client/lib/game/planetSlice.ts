import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface Planet {
  planetName: string;
  currentHealth: string;
  healthPercent: number;
  maxHealth: string;
}

const planetObject: Planet = {
  planetName: "Planet_name",
  currentHealth: "10",
  healthPercent: 100,
  maxHealth: "10",
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
