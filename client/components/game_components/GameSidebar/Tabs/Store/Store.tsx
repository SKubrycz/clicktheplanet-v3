"use client";

import StoreElement from "./StoreElement";

import "./Store.scss";
import { useAppSelector } from "@/lib/hooks";

interface IStoreElement {
  title: string;
  desc: string;
}

export default function Store() {
  const gameData = useAppSelector((state) => state.game);

  const storeElementsArr: IStoreElement[] = [
    { title: "Element 1 title", desc: "Brief element description" },
    { title: "Element 2 title", desc: "Brief element description" },
    { title: "Element 3 title", desc: "Brief element description" },
    { title: "Element 4 title", desc: "Brief element description" },
  ];

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
