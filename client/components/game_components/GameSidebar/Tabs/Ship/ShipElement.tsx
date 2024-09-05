interface ShipElementProps {
  title: string;
  description: string;
}

export default function ShipElement({ title, description }: ShipElementProps) {
  return (
    <div className="ship-element">
      <div className="ship-element-info-wrapper">
        <div className="ship-element-title">{title}</div>
        <div className="ship-element-description">{description}</div>
      </div>
      <div className="ship-element-action-wrapper">
        <div>Upgrade</div>
      </div>
    </div>
  );
}
