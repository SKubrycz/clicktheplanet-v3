"use client";

import React, { useMemo, useRef, useState, useEffect } from "react";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetErrorMessage, SetError } from "@/lib/game/errorSlice";

import { ArrowUpward } from "@mui/icons-material";

interface ShipElementProps {
  index: number;
  title: string;
  description: string;
}

export default function ShipElement({
  index,
  title,
  description,
}: ShipElementProps) {
  const gameData = useAppSelector((state) => state.game);
  const upgradeData = useAppSelector((state) => state.upgrade);
  const errorData = useAppSelector((state) => state.error);
  const dispatch = useAppDispatch();

  const [levels, setLevels] = useState<number | undefined>(undefined);

  const timeout = useRef<NodeJS.Timeout | null>(null);

  const levelsAmount = useMemo(() => {
    setLevels(upgradeData?.levels);
  }, [upgradeData?.levels]);

  const handleUpgradeClick = (e: React.MouseEvent) => {
    dispatch(
      UpgradeElement({
        upgrade: "ship",
        index: index,
        levels: upgradeData?.levels,
      })
    );
    dispatch(
      SetError({
        isVisible: true,
        pos: {
          x: e.clientX,
          y: e.clientY,
        },
      })
    );
  };

  useEffect(() => {
    if (timeout.current) {
      clearTimeout(timeout.current);
    }
    timeout.current = null;
    timeout.current = setTimeout(() => {
      dispatch(
        SetError({
          isVisible: false,
          pos: {
            x: 0,
            y: 0,
          },
        })
      );
    }, 1500);

    dispatch(SetErrorMessage(""));

    return () => {
      if (timeout.current) clearTimeout(timeout.current);
    };
  }, [errorData?.isVisible]);

  return (
    <div className="ship-element">
      <div className="ship-element-info-wrapper">
        <div className="ship-element-title">
          {title} - Level: {gameData?.ship[index]?.level} | Multiplier: x
          {gameData?.ship[index]?.multiplier}
        </div>
        <div className="ship-element-description">
          {description}{" "}
          {index !== 4 && index !== 3 ? gameData?.ship[index]?.damage : ""}
          {index == 4 ? gameData?.planetGold : ""}
        </div>
        <div>Cost: {gameData?.ship[index]?.cost}</div>
      </div>

      <div className="ship-element-action-wrapper">
        {levels && levels > 1 ? <div>x{levels}</div> : undefined}
        <ArrowUpward
          onClick={(e) => handleUpgradeClick(e)}
          sx={{ width: 40, height: 40, cursor: "pointer" }}
        ></ArrowUpward>
      </div>
    </div>
  );
}
