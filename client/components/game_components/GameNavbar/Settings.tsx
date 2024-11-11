"use client";

import { useState } from "react";

import { Box, Modal, ThemeProvider, Typography } from "@mui/material";
import { Settings as SettingsIcon } from "@mui/icons-material";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { SetSettings, SettingsState } from "@/lib/game/settingsSlice";
import { defaultTheme } from "@/assets/defaultTheme";

export default function Settings() {
  const settingsData = useAppSelector((state) => state.settings);
  const dispatch = useAppDispatch();

  const [open, setOpen] = useState<boolean>(false);
  const [settings, setSettings] = useState<SettingsState>(settingsData);

  const descriptions = [
    "Toggle Planet idle animation",
    "Toggle Planet destroy animation",
    "Toggle previous & next level animation",
    "Toggle all animations",
    // Maybe later think about DamageText to be toggled
  ];

  const handleModalOpen = () => {
    setOpen(true);
  };

  const handleModalClose = () => {
    setOpen(false);
  };

  const handleOptionToggle = (option: number) => {
    // Toggle all animations
    if (option != 4) {
      dispatch(
        SetSettings({ [`option${option}`]: !settingsData[`option${option}`] })
      );
      setSettings({
        ...settings,
        [`option${option}`]: !settingsData[`option${option}`],
      });
    } else {
      dispatch(
        SetSettings({
          option1: !settingsData.option4,
          option2: !settingsData.option4,
          option3: !settingsData.option4,
          option4: !settingsData.option4,
        })
      );
      setSettings({
        option1: !settingsData.option4,
        option2: !settingsData.option4,
        option3: !settingsData.option4,
        option4: !settingsData.option4,
      });
    }
  };

  return (
    <>
      <ThemeProvider theme={defaultTheme}>
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
                      checked={settingsData[`option${i + 1}`]}
                      onChange={() => handleOptionToggle(i + 1)}
                    ></input>
                  </Box>
                );
              })}
            </Box>
          </Box>
        </Modal>
      </ThemeProvider>
    </>
  );
}
