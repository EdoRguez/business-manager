'use client';

import { GetCustomersRequest } from '@/app/api/customers/route';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import SimpleCard from '@/app/components/cards/SimpleCard';
import DeleteModal from '@/app/components/modals/DeleteModal';
import SimpleTable from '@/app/components/tables/SimpleTable';
import { ColumnType, SimpleTableColumn } from '@/app/components/tables/SimpleTable.types';
import useDeleteModal from '@/app/hooks/useDeleteModal';
import { BreadcrumItem } from '@/app/types';
import { Customer } from '@/app/types/customer';
import { Button, Input } from '@chakra-ui/react';
import { Icon } from '@iconify/react';
import Link from "next/link";
import { useCallback, useEffect, useState } from 'react';

const CustomersClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Clientes",
      href: "/management/customers"
    }
  ];

  const customerCols: SimpleTableColumn[] = [
    {
      key: "identificationNumber",
      name: "Cédula",
      type: ColumnType.String
    },
    {
      key: "firstName",
      name: "Nombre",
      type: ColumnType.String
    },
    {
      key: "lastName",
      name: "Apellido",
      type: ColumnType.String
    },
    {
      key: "email",
      name: "Correo",
      type: ColumnType.String
    },
    {
      key: "phone",
      name: "Teléfono",
      type: ColumnType.String
    },
  ]

  const [customerData, setCustomerData] = useState<Customer[]>([]);

  const getCustomers = useCallback(async () => {
    let data: Customer[] = await GetCustomersRequest({ companyId: 1, limit: 10, offset: 0 });
    const formatData: Customer[] = data.map(x => {
      return {
        ...x,
        identificationNumber: `${x.identificationType}-${x.identificationNumber}`
      }
    })
    setCustomerData(formatData);
  }, [])

  useEffect(() => {
    getCustomers()
  }, [getCustomers]);

  const deleteCustomerModal = useDeleteModal();

  return (
    <div>
      <SimpleCard>
        <DeleteModal onSubmit={() => { }} title="Eliminar Cliente" description="¿Estás seguro que quieres eliminar este cliente?" />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div>
            <label className="text-sm">Nombre</label>
            <Input size="sm" />
          </div>
          <div>
            <label className="text-sm">Apellido</label>
            <Input size="sm" />
          </div>
          <div>
            <label className="text-sm">Cédula</label>
            <Input size="sm" />
          </div>
          <div className="flex flex-col">
            <span className="opacity-0">.</span>
            <Button size="sm" variant="main">
              <Icon icon="tabler:search" />
            </Button>
          </div>
          <div className="xl:col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              <Link href="/management/customers/create">
                <Button size="sm" variant="main" className="w-full">
                  Crear Cliente
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable columns={customerCols} data={customerData} showDetails showEdit showDelete onDelete={deleteCustomerModal.onOpen} />
        </SimpleCard>
      </div>
    </div>
  )
}

export default CustomersClient;
