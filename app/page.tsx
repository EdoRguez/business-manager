"use client";

import Card from "./components/card/Card";

export default function Home() {
  return (
    <div>
      <div
        className="
            mt-24
            grid 
            grid-cols-1 
            sm:grid-cols-2 
            2xl:grid-cols-4
            gap-8
          "
      >
        <div>
          <Card />
        </div>

        <div>
          <Card />
        </div>

        <div>
          <Card />
        </div>

        <div>
          <Card />
        </div>

        <div>
          <Card />
        </div>

        <div>
          <Card />
        </div>

        <div>
          <Card />
        </div>
      </div>
    </div>
  );
}
