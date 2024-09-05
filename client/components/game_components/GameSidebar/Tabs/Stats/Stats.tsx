import StatsElement from "./StatsElement";

import "./Stats.scss";

export default function Stats() {
  return (
    <div className="stats-content">
      <StatsElement title="Current damage:" data={5}></StatsElement>
      <StatsElement title="Highest level ever:" data={16}></StatsElement>
      <StatsElement title="Planets destroyed:" data={207}></StatsElement>
    </div>
  );
}
