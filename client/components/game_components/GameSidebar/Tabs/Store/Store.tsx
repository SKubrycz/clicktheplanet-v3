"use client";

import { useContext } from "react";

import StoreElement from "./StoreElement";

import "./Store.scss";
import { GameContext } from "@/app/game/page";

interface IStoreElement {
  title: string;
  desc: string;
}

export default function Store() {
  let data = useContext(GameContext);

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
            data={data}
          ></StoreElement>
        );
      })}
    </div>
  );
}
