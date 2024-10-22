"use client";

import ProfileModal from "./ProfileModal";

import "./GameNavbar.scss";
import Settings from "./Settings";

interface GameNavbarProps {
  handleSocketClose: () => Promise<void>;
}

export default function GameNavbar({ handleSocketClose }: GameNavbarProps) {
  return (
    <nav className="game-navbar">
      <div className="game-navbar-content">
        <div className="game-navbar-title">Click the planet</div>
      </div>
      <div className="game-navbar-content">
        <Settings></Settings>
        <ProfileModal handleSocketClose={handleSocketClose}></ProfileModal>
      </div>
    </nav>
  );
}
