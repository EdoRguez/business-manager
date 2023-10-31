"use client";

import useModalAddProductOrder from "@/app/hooks/useModalAddProductOrder";
import { Button, Input, Modal } from "antd";
import React, { useState } from "react";

interface modalProps {
  title: string;
}

interface formData {
  productName: string | undefined;
  cost: number | undefined;
  quantity: number | undefined;
  sell: number | undefined;
}

const ModalAddProductOrder: React.FC<modalProps> = ({ title }) => {
  const [loading, setLoading] = useState(false);
  const modalAddProductOrder = useModalAddProductOrder();
  const [errors, setErrors] = useState<any>({});
  const [formData, setFormData] = useState<formData>({
    productName: undefined,
    cost: undefined,
    quantity: undefined,
    sell: undefined,
  });

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleSubmit = () => {
    setLoading(true);

    setErrors(validateValues());

    if(!errors) {
      modalAddProductOrder.onClose();
    } else {
      setLoading(false);
    }

  };

  const handleCancel = () => {
    modalAddProductOrder.onClose();
  };

  const validateValues = () => {
    let errors: any = {};

    if (!formData.productName || formData.productName.length > 15) {
      errors.productName = "Producto es requerido";
    }

    if (!formData.cost || formData.cost <= 0 || formData.cost >= 999) {
      errors.cost = "Costo es requerido y entre 1$ - 999$";
    }

    if (!formData.quantity || formData.quantity <= 0 || formData.quantity >= 999) {
      errors.quantity = "Cantidad es requerida y entre 1 - 999";
    }

    if (!formData.sell || formData.sell <= 0 || formData.sell >= 999) {
      errors.sell = "Venta es requerido y entre 1$ - 999$";
    }

    return errors;
  }

  return (
    <Modal
      open={modalAddProductOrder.isOpen}
      title={title}
      onOk={handleSubmit}
      onCancel={handleCancel}
      footer={[
        <Button
          key="submit"
          type="primary"
          loading={loading}
          onClick={handleSubmit}
          className="
            rounded-lg
            hover:opacity-80
            hover:bg-green-500
            transition
            w-full
            bg-green-500
            border-green-500
            text-white
            text-sm
            small
            font-light
            border-[1px]
          "
        >
          Search on Google
        </Button>,
      ]}
    >
      <form className="my-8">
        <div>
          <label>Nombre Producto:</label>
          <Input
            status={errors.productName ? 'error' : ''}
            placeholder="Basic usage"
            name="productName"
            maxLength={15}
            value={formData.productName}
            onChange={handleChange}
          />
        </div>

        <div className="flex mt-3">
          <div className="mr-1">
            <label>Costo c/u:</label>
            <Input
            status={errors.cost ? 'error' : ''}
              placeholder="Basic usage"
              name="cost"
              maxLength={3}
              value={formData.cost}
              onChange={handleChange}
            />
          </div>
          <div className="ml-1">
            <label>Cantidad:</label>
            <Input
              status={errors.quantity ? 'error' : ''}
              placeholder="Basic usage"
              name="quantity"
              maxLength={3}
              value={formData.quantity}
              onChange={handleChange}
            />
          </div>
        </div>

        <div className="mt-3">
          <label>Venta Individual:</label>
          <Input
          status={errors.sell ? 'error' : ''}
            placeholder="Basic usage"
            name="sell"
            value={formData.sell}
            onChange={handleChange}
          />
        </div>
      </form>
    </Modal>
  );
};

export default ModalAddProductOrder;
