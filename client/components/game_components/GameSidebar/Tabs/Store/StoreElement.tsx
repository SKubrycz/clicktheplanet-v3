"use client";

import { useMemo, useRef, useState, useEffect } from "react";
import Image from "next/image";

import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetError, SetErrorMessage } from "@/lib/game/errorSlice";
import { IconButton } from "@mui/material";
import { Add } from "@mui/icons-material";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";

import Gold from "@/assets/svg/gold.svg";

interface StoreElementProps {
  index: number;
  title: string;
  description: string;
  image?: string;
}

export default function StoreElement({
  index,
  title,
  description,
  image,
}: StoreElementProps) {
  const storeData = useAppSelector((state) => state.store);
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
    if (storeData[index]?.locked) setLocked(true);
    else if (!storeData[index]?.locked) setLocked(false);
  }, [storeData[index]?.locked]);

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
      {storeData[index]?.locked ? <div className="element-locked"></div> : ""}
      <div className="store-element-left-wrapper">
        {image ? (
          <Image
            src={image}
            alt={`store-${index}-image`}
            className="store-element-image"
          ></Image>
        ) : (
          <div className="store-element-image-div">Image</div>
        )}
        <div className="store-element-info-wrapper">
          <div className="store-element-title">
            {title} - Level {storeData[index]?.level}
          </div>
          <div className="store-element-description">
            {description} - Damage: {storeData[index]?.damage}
          </div>
          <div>
            Cost: {storeData[index]?.cost}{" "}
            <Image src={Gold} alt="gold" width={15} height={15}></Image>
          </div>
        </div>
      </div>
      <div className="store-element-action-wrapper">
        {levels && levels > 1 ? <div>x{levels}</div> : undefined}
        <IconButton
          aria-label={`store-upgrade-button-${index}`}
          sx={{ margin: 0, padding: 0, color: "white" }}
        >
          <Add
            onClick={(e) => handleUpgradeClick(e)}
            sx={{ width: 50, height: 50, padding: "0.4em", cursor: "pointer" }}
          ></Add>
        </IconButton>
      </div>
    </div>
  );
}
