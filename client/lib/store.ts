import { configureStore } from "@reduxjs/toolkit";

import gameReducer from "@/lib/game/gameSlice";
import upgradeReducer from "@/lib/game/upgradeSlice";
import levelReducer from "@/lib/game/levelSlice";
import errorReducer from "@/lib/game/errorSlice";

export const makeStore = () => {
  return configureStore({
    reducer: {
      game: gameReducer,
      upgrade: upgradeReducer,
      level: levelReducer,
      error: errorReducer,
    },
  });
};

export type AppStore = ReturnType<typeof makeStore>;

export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];
