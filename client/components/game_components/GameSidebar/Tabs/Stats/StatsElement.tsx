"use client";

import { useEffect, useRef } from "react";

interface StatsElementProps {
  title: string;
  data: number | string | undefined;
}

export default function StatsElement({ title, data }: StatsElementProps) {
  const dataRef = useRef<HTMLDivElement>(null);
  const count = useRef<number>(0);

  useEffect(() => {
    if (count.current < 2) count.current++;
    if (dataRef.current && count.current > 1) {
      dataRef.current.style.animation = "none";
      dataRef.current.offsetHeight;
      dataRef.current.style.animation = "200ms flash 1";
    }
  }, [data]);

  return (
    <div className="stats-element">
      <div className="stats-element-left">{title}</div>
      <div ref={dataRef} className="stats-element-right">
        {data}
      </div>
    </div>
  );
}
