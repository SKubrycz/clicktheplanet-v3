import { createSlice, PayloadAction } from "@reduxjs/toolkit";

import type { Coords } from "@/components/game_components/ErrorMessage/ErrorMessage";

interface ErrorState {
  isVisible: boolean;
  pos: Coords;
}

export const errorSlice = createSlice({
  name: "error",
  initialState: {
    message: "",
    isVisible: false,
    pos: {
      x: 0,
      y: 0,
    },
  },
  reducers: {
    SetErrorMessage: (state, action: PayloadAction<string>) => {
      return { ...state, message: action.payload };
    },
    SetError: (state, action: PayloadAction<ErrorState>) => {
      return {
        ...state,
        isVisible: action.payload.isVisible,
        pos: action.payload.pos,
      };
    },
  },
});

export const { SetErrorMessage, SetError } = errorSlice.actions;
export default errorSlice.reducer;
