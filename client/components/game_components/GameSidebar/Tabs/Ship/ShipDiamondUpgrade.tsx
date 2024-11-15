"use client";

import { Box, Dialog, Typography } from "@mui/material";

import DiamondUpgrade from "./DiamondUpgrade";

interface ShipDiamondUpgradeProps {
  open: boolean;
  handleModalClose: () => void;
}

export default function ShipDiamondUpgrade({
  open,
  handleModalClose,
}: ShipDiamondUpgradeProps) {
  return (
    <Dialog
      open={open}
      onClose={() => handleModalClose()}
      sx={{
        ".MuiPaper-root": {
          background: "none",
        },
      }}
    >
      <div className="ship-diamond-upgrade-wrapper">
        <Typography variant="h6">Diamond Upgrades</Typography>
        <Box sx={{ width: "100%" }}>
          <DiamondUpgrade title={"Dps"} index={1}></DiamondUpgrade>
          <DiamondUpgrade title={"Click damage"} index={2}></DiamondUpgrade>
          <DiamondUpgrade title={"Critical damage"} index={3}></DiamondUpgrade>
          <DiamondUpgrade title={"Planet gold"} index={4}></DiamondUpgrade>
        </Box>
      </div>
    </Dialog>
  );
}
