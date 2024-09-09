import GameNavbar from "@/components/game_components/GameNavbar/GameNavbar";
import GameSidebar from "@/components/game_components/GameSidebar/GameSidebar";
import GameMain from "@/components/game_components/GameMain/GameMain";

export default function Game() {
  const socket = new WebSocket("ws://localhost:8000/ws_game");

  socket.addEventListener("open", (e) => {
    console.log("WebSocket connection established");
    socket.send("Message to server");
  });

  socket.addEventListener("message", (e) => {
    console.log("From the server: ", e.data);
  });

  socket.addEventListener("close", (e) => {
    console.log("Disconnected");
  });

  return (
    <div className="game-wrapper">
      <GameNavbar></GameNavbar>
      <div className="game-content-wrapper">
        <GameSidebar></GameSidebar>
        <GameMain></GameMain>
      </div>
    </div>
  );
}
