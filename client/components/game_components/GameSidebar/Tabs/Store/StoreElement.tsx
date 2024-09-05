interface StoreElementProps {
  title: string;
  description: string;
}

export default function StoreElement({
  title,
  description,
}: StoreElementProps) {
  return (
    <div className="store-element">
      <div className="store-element-info-wrapper">
        <div className="store-element-title">{title}</div>
        <div className="store-element-description">{description}</div>
      </div>
      <div>Action buttons?</div>
    </div>
  );
}
