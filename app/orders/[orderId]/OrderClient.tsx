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
import { IoIosMan, IoIosWoman, IoMdAddCircle } from "react-icons/io";
import Button from "@/app/components/Button";
import ModalAddProductOrder from "@/app/components/modals/ModalAddProductOrder";
import { useState } from "react";
import useModalAddProductOrder from "@/app/hooks/useModalAddProductOrder";

const OrderClient = () => {
  const modalAddProductOrder = useModalAddProductOrder();

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
    <div>
      <ModalAddProductOrder title="Agregar Producto"/>

      <div
        className="
          grid
          grid-cols-1
          2xl:grid-cols-7
          gap-4
      "
      >

        <div className="2xl:col-start-1 2xl:col-span-5">
          <div className="flex justify-end mb-3">
            <div className="w-56">
              <Button
                label="Agregar Producto"
                icon={IoMdAddCircle}
                onClick={modalAddProductOrder.onOpen}
              />
            </div>
          </div>
          <TableImage columns={tableColumns} data={tableData} />
        </div>

        <div className="2xl:col-start-6 2xl:col-span-2">
          <div
            className="
                grid
                grid-cols-2
                gap-2
              "
          >
            <CardAmount
              icon={BsBoxSeamFill}
              iconColor="text-stone-900"
              iconBackgroundColor="bg-stone-200"
              title="Gastos Ropa"
              amount={100}
              amountType={AmountType.Money}
            />
            <CardAmount
              icon={FaPlaneDeparture}
              iconColor="text-stone-900"
              iconBackgroundColor="bg-stone-200"
              title="Gastos Envío"
              amount={20}
              amountType={AmountType.Money}
            />
            <CardAmount
              icon={GiPayMoney}
              iconColor="text-rose-900"
              iconBackgroundColor="bg-rose-200"
              title="Gastos Ropa + Envío"
              amount={120}
              amountType={AmountType.Money}
            />
            <CardAmount
              icon={GiReceiveMoney}
              iconColor="text-green-900"
              iconBackgroundColor="bg-green-200"
              title="Total a Vender"
              amount={220}
              amountType={AmountType.Money}
            />
            <div className="col-start-1 col-span-2">
              <CardAmount
                icon={BiSolidBadgeDollar}
                iconColor="text-yellow-600"
                iconBackgroundColor="bg-yellow-100"
                title="Ganancias Totales"
                amount={100}
                amountType={AmountType.Money}
              />
            </div>
            <CardAmount
              icon={IoIosMan}
              iconColor="text-sky-900"
              iconBackgroundColor="bg-sky-200"
              title="Ganancias Eduardo"
              amount={30}
              amountType={AmountType.Money}
            />
            <CardAmount
              icon={IoIosWoman}
              iconColor="text-fuchsia-900"
              iconBackgroundColor="bg-fuchsia-200"
              title="Ganancias Daniela"
              amount={70}
              amountType={AmountType.Money}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default OrderClient;
