import { configureStore } from "@reduxjs/toolkit";

import gameReducer from "@/lib/game/gameSlice";
import planetReducer from "@/lib/game/planetSlice";
import storeReducer from "@/lib/game/storeSlice";
import shipReducer from "@/lib/game/shipSlice";
import diamondUpgradeReducer from "@/lib/game/diamondUpgradeSlice";
import upgradeReducer from "@/lib/game/upgradeSlice";
import levelReducer from "@/lib/game/levelSlice";
import errorReducer from "@/lib/game/errorSlice";
import globalErrorReducer from "@/lib/game/globalErrorSlice";
import settingsReducer from "@/lib/game/settingsSlice";
import userReducer from "@/lib/game/userSlice";

export const makeStore = () => {
  return configureStore({
    reducer: {
      game: gameReducer,
      planet: planetReducer,
      store: storeReducer,
      ship: shipReducer,
      diamondUpgrade: diamondUpgradeReducer,
      upgrade: upgradeReducer,
      level: levelReducer,
      error: errorReducer,
      globalError: globalErrorReducer,
      settings: settingsReducer,
      user: userReducer,
    },
  });
};

export type AppStore = ReturnType<typeof makeStore>;

export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];
