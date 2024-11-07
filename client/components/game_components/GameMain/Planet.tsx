"use client";

import { useState, useRef, useEffect } from "react";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";

import { KeyboardArrowLeft, KeyboardArrowRight } from "@mui/icons-material";
import { setLevel } from "@/lib/game/levelSlice";

interface Loaded {
  isLoaded: boolean;
  count: number;
}

interface PlanetProps {
  planetRef: React.RefObject<HTMLCanvasElement>;
  click: (e: React.MouseEvent<HTMLCanvasElement>) => void;
}

export default function Planet({ planetRef, click }: PlanetProps) {
  const gameData = useAppSelector((state) => state.game);
  const settingsData = useAppSelector((state) => state.settings);
  const dispatch = useAppDispatch();

  const [loaded, setLoaded] = useState<Loaded>({ isLoaded: false, count: 0 });

  const workerRef = useRef<Worker>();
  // Timeout for level change arrows
  const timeout = useRef<NodeJS.Timeout | null>(null);
  // Blocks the level choice when the timeout hasn't ran yet
  const block = useRef<boolean>(false);

  let seed =
    gameData?.currentLevel * gameData?.currentLevel +
    gameData?.currentStage * gameData?.currentStage;
  let size = 1;
  let ctx: CanvasRenderingContext2D | null;
  let imageData: ImageData;
  let weights = useRef<number[]>([]);
  const rgb = 200;
  let breatheAnimationKeyframes = useRef([
    {
      filter: `drop-shadow(0px 0px 2px rgb(${rgb * weights.current[0]}, ${
        rgb * weights.current[1]
      }, ${rgb * weights.current[1]}))`,
      offset: 0.0,
    },
    {
      filter: `drop-shadow(0px 0px 15px rgb(${rgb * weights.current[0]}, ${
        rgb * weights.current[1]
      }, ${rgb * weights.current[1]}))`,
      offset: 0.5,
    },
    {
      filter: `drop-shadow(0px 0px 2px rgb(${rgb * weights.current[0]}, ${
        rgb * weights.current[1]
      }, ${rgb * weights.current[1]}))`,
      offset: 1.0,
    },
  ]);

  if (planetRef.current) {
    ctx = planetRef.current.getContext("2d");
    if (ctx)
      imageData = ctx.createImageData(
        planetRef.current.width,
        planetRef.current.height
      );
  }

  if (planetRef.current)
    size = planetRef.current.width * planetRef.current.height;

  function draw() {
    if (planetRef.current && ctx) {
      ctx.putImageData(imageData, 0, 0);
    }
  }

  useEffect(() => {
    if (planetRef?.current) {
      seed =
        gameData?.currentLevel * gameData?.currentLevel +
        gameData?.currentStage * gameData?.currentStage;

      workerRef.current = new Worker(
        new URL("../../../utilities/workers/drawWorker.ts", import.meta.url)
      );
      workerRef.current.postMessage([
        planetRef.current.width,
        planetRef.current.height,
        size,
        seed,
        gameData?.isBoss,
        imageData,
      ]);

      workerRef.current.onmessage = (e) => {
        imageData = e.data[0];
        weights.current = e.data[1];

        draw();

        breatheAnimationKeyframes.current = [
          {
            filter: `drop-shadow(0px 0px 2px rgb(${rgb * weights.current[0]}, ${
              rgb * weights.current[1]
            }, ${rgb * weights.current[1]}))`,
            offset: 0.0,
          },
          {
            filter: `drop-shadow(0px 0px 15px rgb(${
              rgb * weights.current[0]
            }, ${rgb * weights.current[1]}, ${rgb * weights.current[1]}))`,
            offset: 0.5,
          },
          {
            filter: `drop-shadow(0px 0px 2px rgb(${rgb * weights.current[0]}, ${
              rgb * weights.current[1]
            }, ${rgb * weights.current[1]}))`,
            offset: 1.0,
          },
        ];
      };
    }

    return () => {
      if (workerRef.current) workerRef.current.terminate();
    };
  }, [gameData.planetName]);

  useEffect(() => {
    if (weights.current.length > 1 && typeof weights.current[0] == "number") {
      if (settingsData.option1) {
        if (planetRef.current)
          planetRef.current.animate(breatheAnimationKeyframes.current, {
            duration: 4000,
            iterations: Infinity,
          });
      } else {
        if (planetRef.current) {
          const animations = planetRef.current.getAnimations();
          if (animations.length > 0) {
            console.log(animations);
            animations.forEach((anim, i) => {
              animations[i].cancel();
            });
          }
        }
      }
    }
  }, [breatheAnimationKeyframes.current, settingsData.option1]);

  const animationStyle = {
    previous: "200ms swipeRight linear 1, levitate 5.5s infinite",
    next: "200ms swipeLeft linear 1, levitate 5.5s infinite",
    destroyed: "250ms destroyed linear 1, levitate 5.5s infinite",
  };

  useEffect(() => {
    if (settingsData.option2) {
      if (planetRef.current && loaded.count > 1) {
        planetRef.current.style.animation = "none";
        planetRef.current.offsetHeight;
        planetRef.current.style.animation = animationStyle.destroyed;
      }

      // To keep the animation from executing on initial component load
      if (loaded.count < 2) {
        setLoaded({ ...loaded, count: loaded.count + 1 });
      } else if (loaded.count == 2) {
        setLoaded({ ...loaded, isLoaded: true });
      }
    }
  }, [gameData.planetsDestroyed]);

  const animatePrevious = () => {
    dispatch(setLevel({ action: "previous" }));
    if (settingsData.option3) {
      if (planetRef.current && gameData.currentLevel > 1 && !block.current) {
        block.current = true;
        planetRef.current.style.animation = "none";
        planetRef.current.offsetHeight;
        planetRef.current.style.animation = animationStyle.previous;
        timeout.current = setTimeout(() => {
          block.current = false;
        }, 200);
      }
    }
  };

  const animateNext = () => {
    dispatch(setLevel({ action: "next" }));
    if (settingsData.option3) {
      if (
        planetRef.current &&
        gameData.currentLevel !== gameData.maxLevel &&
        !block.current
      ) {
        block.current = true;
        planetRef.current.style.animation = "none";
        planetRef.current.offsetHeight;
        planetRef.current.style.animation = animationStyle.next;
        timeout.current = setTimeout(() => {
          block.current = false;
        }, 200);
      }
    }
  };

  return (
    <>
      <KeyboardArrowLeft
        aria-label="previous-level-button"
        onClick={() => {
          animatePrevious();
        }}
        sx={{
          marginRight: "3em",
          cursor: "pointer",
          fontSize: "30px",
          "&:hover": {
            filter: "drop-shadow(0px 0px 12px white)",
          },
        }}
      ></KeyboardArrowLeft>
      <canvas
        className="main-planet-image"
        width="200"
        height="200"
        ref={planetRef}
        onClick={click}
      ></canvas>
      <KeyboardArrowRight
        aria-label="next-level-button"
        onClick={() => {
          animateNext();
        }}
        sx={{
          marginLeft: "3em",
          cursor: "pointer",
          fontSize: "30px",
          "&:hover": {
            filter: "drop-shadow(0px 0px 12px white)",
          },
        }}
      ></KeyboardArrowRight>
    </>
  );
}
