"use client";

import { useState, useEffect } from "react";

import StoreElement from "./StoreElement";

import "./Store.scss";
import { useAppSelector } from "@/lib/hooks";

interface IStoreElement {
  title: string;
  desc: string;
}

export default function Store() {
  const gameData = useAppSelector((state) => state.game);
  const [storeElementsArr, setStoreElementsArr] = useState<IStoreElement[]>([
    { title: `Element 1`, desc: `Element 1 description` },
    { title: `Element 2`, desc: `Element 2 description` },
    { title: `Element 3`, desc: `Element 3 description` },
    { title: `Element 4`, desc: `Element 4 description` },
    { title: `Element 5`, desc: `Element 5 description` },
    { title: `Element 6`, desc: `Element 6 description` },
  ]);

  return (
    <div className="store-content-scroll">
      {storeElementsArr.map((el, i) => {
        return (
          <StoreElement
            key={i}
            index={i + 1}
            title={el.title}
            description={el.desc}
            data={gameData}
          ></StoreElement>
        );
      })}
    </div>
  );
}
