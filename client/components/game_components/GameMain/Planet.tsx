"use client";

import { useState, useRef, useEffect } from "react";

import { useAppSelector, useAppDispatch } from "@/lib/hooks";

import { KeyboardArrowLeft, KeyboardArrowRight } from "@mui/icons-material";
import { setLevel } from "@/lib/game/levelSlice";

import {
  generateGradients,
  simplexNoise,
} from "@/utilities/simplexNoise/simplexNoise";

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

  function avg(a: number, b: number) {
    return (a + b) / 2;
  }

  function color(i: number, r: number, g: number, b: number, val = 1) {
    let rAvg = avg(r, r * val);
    let gAvg = avg(g, g * val);
    let bAvg = avg(b, b * val);

    imageData.data[i] = rAvg;
    imageData.data[i + 1] = gAvg;
    imageData.data[i + 2] = bAvg;
    imageData.data[i + 3] = 255;
  }

  function draw() {
    if (planetRef.current && ctx) {
      const s = [];

      let max = 0;
      let min = 0;

      const cCenter = [
        planetRef.current.width / 2,
        planetRef.current.height / 2,
      ];
      const cRadius = planetRef.current.width / 2;

      const gradients = generateGradients(
        seed,
        planetRef.current?.width,
        planetRef?.current.height
      );

      //console.log(imageData);

      for (let i = 0; i < planetRef.current.height * 1; i++) {
        for (let j = 0; j < planetRef.current.width * 1; j++) {
          if (
            Math.sqrt(
              Math.pow(j - cCenter[0], 2) + Math.pow(i - cCenter[1], 2)
            ) <=
            cRadius + 2
          ) {
            // const value1 = simplexNoise(0.005 * j, 0.005 * i, gradients, size);
            // const value2 = simplexNoise(0.01 * j, 0.01 * i, gradients, size);
            // const value3 = simplexNoise(0.05 * j, 0.05 * i, gradients, size);
            // const value4 = simplexNoise(0.025 * j, 0.025 * i, gradients, size);

            // let interpolate = value1 + (value2 - value1) * 0.25;
            // interpolate = interpolate + (value3 - interpolate) * 0.05;
            // interpolate = interpolate + (value4 - interpolate) * 0.15;

            const value1 = simplexNoise(0.005 * j, 0.005 * i, gradients, size);
            const value2 = simplexNoise(0.01 * j, 0.01 * i, gradients, size);
            const value3 = simplexNoise(0.015 * j, 0.015 * i, gradients, size);
            const value4 = simplexNoise(0.05 * j, 0.05 * i, gradients, size);
            const value5 = simplexNoise(0.1 * j, 0.1 * i, gradients, size);
            const value6 = simplexNoise(0.25 * j, 0.25 * i, gradients, size);

            let interpolate = value1 + (value2 - value1) * 0.5;
            interpolate = interpolate + (value3 - interpolate) * 0.25;
            interpolate = interpolate + (value4 - interpolate) * 0.1;
            interpolate = interpolate + (value5 - interpolate) * 0.05;
            interpolate = interpolate + (value6 - interpolate) * 0.04;

            let final = interpolate;

            if (final > max) max = final;
            if (final < min) min = final;

            s.push(final);
          } else {
            s.push(0);
          }
        }
      }

      weights.current = [
        ((5578 * seed * 99999 + 6785464) % 1739013789) / 1739013789,
        ((1117 * seed * 99999 + 4649587) % 998391231) / 998391231,
        ((71238 * seed * 99999 + 8956870) % 691270398) / 691270398,
      ];

      let w = weights.current;

      let k = 0;

      for (let i = 0; i < imageData.data.length; i += 4) {
        let val = ((s[k] - min) / (max - min)) * 255;
        let norm = (s[k] - min) / (max - min);

        //if (i === 1800) console.log(255 * w[0]);

        // if (val <= 255 && val > 220) {
        //   color(i, 255, 255, 255, norm);
        // } else if (val <= 220 && val > 160) {
        //   color(i, 240, 240, 240, norm);
        // } else if (val <= 160 && val > 150) {
        //   color(i, 180, 180, 180, norm);
        // } else if (val <= 150 && val > 90) {
        //   color(i, 40, 40, 40, norm);
        // } else if (val <= 90 && val > 70) {
        //   color(i, 30, 30, 30, norm);
        // } else {
        //   color(i, 20, 20, 20, norm);
        // }

        if (val <= 255 && val > 220) {
          color(i, 255 * w[0], 255 * w[1], 255 * w[2], norm);
        } else if (val <= 220 && val > 150) {
          color(i, 240 * w[0], 240 * w[1], 240 * w[2], norm);
        } else if (val <= 150 && val > 100) {
          color(i, 180 * w[0], 180 * w[1], 180 * w[2], norm);
        } else if (val <= 100 && val > 80) {
          color(i, 170 * w[1], 170 * w[2], 170 * w[0], norm);
        } else if (val <= 80 && val > 60) {
          color(i, 150 * w[1], 150 * w[2], 150 * w[0], norm);
        } else {
          color(i, 130 * w[1], 130 * w[2], 130 * w[0], norm);
        }

        k++;
      }

      ctx.putImageData(imageData, 0, 0);
    }
  }

  useEffect(() => {
    if (planetRef?.current) {
      seed =
        gameData?.currentLevel * gameData?.currentLevel +
        gameData?.currentStage * gameData?.currentStage;
      draw();

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
  }, [gameData.planetName]);

  const animationStyle = {
    // breathe 4s infinite
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
    if (planetRef.current && gameData.currentLevel > 1) {
      planetRef.current.style.animation = "none";
      planetRef.current.offsetHeight;
      planetRef.current.style.animation = animationStyle.previous;
    }
  };

  const animateNext = () => {
    if (planetRef.current && gameData.currentLevel !== gameData.maxLevel) {
      planetRef.current.style.animation = "none";
      planetRef.current.offsetHeight;
      planetRef.current.style.animation = animationStyle.next;
    }
  };

  return (
    <>
      <KeyboardArrowLeft
        onClick={() => {
          dispatch(setLevel({ action: "previous" }));
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
          dispatch(setLevel({ action: "next" }));
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
