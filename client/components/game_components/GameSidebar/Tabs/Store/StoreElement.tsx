import { Add } from "@mui/icons-material";

interface StoreElementProps {
  title: string;
  description: string;
  image?: string;
}

export default function StoreElement({
  title,
  description,
  image,
}: StoreElementProps) {
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
      <Add sx={{ width: 50, height: 50, cursor: "pointer" }}></Add>
    </div>
  );
}
