import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export const upgradeSlice = createSlice({
  name: "upgrade",
  initialState: {
    upgrade: "",
    index: -1,
    levels: 1,
  },
  reducers: {
    UpgradeElement: (state, action: PayloadAction<any>) => {
      if (action.payload.levels === undefined || state.levels === undefined) {
        return {
          ...state,
          upgrade: action.payload.upgrade,
          index: action.payload.index,
          levels: 1,
        };
      } else {
        return {
          ...state,
          upgrade: action.payload.upgrade,
          index: action.payload.index,
          levels: action.payload.levels,
        };
      }
    },
  },
});

export const { UpgradeElement } = upgradeSlice.actions;

export default upgradeSlice.reducer;
