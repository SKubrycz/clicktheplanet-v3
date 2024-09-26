"use client";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";
import { Upgrade } from "@/lib/game/upgradeSlice";

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
  const dispatch = useAppDispatch();

  return (
    <div className="ship-element">
      <div className="ship-element-info-wrapper">
        <div className="ship-element-title">
          {title} - Level: {gameData?.ship[index]?.level} | Multiplier:{" "}
          {gameData?.ship[index]?.multiplier}
        </div>
        <div className="ship-element-description">
          {description} - Damage: {gameData?.ship[index]?.damage}
        </div>
        <div>Cost: {gameData?.ship[index]?.cost}</div>
      </div>
      <div className="ship-element-action-wrapper">
        <ArrowUpward
          onClick={() => dispatch(Upgrade({ upgrade: "ship", index: index }))}
          sx={{ width: 40, height: 40 }}
        ></ArrowUpward>
      </div>
    </div>
  );
}
