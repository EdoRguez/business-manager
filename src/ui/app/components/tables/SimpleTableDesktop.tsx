'use client';

import { SimpleTableColumn, SimpleTableProps } from './SimpleTable.types';
import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
  Button,
  TableCaption,
  Switch
} from '@chakra-ui/react'
import { Icon } from '@iconify/react';
import SimpleTableField from './SimpleTableField';
import { useState } from 'react';

const SimpleTableDesktop: React.FC<SimpleTableProps> = ({
  columns,
  data,
  offset,
  size = 'md',
  showToggleActive = false,
  showDetails = false,
  showEdit = false,
  showDelete = false,
  onDelete,
  onChangePage
}) => {

  const [pageNumber, setPageNumber] = useState<number>(1);

  const handleChangePage = (val: string) => {
    setPageNumber((prevValue) => val === 'NEXT' ? ++prevValue : --prevValue);

    if (onChangePage) {
      return onChangePage(val);
    }
  }

  return (
    <TableContainer>
      <Table size={size}>
        <TableCaption>
          <div className='flex justify-end'>
            <div className='flex items-center select-none text-thirdcolor'>
              {
                offset >= 10 &&
                <div className='cursor-pointer p-2 rounded hover:bg-maincolor duration-150 hover:text-white' onClick={() => handleChangePage('PREV')}>
                  <Icon icon="fa:chevron-left" />
                </div>
              }
              <span className='mx-2 font-bold'>
                Página {pageNumber}
              </span>
              {
                data.length !== 0 &&
                <div className='cursor-pointer p-2 rounded hover:bg-maincolor duration-150 hover:text-white' onClick={() => handleChangePage('NEXT')}>
                  <Icon icon="fa:chevron-right" />
                </div>
              }
            </div>
          </div>
        </TableCaption>
        <Thead>
          <Tr>
            {columns.map((col: SimpleTableColumn, index: number) => (
              <Th key={index}>
                {col.name}
              </Th>
            ))}

            {(showEdit || showDelete) &&
              <Th></Th>
            }
          </Tr>
        </Thead>
        <Tbody>
          {data.map((dataVal: any, dataIndex: number) => (

            <Tr key={dataIndex} className='hover:bg-thirdcolorhov transition duration-150 text-sm'>
              {columns.map((col: SimpleTableColumn, colIndex: number) => (

                <Td key={colIndex}>
                  {
                    <SimpleTableField data={dataVal} col={col} />
                  }
                </Td>

              ))}

              <Td>
                <div className='flex justify-end'>

                  {showToggleActive && (
                    <div className='mr-2 flex items-center justify-center'>
                      <span className="text-xs font-bold text-maincolor mr-2">Activo</span>
                      <Switch size='md' colorScheme='main' />
                    </div>
                  )}

                  {showDetails && (
                    <Button size="sm" variant="fifth">
                      <Icon icon="lucide:info" />
                    </Button>
                  )}

                  {showEdit && (
                    <Button size="sm" variant="main" className="mx-1">
                      <Icon icon="lucide:edit" />
                    </Button>
                  )}

                  {showDelete && (
                    <Button size="sm" variant="third" onClick={onDelete}>
                      <Icon icon="wpf:delete" />
                    </Button>
                  )}
                </div>
              </Td>

            </Tr>

          ))}

          {
            data.length === 0 && (
              <Tr>
                <Td colSpan={6} >
                  <div className='flex justify-center'>

                  <span>No hay registros para mostrar</span>
                  </div>
                </Td>
              </Tr>
            )
          }

        </Tbody>
      </Table>
    </TableContainer>
  );
}

export default SimpleTableDesktop;
