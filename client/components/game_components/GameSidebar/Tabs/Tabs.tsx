import Store from "./Store/Store";

interface TabProps {
  tabTitle: string;
}

export default function Tabs({ tabTitle }: TabProps) {
  switch (tabTitle) {
    case "Store":
      return <Store></Store>;
    default:
      return <></>;
  }
}
