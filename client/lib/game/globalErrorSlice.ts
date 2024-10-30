import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface GlobalErrorState {
  message: string;
  statusCode: number | string;
}

const initialState: GlobalErrorState = {
  message: "",
  statusCode: 0,
};

export const globalErrorSlice = createSlice({
  name: "globalerror",
  initialState,
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
