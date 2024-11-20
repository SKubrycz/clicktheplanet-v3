import { generateGradients } from "../simplexNoise/simplexNoise";
import patternSelect from "../simplexNoise/patternSelect";

function avg(a: number, b: number) {
  return (a + b) / 2;
}

function color(
  imageData: ImageData,
  i: number,
  r: number,
  g: number,
  b: number,
  val = 1
) {
  let rAvg = avg(r, r * val);
  let gAvg = avg(g, g * val);
  let bAvg = avg(b, b * val);

  imageData.data[i] = rAvg;
  imageData.data[i + 1] = gAvg;
  imageData.data[i + 2] = bAvg;
  imageData.data[i + 3] = 255;
}

let size = 1;
let imageData: ImageData;
let weights: number[] = [];

let width = 200;
let height = 200;
let seed = 123;
let isBoss = false;

self.onmessage = (msg) => {
  width = msg.data[0];
  height = msg.data[1];
  size = msg.data[2];
  seed = msg.data[3];
  isBoss = msg.data[4];
  imageData = msg.data[5];

  draw();

  postMessage([imageData, weights]);
};

function draw() {
  const s = [];

  let max = 0;
  let min = 0;

  const cCenter = [width / 2, height / 2];
  const cRadius = width / 2;

  const gradients = generateGradients(seed, width, height);

  let planetPattern = (89213 * seed + 32894789) % 5346785487;
  planetPattern = (91235 * planetPattern + 56426456) % 78269681123;
  console.log(planetPattern % 3);

  for (let i = 0; i < height * 1; i++) {
    for (let j = 0; j < width * 1; j++) {
      if (
        Math.sqrt(Math.pow(j - cCenter[0], 2) + Math.pow(i - cCenter[1], 2)) <=
        cRadius + 2
      ) {
        let final = patternSelect(planetPattern, j, i, gradients, size);

        if (final > max) max = final;
        if (final < min) min = final;

        s.push(final);
      } else {
        s.push(0);
      }
    }
  }

  weights = [
    ((5578 * seed * 99999 + 6785464) % 1739013789) / 1739013789,
    ((1117 * seed * 99999 + 4649587) % 998391231) / 998391231,
    ((71238 * seed * 99999 + 8956870) % 691270398) / 691270398,
  ];
  if (weights[0] < 0.4 && weights[1] < 0.4 && weights[2] < 0.4) {
    weights[0] += 0.2;
    weights[1] += 0.2;
    weights[2] += 0.2;
  }

  let w = weights;

  let k = 0;

  for (let i = 0; i < imageData.data.length; i += 4) {
    let val = ((s[k] - min) / (max - min)) * 255;
    let norm = (s[k] - min) / (max - min);

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

    if (!isBoss) {
      if (planetPattern % 3 === 0) {
        if (val <= 255 && val > 220) {
          color(imageData, i, 255 * w[0], 255 * w[1], 255 * w[2], norm);
        } else if (val <= 220 && val > 150) {
          color(imageData, i, 245 * w[0], 245 * w[1], 245 * w[2], norm);
        } else if (val <= 150 && val > 100) {
          color(imageData, i, 195 * w[0], 195 * w[1], 195 * w[2], norm);
        } else if (val <= 100 && val > 80) {
          color(imageData, i, 190 * w[1], 190 * w[2], 190 * w[0], norm);
        } else if (val <= 80 && val > 60) {
          color(imageData, i, 170 * w[1], 170 * w[2], 170 * w[0], norm);
        } else {
          color(imageData, i, 150 * w[1], 150 * w[2], 150 * w[0], norm);
        }
      } else if (planetPattern % 3 === 1) {
        if (val <= 255 && val > 210) {
          color(imageData, i, 255 * w[0], 255 * w[1], 255 * w[2], norm);
        } else if (val <= 210 && val > 135) {
          color(imageData, i, 240 * w[0], 240 * w[1], 240 * w[2], norm);
        } else if (val <= 135 && val > 100) {
          color(imageData, i, 210 * w[2], 210 * w[0], 210 * w[1], norm);
        } else if (val <= 100 && val > 80) {
          color(imageData, i, 200 * w[2], 200 * w[0], 200 * w[1], norm);
        } else if (val <= 80 && val > 50) {
          color(imageData, i, 190 * w[2], 190 * w[0], 190 * w[1], norm);
        } else {
          color(imageData, i, 160 * w[2], 160 * w[0], 160 * w[1], norm);
        }
      } else {
        if (val <= 255 && val > 190) {
          color(imageData, i, 255 * w[2], 255 * w[1], 255 * w[0], norm);
        } else if (val <= 190 && val > 120) {
          color(imageData, i, 240 * w[2], 240 * w[1], 240 * w[0], norm);
        } else if (val <= 130 && val > 100) {
          color(imageData, i, 230 * w[2], 230 * w[1], 230 * w[0], norm);
        } else if (val <= 100 && val > 80) {
          color(imageData, i, 205 * w[0], 205 * w[1], 205 * w[2], norm);
        } else if (val <= 80 && val > 50) {
          color(imageData, i, 200 * w[0], 200 * w[1], 200 * w[2], norm);
        } else {
          color(imageData, i, 180 * w[0], 180 * w[1], 180 * w[2], norm);
        }
      }
    } else {
      color(imageData, i, 255 * w[0], 255 * w[1], 255 * w[2], norm);
    }

    k++;
  }
}
