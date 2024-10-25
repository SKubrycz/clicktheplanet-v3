"use client";

import { useState, useRef } from "react";

import { Box, Modal, Typography } from "@mui/material";
import { Settings as SettingsIcon } from "@mui/icons-material";

export default function Settings() {
  const [open, setOpen] = useState<boolean>(false);

  const option1Ref = useRef<HTMLInputElement>(null);
  const option2Ref = useRef<HTMLInputElement>(null);
  const option3Ref = useRef<HTMLInputElement>(null);

  const handleModalOpen = () => {
    setOpen(true);
  };

  const handleModalClose = () => {
    setOpen(false);
  };

  return (
    <>
      <div
        aria-labelledby="modal-modal-title"
        onClick={() => handleModalOpen()}
      >
        <SettingsIcon></SettingsIcon>
      </div>
      <Modal open={open} onClose={() => handleModalClose()}>
        <Box
          sx={{
            minWidth: 300,
            p: "1em",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
            position: "absolute",
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            border: "1px solid rgb(100, 100, 100)",
            borderRadius: "5px",
            outline: 0,
            boxShadow: "0px 0px 15px black",
            backgroundColor: "rgba(10, 10, 10, 0.9)",
          }}
        >
          <Typography
            id="modal-modal-title"
            variant="h5"
            sx={{ paddingBottom: "0.5em" }}
          >
            Settings
          </Typography>
          <Box
            sx={{
              width: "100%",
              display: "flex",
              flexDirection: "column",
              ".MuiBox-root": {
                padding: "0.5em",
              },
            }}
          >
            <Box sx={{ display: "flex", justifyContent: "space-between" }}>
              <Typography variant="body1">Setting 1 example</Typography>
              <input ref={option1Ref} type="checkbox"></input>
            </Box>
            <Box sx={{ display: "flex", justifyContent: "space-between" }}>
              <Typography variant="body1">Setting 2 example</Typography>
              <input ref={option2Ref} type="checkbox"></input>
            </Box>
            <Box sx={{ display: "flex", justifyContent: "space-between" }}>
              <Typography variant="body1">
                Lorem ipsum dolor sit amet
              </Typography>
              <input ref={option3Ref} type="checkbox"></input>
            </Box>
          </Box>
        </Box>
      </Modal>
    </>
  );
}
