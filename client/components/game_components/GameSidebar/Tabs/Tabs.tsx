import Store from "./Store/Store";
import Ship from "./Ship/Ship";

interface TabProps {
  tabTitle: string;
}

export default function Tabs({ tabTitle }: TabProps) {
  switch (tabTitle) {
    case "Store":
      return <Store></Store>;
    case "Ship":
      return <Ship></Ship>;
    default:
      return <></>;
  }
}
