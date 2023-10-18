"use client";

import { DataType } from "@/app/common/types";
import CardAmount from "@/app/components/cards/CardAmount";
import TableImage from "@/app/components/tables/TableImage";
import { GiPayMoney, GiPoloShirt, GiReceiveMoney } from "react-icons/gi";
import { Popconfirm } from "antd";
import { AmountType } from "@/app/common/enums";
import { BsBoxSeamFill } from "react-icons/bs";
import { FaPlaneDeparture } from "react-icons/fa";
import { BiSolidBadgeDollar } from "react-icons/bi";
import { IoIosMan, IoIosWoman } from "react-icons/io";

const OrderClient = () => {
  const tableColumns: any[] = [
    {
      title: "Imagen",
      dataIndex: "image",
      render: (_: any, record: { key: React.Key }) => {
        return <h1>Hola</h1>;
      },
    },
    {
      title: "Nombre",
      dataIndex: "name",
      editable: true,
    },
    {
      title: "Cantidad",
      dataIndex: "quantity",
      editable: true,
    },
    {
      title: "Costo",
      dataIndex: "cost",
      editable: true,
    },
    {
      title: "Total Costo",
      dataIndex: "totalCost",
      editable: true,
    },
    {
      title: "Venta",
      dataIndex: "sell",
      editable: true,
    },
    {
      title: "Total Venta",
      dataIndex: "totalSell",
      editable: true,
    },
    {
      title: "operation",
      dataIndex: "operation",
      render: (_: any, record: { key: React.Key }) => (
        <Popconfirm title="Deseas borrar el registro?" onConfirm={() => {}}>
          <a>Borrar</a>
        </Popconfirm>
      ),
    },
  ];

  const tableData: DataType[] = [
    {
      key: "0",
      name: "Edward King 0",
      age: "32",
      address: "London, Park Lane no. 0",
    },
    {
      key: "1",
      name: "Edward King 1",
      age: "32",
      address: "London, Park Lane no. 1",
    },
  ];

  return (
    <div
      className="
        grid
        grid-cols-1
        sm:grid-cols-4 
        2xl:grid-cols-6
        gap-4
      "
    >
      <div className="col-start-1 col-span-4">
        <TableImage columns={tableColumns} data={tableData} />
      </div>
      <div className="col-start-5 col-span-2">
        <div
          className="
              grid
              grid-cols-1
              sm:grid-cols-2 
              gap-2
            "
        >
          <CardAmount
            icon={BsBoxSeamFill}
            iconColor="text-blue-800"
            iconBackgroundColor="bg-blue-300"
            title="Gastos Ropa"
            amount={100}
            amountType={AmountType.Money}
          />
          <CardAmount
            icon={FaPlaneDeparture}
            iconColor="text-red-800"
            iconBackgroundColor="bg-blue-300"
            title="Gastos Envío"
            amount={20}
            amountType={AmountType.Money}
          />
          <CardAmount
            icon={GiPayMoney}
            iconColor="text-red-800"
            iconBackgroundColor="bg-blue-300"
            title="Gastos Ropa + Envío"
            amount={120}
            amountType={AmountType.Money}
          />
          <CardAmount
            icon={GiReceiveMoney}
            iconColor="text-red-800"
            iconBackgroundColor="bg-blue-300"
            title="Total a Vender"
            amount={220}
            amountType={AmountType.Money}
          />
          <div className="col-start-1 col-span-2">
            <CardAmount
              icon={BiSolidBadgeDollar}
              iconColor="text-red-800"
              iconBackgroundColor="bg-blue-300"
              title="Ganancias Totales"
              amount={100}
              amountType={AmountType.Money}
            />
          </div>
          <CardAmount
            icon={IoIosMan}
            iconColor="text-red-800"
            iconBackgroundColor="bg-blue-300"
            title="Ganancias Eduardo"
            amount={30}
            amountType={AmountType.Money}
          />
          <CardAmount
            icon={IoIosWoman}
            iconColor="text-red-800"
            iconBackgroundColor="bg-blue-300"
            title="Ganancias Daniela"
            amount={70}
            amountType={AmountType.Money}
          />
        </div>
      </div>
    </div>
  );
};

export default OrderClient;
