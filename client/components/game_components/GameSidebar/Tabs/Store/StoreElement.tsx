"use client";

import { useMemo, useState } from "react";

import type { Data } from "@/lib/game/gameSlice";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { Add } from "@mui/icons-material";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";

interface StoreElementProps {
  index: number;
  title: string;
  description: string;
  image?: string;
  data: Data | undefined;
}

export default function StoreElement({
  index,
  title,
  description,
  image,
  data,
}: StoreElementProps) {
  const gameData = useAppSelector((state) => state.game);
  const upgradeData = useAppSelector((state) => state.upgrade);
  const dispatch = useAppDispatch();

  const [levels, setLevels] = useState<number | undefined>(undefined);

  const levelsAmount = useMemo(() => {
    setLevels(upgradeData?.levels);
  }, [upgradeData?.levels]); 

  return (
    <div className="store-element">
      <div className="store-element-left-wrapper">
        {image ? (
          <img src={image} className="store-element-image"></img>
        ) : (
          <div className="store-element-image-div">Image</div>
        )}
        <div className="store-element-info-wrapper">
          <div className="store-element-title">
            {title} - Level {gameData?.store[index]?.level}
          </div>
          <div className="store-element-description">
            {description} - Damage: {gameData?.store[index]?.damage}
          </div>
          <div>Cost: {gameData?.store[index]?.cost}</div>
        </div>
      </div>
      <div className="store-element-action-wrapper">
      {(levels && levels > 1) ? <div>x{levels}</div> : undefined}
      <Add
        onClick={() =>
          dispatch(
            UpgradeElement({
              upgrade: "store",
              index: index,
              levels: upgradeData.levels,
            })
          )
        }
        sx={{ width: 50, height: 50, cursor: "pointer" }}
      ></Add>
      </div>
    </div>
  );
}
