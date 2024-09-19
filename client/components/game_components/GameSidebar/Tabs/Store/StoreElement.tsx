import { useContext } from "react";

import { UpgradeContext } from "@/app/game/page";
import { Add } from "@mui/icons-material";

interface StoreElementProps {
  index: number;
  title: string;
  description: string;
  image?: string;
}

export default function StoreElement({
  index,
  title,
  description,
  image,
}: StoreElementProps) {
  let upgradeFunc = useContext(UpgradeContext);

  return (
    <div className="store-element">
      <div className="store-element-left-wrapper">
        {image ? (
          <img src={image} className="store-element-image"></img>
        ) : (
          <div className="store-element-image-div">Image</div>
        )}
        <div className="store-element-info-wrapper">
          <div className="store-element-title">{title}</div>
          <div className="store-element-description">{description}</div>
        </div>
      </div>
      <Add
        onClick={() => {
          if (upgradeFunc) upgradeFunc("store", index);
        }}
        sx={{ width: 50, height: 50, cursor: "pointer" }}
      ></Add>
    </div>
  );
}
