'use client';

import SimpleCard from "../components/cards/SimpleCard";
import PaymentCard from "../components/cards/PaymentCard";
import PaymentFilterCard from "../components/cards/PaymentFilterCard";
import BreadcrumbNavigation from "../components/BreadcrumbNavigation";
import { BreadcrumItem } from "../types";
import { Button } from "@chakra-ui/react";
import Link from "next/link";

const PaymentsClient = () => {

  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/payments"
    }
  ];

  const payments: any[] = [...Array(11)];

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 lg:grid-cols-5 gap-4">
          <div>
            <h1 className='ml-2 font-bold'>Métodos de Pago</h1>
          </div>
          <div className="lg:col-start-5">
            <Link href="/payments/create">
              <Button size="sm" variant="main" className="w-full">
                Crear Método de Pago
              </Button>
            </Link>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <div className="grid grid-cols-1 xl:grid-cols-2 gap-3">
          <SimpleCard>
            <div className="p-2">
              <h3>Filtrar</h3>
              <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 mt-1">
                {
                  payments.map((val: any, index: number) => (
                    <PaymentFilterCard key={index} paymentTypeEnum={index} description="Description" isSelected={index == 0} />
                  ))
                }
              </div>
            </div>
          </SimpleCard>
          <SimpleCard>
            <div className="p-2">
              <h3>Métodos de Pago</h3>
              <div className="mt-1">
                <PaymentCard name="Transferencia" paymentTypeEnum={0} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Pago Movil" paymentTypeEnum={1} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Binance" paymentTypeEnum={2} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Colombia" paymentTypeEnum={3} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Panama" paymentTypeEnum={4} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Paypal" paymentTypeEnum={5} />
              </div>
              <div className="mt-1">
                <PaymentCard name="USA" paymentTypeEnum={6} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Zelle" paymentTypeEnum={7} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Zinli" paymentTypeEnum={8} />
              </div>
              <div className="mt-1">
                <PaymentCard name="Otro" paymentTypeEnum={9} />
              </div>
            </div>
          </SimpleCard>
        </div>
      </div>
    </div>
  )

}

export default PaymentsClient;
