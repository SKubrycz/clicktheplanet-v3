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
  let values: number[] = [];
  let interpolated: number = 0;

  switch (random % 3) {
    case 0:
      values = createValues(
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

      interpolated = interpolateValues(values);
      return interpolated;
    case 1:
      values = createValues(
        [
          [0.004 * j, 0.004 * i],
          [0.015 * j, 0.015 * i],
          [0.025 * j, 0.025 * i],
          [0.05 * j, 0.05 * i],
          [0.1 * j, 0.1 * i],
          [0.3 * j, 0.3 * i],
        ],
        gradients,
        size
      );

      interpolated = interpolateValues(values);

      return interpolated;
    case 2:
      values = createValues(
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

      interpolated = interpolateValues(values);
      return interpolated;
    default:
      values = createValues(
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

      interpolated = interpolateValues(values);
      return interpolated;
  }
}
