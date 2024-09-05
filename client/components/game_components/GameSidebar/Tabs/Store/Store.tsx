import StoreElement from "./StoreElement";

import "./Store.scss";

export default function Store() {
  return (
    <div className="store-content-scroll">
      <StoreElement
        title="Element 1 title"
        description="Brief element description"
      ></StoreElement>
      <StoreElement
        title="Element 2 title"
        description="Brief element description"
      ></StoreElement>
      <StoreElement
        title="Element 3 title"
        description="Brief element description"
      ></StoreElement>
      <StoreElement
        title="Element 4 title"
        description="Brief element description"
      ></StoreElement>
      <StoreElement
        title="Element 5 title"
        description="Brief element description"
      ></StoreElement>
      <StoreElement
        title="Element 6 title"
        description="Brief element description"
      ></StoreElement>
      <StoreElement
        title="Element 7 title"
        description="Brief element description"
      ></StoreElement>
    </div>
  );
}
