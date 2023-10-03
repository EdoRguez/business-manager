import { Table } from "antd";

export type EditableTableProps = Parameters<typeof Table>[0];

export type ColumnTypes = Exclude<EditableTableProps['columns'], undefined>;

export interface DataType {
  key: React.Key;
  name: string;
  age: string;
  address: string;
}