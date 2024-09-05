import Store from "./Store/Store";
import Ship from "./Ship/Ship";
import Stats from "./Stats/Stats";

interface TabProps {
  tabTitle: string;
}

export default function Tabs({ tabTitle }: TabProps) {
  switch (tabTitle) {
    case "Store":
      return <Store></Store>;
    case "Ship":
      return <Ship></Ship>;
    case "Stats":
      return <Stats></Stats>;
    default:
      return <></>;
  }
}
