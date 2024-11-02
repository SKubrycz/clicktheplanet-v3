"use client";

import { useMemo, useRef, useState, useEffect } from "react";
import Image from "next/image";

import type { Data } from "@/lib/game/gameSlice";
import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetError, SetErrorMessage } from "@/lib/game/errorSlice";
import { IconButton } from "@mui/material";
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
        upgrade: "store",
        index: index,
        levels: upgradeData.levels,
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
    <div className="store-element">
      {gameData?.store[index]?.locked ? (
        <div className="element-locked"></div>
      ) : (
        ""
      )}
      <div className="store-element-left-wrapper">
        {image ? (
          <Image src={image} alt={`store-${index}-image`} className="store-element-image"></Image>
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
        {levels && levels > 1 ? <div>x{levels}</div> : undefined}
        <IconButton
          aria-label={`store-upgrade-button-${index}`}
          sx={{ margin: 0, color: "white" }}
        >
          <Add
            onClick={(e) => handleUpgradeClick(e)}
            sx={{ width: 50, height: 50, cursor: "pointer" }}
          ></Add>
        </IconButton>
      </div>
    </div>
  );
}
