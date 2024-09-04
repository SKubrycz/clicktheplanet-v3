import GameNavbar from "@/components/game_components/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar";
import GameMain from "@/components/game_components/GameMain";

export default function Game() {
  return (
    <div>
      <GameNavbar></GameNavbar>
      <div className="game-content-wrapper">
        <GameSidebar></GameSidebar>
        <GameMain></GameMain>
      </div>
    </div>
  );
}
