import { configureStore } from "@reduxjs/toolkit";

import gameReducer from "@/lib/game/gameSlice";
import upgradeReducer from "@/lib/game/upgradeSlice";

export const makeStore = () => {
  return configureStore({
    reducer: {
      game: gameReducer,
      upgrade: upgradeReducer,
    },
  });
};

export type AppStore = ReturnType<typeof makeStore>;

export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];
