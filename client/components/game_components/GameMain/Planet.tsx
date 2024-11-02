"use client";

import { useState, useRef, useEffect } from "react";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";

import { KeyboardArrowLeft, KeyboardArrowRight } from "@mui/icons-material";
import { setLevel } from "@/lib/game/levelSlice";

// import { generateGradients } from "@/utilities/simplexNoise/simplexNoise";

// import patternSelect from "@/utilities/simplexNoise/patternSelect";

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


  // TODO: Move this to Web Worker
  function draw() {
    if (planetRef.current && ctx) {
  //     const s = [];

  //     let max = 0;
  //     let min = 0;

  //     const cCenter = [
  //       planetRef.current.width / 2,
  //       planetRef.current.height / 2,
  //     ];
  //     const cRadius = planetRef.current.width / 2;

  //     const gradients = generateGradients(
  //       seed,
  //       planetRef.current?.width,
  //       planetRef?.current.height
  //     );

  //     let planetPattern = (89213 * seed + 32894789) % 5346785487;
  //     planetPattern = (89213 * planetPattern + 32894789) % 5346785487;

  //     for (let i = 0; i < planetRef.current.height * 1; i++) {
  //       for (let j = 0; j < planetRef.current.width * 1; j++) {
  //         if (
  //           Math.sqrt(
  //             Math.pow(j - cCenter[0], 2) + Math.pow(i - cCenter[1], 2)
  //           ) <=
  //           cRadius + 2
  //         ) {
  //           let final = patternSelect(planetPattern, j, i, gradients, size);

  //           if (final > max) max = final;
  //           if (final < min) min = final;

  //           s.push(final);
  //         } else {
  //           s.push(0);
  //         }
  //       }
  //     }

  //     weights.current = [
  //       ((5578 * seed * 99999 + 6785464) % 1739013789) / 1739013789,
  //       ((1117 * seed * 99999 + 4649587) % 998391231) / 998391231,
  //       ((71238 * seed * 99999 + 8956870) % 691270398) / 691270398,
  //     ];
  //     if (
  //       weights.current[0] < 0.4 &&
  //       weights.current[1] < 0.4 &&
  //       weights.current[2] < 0.4
  //     ) {
  //       weights.current[0] += 0.2;
  //       weights.current[1] += 0.2;
  //       weights.current[2] += 0.2;
  //     }

  //     let w = weights.current;

  //     let k = 0;

  //     for (let i = 0; i < imageData.data.length; i += 4) {
  //       let val = ((s[k] - min) / (max - min)) * 255;
  //       let norm = (s[k] - min) / (max - min);

  //       // if (val <= 255 && val > 220) {
  //       //   color(i, 255, 255, 255, norm);
  //       // } else if (val <= 220 && val > 160) {
  //       //   color(i, 240, 240, 240, norm);
  //       // } else if (val <= 160 && val > 150) {
  //       //   color(i, 180, 180, 180, norm);
  //       // } else if (val <= 150 && val > 90) {
  //       //   color(i, 40, 40, 40, norm);
  //       // } else if (val <= 90 && val > 70) {
  //       //   color(i, 30, 30, 30, norm);
  //       // } else {
  //       //   color(i, 20, 20, 20, norm);
  //       // }

  //       if (!gameData?.isBoss) {
  //         if (val <= 255 && val > 220) {
  //           color(imageData, i, 255 * w[0], 255 * w[1], 255 * w[2], norm);
  //         } else if (val <= 220 && val > 150) {
  //           color(imageData, i, 245 * w[0], 245 * w[1], 245 * w[2], norm);
  //         } else if (val <= 150 && val > 100) {
  //           color(imageData, i, 195 * w[0], 195 * w[1], 195 * w[2], norm);
  //         } else if (val <= 100 && val > 80) {
  //           color(imageData, i, 190 * w[1], 190 * w[2], 190 * w[0], norm);
  //         } else if (val <= 80 && val > 60) {
  //           color(imageData, i, 170 * w[1], 170 * w[2], 170 * w[0], norm);
  //         } else {
  //           color(imageData, i, 150 * w[1], 150 * w[2], 150 * w[0], norm);
  //         }
  //       } else {
  //         color(imageData, i, 255 * w[0], 255 * w[1], 255 * w[2], norm);
  //       }

  //       k++;
  //     }

      ctx.putImageData(imageData, 0, 0);
    }
  }

  useEffect(() => {
    if (planetRef?.current) {
      seed =
        gameData?.currentLevel * gameData?.currentLevel +
        gameData?.currentStage * gameData?.currentStage;
      //draw();

      workerRef.current = new Worker(new URL("../../../utilities/workers/drawWorker.ts", import.meta.url));
      workerRef.current.postMessage([planetRef.current.width, planetRef.current.height, size, seed, gameData?.isBoss, imageData]);

      workerRef.current.onmessage = (e) => {
        imageData = e.data[0];
        weights.current = e.data[1];
        draw();
      }

      breatheAnimationKeyframes.current = [
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
      ];

      if (weights.current.length > 1 && typeof weights.current[0] == "number") {
        planetRef.current.animate(breatheAnimationKeyframes.current, {
          duration: 4000,
          iterations: Infinity,
        });
      }
    }

    return () => {
      if (workerRef.current) workerRef.current.terminate();
    }
  }, [gameData.planetName]);

  const animationStyle = {
    previous: "200ms swipeRight linear 1, levitate 5.5s infinite",
    next: "200ms swipeLeft linear 1, levitate 5.5s infinite",
    destroyed: "250ms destroyed linear 1, levitate 5.5s infinite",
  };

  useEffect(() => {
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
  }, [gameData.planetsDestroyed]);

  const animatePrevious = () => {
    if (planetRef.current && gameData.currentLevel > 1 && !block.current) {
      block.current = true;
      dispatch(setLevel({ action: "previous" }));
      planetRef.current.style.animation = "none";
      planetRef.current.offsetHeight;
      planetRef.current.style.animation = animationStyle.previous;
      timeout.current = setTimeout(() => {
        block.current = false;
      }, 400);
    }
  };

  const animateNext = () => {
    if (planetRef.current && gameData.currentLevel !== gameData.maxLevel && !block.current) {
      console.log(`block: ${block}`);
      block.current = true;
      dispatch(setLevel({ action: "next" }));
      planetRef.current.style.animation = "none";
      planetRef.current.offsetHeight;
      planetRef.current.style.animation = animationStyle.next;
      timeout.current = setTimeout(() => {
        block.current = false;
      }, 400);
    }
  };

  return (
    <>
      <KeyboardArrowLeft
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
