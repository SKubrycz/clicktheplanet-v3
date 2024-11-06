"use client";

import { useRef } from "react";
import { Provider } from "react-redux";
import { AppStore, makeStore } from "@/lib/store";
import { saveSettings } from "@/lib/localStorage";

interface StoreProviderProps {
  children: React.ReactNode;
}

export default function StoreProvider({ children }: StoreProviderProps) {
  const storeRef = useRef<AppStore>();

  if (!storeRef.current) {
    storeRef.current = makeStore();
    storeRef.current.subscribe(() => {
      const state = storeRef.current?.getState();
      if (state?.settings) saveSettings(state?.settings);
    });
    console.log("---> Store initialized");
  }

  return <Provider store={storeRef.current}>{children}</Provider>;
}
