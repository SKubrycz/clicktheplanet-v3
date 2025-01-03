"use client";

import { useState } from "react";

import StoreElement from "./StoreElement";

import "./Store.scss";
import { useAppSelector } from "@/lib/hooks";

interface IStoreElement {
  title: string;
  desc: string;
}

export default function Store() {
  const storeData = useAppSelector((state) => state.store);

  const [storeElementsArr] = useState<IStoreElement[]>([
    { title: `Element 1`, desc: `Element 1 description` },
    { title: `Element 2`, desc: `Element 2 description` },
    { title: `Element 3`, desc: `Element 3 description` },
    { title: `Element 4`, desc: `Element 4 description` },
    { title: `Element 5`, desc: `Element 5 description` },
    { title: `Element 6`, desc: `Element 6 description` },
    { title: `Element 7`, desc: `Element 7 description` },
    { title: `Element 8`, desc: `Element 8 description` },
    { title: `Element 9`, desc: `Element 9 description` },
    { title: `Element 10`, desc: `Element 10 description` },
    { title: `Element 11`, desc: `Element 11 description` },
    { title: `Element 12`, desc: `Element 12 description` },
    { title: `Element 13`, desc: `Element 13 description` },
    { title: `Element 14`, desc: `Element 14 description` },
    { title: `Element 15`, desc: `Element 15 description` },
    { title: `Element 16`, desc: `Element 16 description` },
    { title: `Element 17`, desc: `Element 17 description` },
    { title: `Element 18`, desc: `Element 18 description` },
    { title: `Element 19`, desc: `Element 19 description` },
    { title: `Element 20`, desc: `Element 20 description` },
    { title: `Element 21`, desc: `Element 21 description` },
    { title: `Element 22`, desc: `Element 22 description` },
    { title: `Element 23`, desc: `Element 23 description` },
    { title: `Element 24`, desc: `Element 24 description` },
    { title: `Element 25`, desc: `Element 25 description` },
    { title: `Element 26`, desc: `Element 26 description` },
    { title: `Element 27`, desc: `Element 27 description` },
    { title: `Element 28`, desc: `Element 28 description` },
    { title: `Element 29`, desc: `Element 29 description` },
    { title: `Element 30`, desc: `Element 30 description` },
  ]);

  return (
    <div className="store-content-scroll">
      {Object.keys(storeData).map((el, i) => {
        return (
          <StoreElement
            key={i}
            index={i + 1}
            title={storeElementsArr[i]?.title}
            description={storeElementsArr[i]?.desc}
            locked={storeData[Number(el)]?.locked}
          ></StoreElement>
        );
      })}
    </div>
  );
}
