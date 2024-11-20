import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Ship {
  index: number;
  level: number;
  cost: string;
  multiplier: number;
  damage: string;
  locked: boolean;
}

interface ShipWrapper {
  [key: number]: Ship;
}

interface ShipMessage {
  gold: string;
  diamonds: number;
  ship: Ship;
}

export const shipObject: ShipWrapper = {
  1: {
    index: 1,
    level: 0,
    cost: "1",
    multiplier: 0.0,
    damage: "1",
    locked: true,
  },
};

export const shipSlice = createSlice({
  name: "ship",
  initialState: shipObject,
  reducers: {
    UpdateShip: (state: ShipWrapper, action: PayloadAction<ShipWrapper>) => {
      return {
        ...state,
        ...action.payload,
      };
    },
  },
});

export const { UpdateShip } = shipSlice.actions;

export default shipSlice.reducer;
