import ShipElement from "./ShipElement";

import "./Ship.scss";

interface IShipElement {
  title: string;
  desc: string;
}

export default function Ship() {
  const shipElementsArr: IShipElement[] = [
    { title: "Dps", desc: "Damage per second based on click damage" },
    { title: "Element 2", desc: "Lorem ipsum dolor sit amet" },
    { title: "Element 3", desc: "Lorem ipsum dolor sit amet" },
    { title: "Element 4", desc: "Lorem ipsum dolor sit amet" },
  ];

  return (
    <div className="ship-content-scroll">
      {shipElementsArr.map((el, i) => {
        return (
          <ShipElement
            key={i}
            index={i + 1}
            title={el.title}
            description={el.desc}
          ></ShipElement>
        );
      })}
    </div>
  );
}
