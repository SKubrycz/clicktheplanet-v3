import { SettingsState } from "./game/settingsSlice";

export const saveSettings = (state: SettingsState) => {
  const stateSettings = JSON.stringify(state);
  localStorage.setItem("settings", stateSettings);
};
