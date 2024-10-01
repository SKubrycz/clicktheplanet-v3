import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Action {
  action: string;
}

export const levelSlice = createSlice({
  name: "level",
  initialState: {
    action: "",
  },
  reducers: {
    setLevel: (state, action: PayloadAction<Action>) => {
      return { ...state, action: action.payload.action };
    },
  },
});

export const { setLevel } = levelSlice.actions;

export default levelSlice.reducer;
