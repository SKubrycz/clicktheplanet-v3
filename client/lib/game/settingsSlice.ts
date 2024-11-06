import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface SettingsState {
  [key: string]: boolean;
}

const initialState: SettingsState = {
  option1: true,
  option2: true,
  option3: true,
  option4: true,
};

export const settingsSlice = createSlice({
  name: "settings",
  initialState,
  reducers: {
    SetSettings: (state, action: PayloadAction<SettingsState>) => {
      return {
        ...state,
        ...action.payload,
      };
    },
  },
});

export const { SetSettings } = settingsSlice.actions;
export default settingsSlice.reducer;
