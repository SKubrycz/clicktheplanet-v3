import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Store {
  index: number;
  level: number;
  cost: string;
  damage: string;
  locked: boolean;
}

interface StoreWrapper {
  [key: number]: Store;
}

interface StoreMessage {
  gold: string;
  diamonds: number;
  store: Store;
}

export const storeObject: StoreWrapper = {
  1: {
    index: 1,
    level: 0,
    cost: "1",
    damage: "1",
    locked: true,
  },
};

export const storeSlice = createSlice({
  name: "store",
  initialState: storeObject,
  reducers: {
    UpdateStore: (state: StoreWrapper, action: PayloadAction<StoreWrapper>) => {
      return {
        ...action.payload,
      };
    },
  },
});

export const { UpdateStore } = storeSlice.actions;

export default storeSlice.reducer;
