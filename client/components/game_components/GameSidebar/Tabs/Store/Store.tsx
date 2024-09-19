import StoreElement from "./StoreElement";

import "./Store.scss";

interface IStoreElement {
  title: string;
  desc: string;
}

export default function Store() {
  const storeElementsArr: IStoreElement[] = [
    { title: "Element 1 title", desc: "Brief element description" },
    { title: "Element 2 title", desc: "Brief element description" },
    { title: "Element 3 title", desc: "Brief element description" },
    { title: "Element 4 title", desc: "Brief element description" },
    { title: "Element 5 title", desc: "Brief element description" },
    { title: "Element 6 title", desc: "Brief element description" },
    { title: "Element 7 title", desc: "Brief element description" },
  ];

  return (
    <div className="store-content-scroll">
      {storeElementsArr.map((el, i) => {
        return (
          <StoreElement
            key={i}
            index={i}
            title={el.title}
            description={el.desc}
          ></StoreElement>
        );
      })}
    </div>
  );
}
