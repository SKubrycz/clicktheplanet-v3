"use client";

import React, { useState, useRef, useEffect } from "react";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";

import { KeyboardArrowLeft, KeyboardArrowRight } from "@mui/icons-material";
import { setLevel } from "@/lib/game/levelSlice";
import { IconButton } from "@mui/material";

interface Loaded {
  isLoaded: boolean;
  count: number;
}

interface PlanetProps {
  planetRef: React.RefObject<HTMLCanvasElement>;
  click: (e: React.MouseEvent<HTMLCanvasElement>) => void;
}

const applyBreathe = (
  planetRef: HTMLCanvasElement,
  breatheAnimationKeyframes: Keyframe[]
): Animation => {
  const breatheAnim = planetRef.animate(breatheAnimationKeyframes, {
    duration: 4000,
    iterations: Infinity,
  });

  return breatheAnim;
};

const applyLevitate = (planetRef: HTMLCanvasElement): Animation => {
  const levitateAnim = planetRef.animate(
    [
      {
        transform: "translateY(3%)",
        offset: 0.5,
      },
    ],
    {
      id: "levitate",
      duration: 5500,
      iterations: Infinity,
    }
  );
  return levitateAnim;
};

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
  const breatheAnim = useRef<Animation | null>(null);
  const levitateAnim = useRef<Animation | null>(null);

  let seed =
    gameData?.currentLevel * gameData?.currentLevel +
    gameData?.currentStage * gameData?.currentStage;
  seed = (31764 * (seed + 936)) % 5492340;

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
      seed = (31764 * (seed + 936)) % 5492340;

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
        if (planetRef.current) {
          breatheAnim.current = planetRef.current.animate(
            breatheAnimationKeyframes.current,
            {
              duration: 4000,
              iterations: Infinity,
            }
          );
        }
      } else {
        if (planetRef.current) {
          const animations = planetRef.current.getAnimations();
          if (animations.length > 0) {
            console.log(animations);
            animations.forEach((anim, i) => {
              animations[i].cancel();
            });
            console.log(
              `animations after cancel(): ${animations[0]} - l: ${animations.length}`
            );
          }
        }
      }
    }
  }, [breatheAnimationKeyframes.current, settingsData.option1]);

  useEffect(() => {
    if (planetRef.current) {
      levitateAnim.current = applyLevitate(planetRef.current);
    }
  }, []);

  const animationStyle = {
    previous: "200ms swipeRight linear 1",
    next: "200ms swipeLeft linear 1",
    destroyed: "250ms destroyed linear 1",
  };

  useEffect(() => {
    if (settingsData.option2) {
      if (planetRef.current && loaded.count > 1) {
        levitateAnim.current?.cancel();
        planetRef.current.style.animation = "none";
        planetRef.current.offsetHeight;
        planetRef.current.style.animation = animationStyle.destroyed;
        levitateAnim.current = applyLevitate(planetRef.current);
      }

      // To keep the animation from executing on initial component load
      if (loaded.count < 2) {
        setLoaded({ ...loaded, count: loaded.count + 1 });
      } else if (loaded.count == 2) {
        setLoaded({ ...loaded, isLoaded: true });
      }
    }
  }, [gameData.planetsDestroyed]);

  const goToPrevious = () => {
    if (planetRef.current && gameData.currentLevel > 1 && !block.current) {
      dispatch(setLevel({ action: "previous" }));
      block.current = true;
      levitateAnim.current?.cancel();
      planetRef.current.style.animation = "none";
      planetRef.current.offsetHeight;
      if (settingsData.option3)
        planetRef.current.style.animation = animationStyle.previous;
      levitateAnim.current = applyLevitate(planetRef.current);
      timeout.current = setTimeout(() => {
        block.current = false;
      }, 200);
    }
  };

  const goToNext = () => {
    if (
      planetRef.current &&
      gameData.currentLevel !== gameData.maxLevel &&
      !block.current
    ) {
      dispatch(setLevel({ action: "next" }));
      block.current = true;
      levitateAnim.current?.cancel();
      planetRef.current.style.animation = "none";
      planetRef.current.offsetHeight;
      if (settingsData.option3)
        planetRef.current.style.animation = animationStyle.next;
      levitateAnim.current = applyLevitate(planetRef.current);
      timeout.current = setTimeout(() => {
        block.current = false;
      }, 200);
    }
  };

  return (
    <>
      <IconButton
        disabled={gameData?.currentLevel > 0 ? false : true}
        sx={{
          marginRight: "3em",
          cursor: "pointer",
          "&:hover": {
            background: "unset",
            filter: "drop-shadow(0px 0px 12px white)",
          },
        }}
      >
        <KeyboardArrowLeft
          aria-label="previous-level-button"
          onClick={() => goToPrevious()}
          sx={{
            fontSize: "30px",
          }}
        ></KeyboardArrowLeft>
      </IconButton>
      <canvas
        className="main-planet-image"
        width="200"
        height="200"
        ref={planetRef}
        onClick={click}
      ></canvas>
      <IconButton
        disabled={gameData?.currentLevel === gameData?.maxLevel ? true : false}
        sx={{
          marginLeft: "3em",
          cursor: "pointer",
          "&:hover": {
            background: "unset",
            filter: "drop-shadow(0px 0px 12px white)",
          },
        }}
      >
        <KeyboardArrowRight
          aria-label="next-level-button"
          onClick={() => goToNext()}
          sx={{
            fontSize: "30px",
          }}
        ></KeyboardArrowRight>
      </IconButton>
    </>
  );
}
