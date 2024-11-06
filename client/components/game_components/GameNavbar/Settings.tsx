"use client";

import { useState } from "react";

import { Box, Modal, Typography } from "@mui/material";
import { Settings as SettingsIcon } from "@mui/icons-material";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { SetSettings } from "@/lib/game/settingsSlice";

export default function Settings() {
  const settingsData = useAppSelector((state) => state.settings);
  const dispatch = useAppDispatch();

  const [open, setOpen] = useState<boolean>(false);

  const descriptions = [
    "Toggle Planet idle animation",
    "Toggle Planet destroy animation",
    "Toggle previous & next level animation",
    "Toggle all animations",
  ];

  const handleModalOpen = () => {
    setOpen(true);
  };

  const handleModalClose = () => {
    setOpen(false);
  };

  const handleOptionToggle = (option: number) => {
    dispatch(
      SetSettings({ [`option${option}`]: !settingsData[`option${option}`] })
    );
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
            {descriptions.map((el, i) => {
              return (
                <Box
                  key={i}
                  sx={{
                    margin: "0.3em 0",
                    display: "flex",
                    justifyContent: "space-between",
                  }}
                >
                  <Typography variant="body1">{String(el)}</Typography>
                  <input
                    type="checkbox"
                    className="settings-checkbox"
                    defaultChecked={settingsData[`option${i + 1}`]}
                    onClick={() => handleOptionToggle(i + 1)}
                  ></input>
                </Box>
              );
            })}
          </Box>
        </Box>
      </Modal>
    </>
  );
}
