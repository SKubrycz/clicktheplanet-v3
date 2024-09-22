import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export const upgradeSlice = createSlice({
  name: "upgrade",
  initialState: {
    upgrade: "",
    index: -1,
  },
  reducers: {
    Upgrade: (state, action: PayloadAction<any>) => {
      console.log("Upgrade() upgradeSlice: ", action.payload);
      return { ...action.payload };
    },
  },
});

export const { Upgrade } = upgradeSlice.actions;

export default upgradeSlice.reducer;
