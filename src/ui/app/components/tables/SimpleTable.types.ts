export interface SimpleTableProps {
  columns: SimpleTableColumn[];
  data: any[];
  size?: string;
  showToggleActive?: boolean;
  showDetails?: boolean;
  showEdit?: boolean;
  showDelete?: boolean;
  onDelete?: () => void;
}

export interface SimpleTableColumn {
  key: string;
  name: string;
  type: ColumnType;
  display?: boolean;
}

export enum ColumnType {
  String,
  Image,
  Number,
  Money
}

// export interface SimpleTableData {
//   value: any;
//   type: SimpleTableDataType;
// }
//
// enum SimpleTableDataType {
//   String,
//   Number,
// }
