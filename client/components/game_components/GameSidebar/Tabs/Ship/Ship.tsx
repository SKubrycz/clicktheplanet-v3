import ShipElement from "./ShipElement";

import "./Ship.scss";

export default function Ship() {
  return (
    <div className="ship-content-scroll">
      <ShipElement
        title="Element 1"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 2"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 3"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 4"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 5"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 6"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 7"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
      <ShipElement
        title="Element 8"
        description="Lorem ipsum dolor sit amet"
      ></ShipElement>
    </div>
  );
}
