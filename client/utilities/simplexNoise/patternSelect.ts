import { simplexNoise } from "./simplexNoise";

function createValues(
  values: number[][],
  gradients: number[][],
  size: number
): number[] {
  let noises: number[] = [];

  values.forEach((el, i) => {
    const value = simplexNoise(el[0], el[1], gradients, size);
    noises.push(value);
  });

  return noises;
}

function interpolateValues(noises: number[]): number {
  const weights = [0.5, 0.25, 0.1, 0.05, 0.04, 0.02, 0.01];

  let interpolate = noises[0] + (noises[1] - noises[0]) * weights[0];

  noises.forEach((el, i) => {
    if (i > 2) {
      interpolate =
        interpolate +
        (noises[i] - interpolate) * weights[(i - 1) % weights.length];
    }
  });

  return interpolate;
}

export default function generatePattern(
  random: number,
  j: number,
  i: number,
  gradients: number[][],
  size: number
): number {
  if (random % 3 === 0) {
    const values = createValues(
      [
        [0.005 * j, 0.005 * i],
        [0.01 * j, 0.01 * i],
        [0.015 * j, 0.015 * i],
        [0.05 * j, 0.05 * i],
        [0.1 * j, 0.1 * i],
        [0.25 * j, 0.25 * i],
      ],
      gradients,
      size
    );

    const interpolated = interpolateValues(values);

    // const value1 = simplexNoise(0.005 * j, 0.005 * i, gradients, size);
    // const value2 = simplexNoise(0.01 * j, 0.01 * i, gradients, size);
    // const value3 = simplexNoise(0.015 * j, 0.015 * i, gradients, size);
    // const value4 = simplexNoise(0.05 * j, 0.05 * i, gradients, size);
    // const value5 = simplexNoise(0.1 * j, 0.1 * i, gradients, size);
    // const value6 = simplexNoise(0.25 * j, 0.25 * i, gradients, size);

    // let interpolate = value1 + (value2 - value1) * 0.5;
    // interpolate = interpolate + (value3 - interpolate) * 0.25;
    // interpolate = interpolate + (value4 - interpolate) * 0.1;
    // interpolate = interpolate + (value5 - interpolate) * 0.05;
    // interpolate = interpolate + (value6 - interpolate) * 0.04;

    return interpolated;
  } else if (random % 3 === 1) {
    const value1 = simplexNoise(0.004 * j, 0.004 * i, gradients, size);
    const value2 = simplexNoise(0.015 * j, 0.015 * i, gradients, size);
    const value3 = simplexNoise(0.025 * j, 0.025 * i, gradients, size);
    const value4 = simplexNoise(0.05 * j, 0.05 * i, gradients, size);
    const value5 = simplexNoise(0.1 * j, 0.1 * i, gradients, size);
    const value6 = simplexNoise(0.3 * j, 0.3 * i, gradients, size);

    let interpolate = value1 + (value2 - value1) * 0.5;
    interpolate = interpolate + (value3 - interpolate) * 0.25;
    interpolate = interpolate + (value4 - interpolate) * 0.1;
    interpolate = interpolate + (value5 - interpolate) * 0.05;
    interpolate = interpolate + (value6 - interpolate) * 0.04;

    return interpolate;
  } else if (random % 3 === 2) {
    const value1 = simplexNoise(0.005 * j, 0.005 * i, gradients, size);
    const value2 = simplexNoise(0.01 * j, 0.01 * i, gradients, size);
    const value3 = simplexNoise(0.015 * j, 0.015 * i, gradients, size);
    const value4 = simplexNoise(0.05 * j, 0.05 * i, gradients, size);
    const value5 = simplexNoise(0.1 * j, 0.1 * i, gradients, size);
    const value6 = simplexNoise(0.25 * j, 0.25 * i, gradients, size);

    let interpolate = value1 + (value2 - value1) * 0.25;
    interpolate = interpolate + (value3 - interpolate) * 0.5;
    interpolate = interpolate + (value4 - interpolate) * 0.1;
    interpolate = interpolate + (value5 - interpolate) * 0.02;
    interpolate = interpolate + (value6 - interpolate) * 0.05;

    return interpolate;
  } else {
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

    return interpolate;
  }
}
