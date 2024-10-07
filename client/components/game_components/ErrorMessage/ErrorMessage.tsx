import "./ErrorMessage.scss";

export interface Coords {
  x: number;
  y: number;
}

interface ErrorMessageProps {
  message: string | undefined;
  pos: Coords;
}

export default function ErrorMessage({ message, pos }: ErrorMessageProps) {
  return (
    <>
      <div
        className="game-error-message"
        style={{
          position: "absolute",
          top: pos.y,
          left: pos.x,
          animation: "1500ms moveAndTransparent 1",
        }}
      >
        {message}
      </div>
    </>
  );
}
