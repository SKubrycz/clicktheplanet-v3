"use client";

import React, { useRef } from "react";
import { Provider } from "react-redux";
import { AppStore, makeStore } from "@/lib/store";

interface StoreProviderProps {
  children: React.ReactNode;
}

export default function StoreProvider({ children }: StoreProviderProps) {
  const storeRef = useRef<AppStore>();

  if (!storeRef.current) {
    storeRef.current = makeStore();
    console.log("---> Store initialized");
  }

  return <Provider store={storeRef.current}>{children}</Provider>;
}
