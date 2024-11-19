import { createSlice, PayloadAction } from "@reduxjs/toolkit";

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

export const diamondUpgradeObject: DiamondUpgradeWrapper = {
  1: {
    index: 1,
    level: 0,
    multiplier: "1.0",
    cost: 1,
  },
};

export const diamondUpgradeSlice = createSlice({
  name: "diamondupgrade",
  initialState: diamondUpgradeObject,
  reducers: {
    UpdateDiamondUpgrade: (
      state: DiamondUpgradeWrapper,
      action: PayloadAction<DiamondUpgradeWrapper>
    ) => {
      return {
        ...action.payload,
      };
    },
  },
});

export const { UpdateDiamondUpgrade } = diamondUpgradeSlice.actions;

export default diamondUpgradeSlice.reducer;
