import ShipElement from "./ShipElement";

import "./Ship.scss";

interface IShipElement {
  title: string;
  desc: string;
}

export default function Ship() {
  const shipElementsArr: IShipElement[] = [
    { title: "Dps", desc: "Damage per second based on click damage - DPS:" },
    {
      title: "Click damage",
      desc: "Boost the amount of damage dealt through clicks - Click damage:",
    },
    {
      title: "Critical click",
      desc: "Increase the chance of clicking critically",
    },
    {
      title: "Planet gold",
      desc: "Gold gained from destroying planets",
    },
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
