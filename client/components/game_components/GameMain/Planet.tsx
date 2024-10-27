"use client";

import { useRef, useEffect } from "react";

import { useAppSelector } from "@/lib/hooks";

import {
  generateGradients,
  simplexNoise,
} from "@/utilities/simplexNoise/simplexNoise";

interface PlanetProps {
  planetRef: React.RefObject<HTMLCanvasElement>;
  click: (e: React.MouseEvent<HTMLCanvasElement>) => void;
}

export default function Planet({ planetRef, click }: PlanetProps) {
  const gameData = useAppSelector((state) => state.game);

  let seed =
    gameData?.currentLevel * gameData?.currentLevel +
    gameData?.currentStage * gameData?.currentStage;
  let size = 1;
  let ctx: CanvasRenderingContext2D | null;
  let imageData: ImageData;

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
            const value1 = simplexNoise(0.005 * j, 0.005 * i, gradients, size);
            const value2 = simplexNoise(0.01 * j, 0.01 * i, gradients, size);
            const value3 = simplexNoise(0.05 * j, 0.05 * i, gradients, size);
            const value4 = simplexNoise(0.025 * j, 0.025 * i, gradients, size);

            let interpolate = value1 + (value2 - value1) * 0.25;
            interpolate = interpolate + (value3 - interpolate) * 0.05;
            interpolate = interpolate + (value4 - interpolate) * 0.1;

            let final = interpolate;

            if (final > max) max = final;
            if (final < min) min = final;

            s.push(final);
          } else {
            s.push(0);
          }
        }
      }

      let w = [
        ((5578 * seed * 99999 + 6785464) % 1739013789) / 1739013789,
        ((1117 * seed * 99999 + 4649587) % 998391231) / 998391231,
        ((71238 * seed * 99999 + 8956870) % 691270398) / 691270398,
      ];

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
        } else if (val <= 220 && val > 160) {
          color(i, 240 * w[0], 240 * w[1], 240 * w[2], norm);
        } else if (val <= 160 && val > 150) {
          color(i, 180 * w[0], 180 * w[1], 180 * w[2], norm);
        } else if (val <= 150 && val > 90) {
          color(i, 40 * w[0], 40 * w[1], 40 * w[2], norm);
        } else if (val <= 90 && val > 70) {
          color(i, 30 * w[0], 30 * w[1], 30 * w[2], norm);
        } else {
          color(i, 20 * w[0], 20 * w[1], 20 * w[2], norm);
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
    }
  }, [gameData.planetName]);

  return (
    <>
      <canvas
        className="main-planet-image"
        width="200"
        height="200"
        ref={planetRef}
        onClick={click}
      ></canvas>
    </>
  );
}
