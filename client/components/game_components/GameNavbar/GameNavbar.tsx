import "./GameNavbar.scss";

export default function GameNavbar() {
  return (
    <nav className="game-navbar">
      <div className="game-navbar-content">
        <div className="game-navbar-title">Click the planet</div>
      </div>
      <div className="game-navbar-content">
        <div>Settings</div>
        <div>Profile</div>
      </div>
    </nav>
  );
}
