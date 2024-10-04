"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";

import { Popover, PopoverProps } from "@mui/material";

import "./GameNavbar.scss";

interface GameNavbarProps {
  handleOnLogout: () => Promise<void>;
}

export default function GameNavbar({ handleOnLogout }: GameNavbarProps) {
  const [open, setOpen] = useState<boolean>(false);
  const [anchorEl, setAnchorEl] = useState<PopoverProps["anchorEl"]>(null);
  const router = useRouter();

  const handlePopoverOpen = (e: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(e.currentTarget);
    setOpen(true);
  };

  const handlePopoverClose = () => {
    setOpen(false);
  };

  const id = open ? "virtual-element-popover" : undefined;

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
      handleOnLogout().then(() => {
        router.push("/");
      });
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <nav className="game-navbar">
      <div className="game-navbar-content">
        <div className="game-navbar-title">Click the planet</div>
      </div>
      <div className="game-navbar-content">
        <div>Settings</div>
        <div
          aria-describedby={id}
          aria-haspopup="true"
          onClick={(e) => handlePopoverOpen(e)}
        >
          Profile
        </div>
        <Popover
          open={open}
          anchorOrigin={{
            vertical: "bottom",
            horizontal: "center",
          }}
          transformOrigin={{
            vertical: "top",
            horizontal: "center",
          }}
          onClose={handlePopoverClose}
          anchorEl={anchorEl}
        >
          <button onClick={(e) => handleLogout(e)}>Logout</button>
        </Popover>
      </div>
    </nav>
  );
}
