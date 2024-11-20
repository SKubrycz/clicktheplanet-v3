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
            locked={storeData[i + 1]?.locked}
          ></StoreElement>
        );
      })}
    </div>
  );
}
