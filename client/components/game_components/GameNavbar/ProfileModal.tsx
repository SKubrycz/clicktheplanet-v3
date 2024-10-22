"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";

import { Box, Modal, Typography } from "@mui/material";

interface ProfileModalProps {
  handleSocketClose: () => Promise<void>;
}

export default function ProfileModal({ handleSocketClose }: ProfileModalProps) {
  const [open, setOpen] = useState<boolean>(false);
  const [deleteOpen, setDeleteOpen] = useState<boolean>(false);
  const router = useRouter();

  const handleModalOpen = () => {
    setOpen(true);
  };

  const handleModalClose = () => {
    setOpen(false);
  };

  const handleLogout = async (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8000/logout", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        mode: "cors",
        credentials: "include",
      });
      const data = await res.json();
      console.log(data);
      handleSocketClose().then(() => {
        router.push("/");
      });
    } catch (err) {
      console.error(err);
    }
  };

  const handleDeleteAccount = () => {
    setDeleteOpen(true);
  };

  const handleDeleteAccountClose = () => {
    setDeleteOpen(false);
  };

  const handleDeleteAccountConfirm = async (e: React.MouseEvent) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8000/delete-account", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        mode: "cors",
        credentials: "include",
      });
      const data = await res.json();
      console.log(data);
      handleSocketClose().then(() => {
        router.push("/");
      });
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <>
      <div
        aria-labelledby="modal-modal-title"
        onClick={() => handleModalOpen()}
      >
        Profile
      </div>
      <Modal open={open} onClose={() => handleModalClose()}>
        <Box
          sx={{
            minWidth: 300,
            p: "0.5em",
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            alignItems: "center",
            position: "absolute",
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            border: "1px solid gray",
            borderRadius: "5px",
            boxShadow: "0px 0px 15px black",
            backgroundColor: "rgba(10, 10, 10, 0.9)",
          }}
        >
          <Typography id="modal-modal-title" variant="h5">
            Profile
          </Typography>
          <Box
            sx={{
              width: "90%",
              p: "0.7em",
              display: "flex",
              justifyContent: "space-around",
              alignItems: "center",
              border: "none",
            }}
          >
            <div className="profile-img-lg">Image</div>
            <Box>
              <Typography variant="h6">Nickname</Typography>
              <button
                className="profile-logout-btn"
                onClick={(e) => handleLogout(e)}
              >
                Logout
              </button>
            </Box>
            <Modal open={deleteOpen} onClose={() => handleDeleteAccountClose()}>
              <Box
                sx={{
                  minWidth: 300,
                  minHeight: 200,
                  position: "absolute",
                  top: "50%",
                  left: "50%",
                  transform: "translate(-50%, -50%)",
                  display: "flex",
                  flexDirection: "column",
                  justifyContent: "center",
                  alignItems: "center",
                  border: "1px solid white",
                  borderRadius: "5px",
                  backgroundColor: "rgb(0, 0, 0)",
                }}
              >
                <Typography variant="h5">Are you sure?</Typography>
                <Typography variant="body1">
                  This action is irreversible!
                </Typography>
                <Box>
                  <button
                    className="profile-delete-account-confirm"
                    onClick={(e) => handleDeleteAccountConfirm(e)}
                  >
                    Yes (Delete account)
                  </button>
                  <button
                    className="profile-delete-account-cancel"
                    onClick={() => handleDeleteAccountClose()}
                  >
                    No (Cancel)
                  </button>
                </Box>
              </Box>
            </Modal>
          </Box>
          <button
            className="profile-delete-account-modal-btn"
            onClick={() => handleDeleteAccount()}
          >
            Delete account
          </button>
        </Box>
      </Modal>
    </>
  );
}
