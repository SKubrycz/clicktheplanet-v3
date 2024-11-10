"use client";

import { Modal } from "@mui/material";

import { useAppSelector } from "@/lib/hooks";

interface ShipDiamondUpgradeProps {
  open: boolean;
  handleModalClose: () => void;
}

export default function ShipDiamondUpgrade({
  open,
  handleModalClose,
}: ShipDiamondUpgradeProps) {
  const gameData = useAppSelector((state) => state.game);

  return (
    <Modal open={open} onClose={() => handleModalClose()}>
      <div className="ship-diamond-upgrade-wrapper">
        <div>
          {Object.entries(gameData.ship).map(([key, ship]) => {
            return <div key={key}>{ship.multiplier}</div>;
          })}
        </div>
      </div>
    </Modal>
  );
}
