import { Table } from "antd";
import { IconType } from "antd/es/notification/interface";

export type EditableTableProps = Parameters<typeof Table>[0];

export type ColumnTypes = Exclude<EditableTableProps['columns'], undefined>;

export interface DataType {
  key: React.Key;
  name: string;
  age: string;
  address: string;
}

export interface MenuElement {
  title: string;
  path: string;
  icon: IconType;
  gap?: boolean;
}