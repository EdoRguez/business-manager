"use client";

import Card from "./components/cards/Card";

export default function Home() {

  const dummyData: number[] = Array.from(Array(7).keys());

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
        {
          dummyData.map((x: number) =>  (
            <div key={x}>
                <Card />
              </div>
            )
          )
        }
      </div>
    </div>
  );
}
