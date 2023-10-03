"use client";

import { DataType } from "@/app/common/types";
import TableImage from "@/app/components/tables/TableImage";
import { Popconfirm } from "antd";

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

  return <TableImage columns={tableColumns} data={tableData} />;
};

export default OrderClient;
