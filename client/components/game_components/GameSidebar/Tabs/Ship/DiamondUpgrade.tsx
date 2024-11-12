"use client";

import { useState, useRef, useEffect, useMemo } from "react";

import Image from "next/image";
import Diamond from "@/assets/svg/diamond.svg";

import { UpgradeElement } from "@/lib/game/upgradeSlice";
import { SetError, SetErrorMessage } from "@/lib/game/errorSlice";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";

import { Box, IconButton, Typography } from "@mui/material";
import { ElectricBoltOutlined } from "@mui/icons-material";

interface DiamondUpgradeProps {
  title: string;
  index: number;
}

export default function DiamondUpgrade({ title, index }: DiamondUpgradeProps) {
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
        upgrade: "diamond-upgrade",
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
    <div className="diamond-upgrade">
      <Box sx={{ padding: "0.4em" }}>
        <Typography variant="h6" sx={{ fontSize: "18px" }}>
          {title}
        </Typography>
        <div>
          Cost:{" "}
          <Image src={Diamond} alt="diamond" width={14} height={14}></Image>
        </div>
      </Box>
      <IconButton
        title="Upgrade"
        aria-label={`diamond-upgrade-button-${index}`}
        sx={{ margin: 0, padding: 0, color: "white" }}
      >
        <ElectricBoltOutlined
          onClick={(e) => handleUpgradeClick(e)}
          sx={{
            width: 30,
            height: 30,
            padding: "0.4em",
            color: "lightskyblue",
            cursor: "pointer",
            "&:hover": {
              color: "deepskyblue",
              filter: "drop-shadow(0px 0px 3px cornflowerblue)",
            },
          }}
        ></ElectricBoltOutlined>
      </IconButton>
    </div>
  );
}
