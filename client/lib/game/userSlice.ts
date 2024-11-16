import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface UserState {
  login: string;
}

const initialState: UserState = {
  login: "",
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    SetUser: (state, action: PayloadAction<UserState>) => {
      return {
        ...action.payload,
      };
    },
  },
});

export const { SetUser } = userSlice.actions;
export default userSlice.reducer;
