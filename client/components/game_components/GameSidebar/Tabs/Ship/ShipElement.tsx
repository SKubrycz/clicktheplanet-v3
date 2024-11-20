"use client";

import React, { useMemo, useRef, useState, useEffect } from "react";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetErrorMessage, SetError } from "@/lib/game/errorSlice";

import { IconButton } from "@mui/material";
import { ArrowUpward, ElectricBoltOutlined } from "@mui/icons-material";

import Gold from "@/assets/svg/gold.svg";
import Image from "next/image";

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
  const shipData = useAppSelector((state) => state.ship);
  const diamondUpgradeData = useAppSelector((state) => state.diamondUpgrade);
  const upgradeData = useAppSelector((state) => state.upgrade);
  const errorData = useAppSelector((state) => state.error);
  const dispatch = useAppDispatch();

  const [levels, setLevels] = useState<number | undefined>(undefined);
  const [locked, setLocked] = useState<boolean>(true);

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
    if (shipData[index]?.locked) setLocked(true);
    else if (!shipData[index]?.locked) setLocked(false);
  }, [shipData[index]?.locked]);

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
      {locked ? <div className="element-locked"></div> : undefined}
      <div className="ship-element-info-wrapper">
        <div className="ship-element-title">
          {title} - Level: {shipData[index]?.level} | Multiplier: x
          {shipData[index]?.multiplier}
          {gameData?.diamondUpgradesUnlocked ? (
            <>
              {" "}
              |{" "}
              <ElectricBoltOutlined
                sx={{
                  width: 22,
                  height: 22,
                  color: "lightskyblue",
                }}
              ></ElectricBoltOutlined>
              x{diamondUpgradeData[index]?.multiplier}
            </>
          ) : undefined}
        </div>
        <div className="ship-element-description">
          {description}{" "}
          {index !== 4 && index !== 3 ? shipData[index]?.damage : ""}
          {index == 4 ? gameData?.planetGold : ""}
        </div>
        <div>
          Cost: {shipData[index]?.cost}{" "}
          <Image src={Gold} alt="gold" width={15} height={15}></Image>
        </div>
      </div>

      <div className="ship-element-action-wrapper">
        {levels && levels > 1 ? <div>x{levels}</div> : undefined}
        <IconButton sx={{ margin: 0, padding: 0, color: "white" }}>
          <ArrowUpward
            onClick={(e) => handleUpgradeClick(e)}
            sx={{ width: 40, height: 40, padding: "0.4em", cursor: "pointer" }}
          ></ArrowUpward>
        </IconButton>
      </div>
    </div>
  );
}
