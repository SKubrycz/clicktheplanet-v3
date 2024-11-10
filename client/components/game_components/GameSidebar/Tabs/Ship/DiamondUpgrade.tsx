import Image from "next/image";
import Diamond from "@/assets/svg/diamond.svg";

import { Typography } from "@mui/material";

interface DiamondUpgradeProps {
  title: string;
  index: number;
}

export default function DiamondUpgrade({ title, index }: DiamondUpgradeProps) {
  return (
    <div className="diamond-upgrade">
      <div>
        <Typography variant="h6" sx={{ fontSize: "18px" }}>
          {title}
        </Typography>
        <div>
          Cost:{" "}
          <Image src={Diamond} alt="diamond" width={14} height={14}></Image>
        </div>
      </div>
      <div>U</div>
    </div>
  );
}
