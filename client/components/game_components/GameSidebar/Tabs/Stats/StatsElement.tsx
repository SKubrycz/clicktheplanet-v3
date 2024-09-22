interface StatsElementProps {
  title: string; // maybe set title and description
  data: number | string | undefined;
}

export default function StatsElement({ title, data }: StatsElementProps) {
  return (
    <div className="stats-element">
      <div className="stats-element-left">{title}</div>
      <div className="stats-element-right">{data}</div>
    </div>
  );
}
