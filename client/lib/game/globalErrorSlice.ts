import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface GlobalErrorState {
  message: string;
  statusCode: number;
}

export const globalErrorSlice = createSlice({
  name: "globalerror",
  initialState: {
    message: "",
    statusCode: 0,
  },
  reducers: {
    SetGlobalError: (state, action: PayloadAction<GlobalErrorState>) => {
      return {
        ...state,
        message: action.payload.message,
        statusCode: action.payload.statusCode,
      };
    },
  },
});

export const { SetGlobalError } = globalErrorSlice.actions;
export default globalErrorSlice.reducer;
