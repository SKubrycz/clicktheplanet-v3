"use client";

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
  return (
    <div className="ship-element">
      <div className="ship-element-info-wrapper">
        <div className="ship-element-title">{title}</div>
        <div className="ship-element-description">{description}</div>
      </div>
      <div className="ship-element-action-wrapper">
        <ArrowUpward sx={{ width: 40, height: 40 }}></ArrowUpward>
      </div>
    </div>
  );
}
