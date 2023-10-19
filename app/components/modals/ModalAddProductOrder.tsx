'use client';

import useModalAddProductOrder from "@/app/hooks/useModalAddProductOrder";
import { Button, Modal } from "antd";
import { useState } from "react";

const ModalAddProductOrder = () => {
    const [loading, setLoading] = useState(false);
    const modalAddProductOrder = useModalAddProductOrder();

    const handleOk = () => {
        setLoading(true);
        setTimeout(() => {
        setLoading(false);
        modalAddProductOrder.onClose();
        }, 3000);
    };

    const handleCancel = () => {
        modalAddProductOrder.onClose();
    };

    return (
        <Modal
        open={modalAddProductOrder.isOpen}
        title="Title"
        onOk={handleOk}
        onCancel={handleCancel}
        footer={[
          <Button key="back" onClick={handleCancel}>
            Return
          </Button>,
          <Button key="submit" type="primary" loading={loading} onClick={handleOk}>
            Submit
          </Button>,
          <Button
            key="link"
            href="https://google.com"
            type="primary"
            loading={loading}
            onClick={handleOk}
          >
            Search on Google
          </Button>,
        ]}
      >
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
      </Modal>
    )
}

export default ModalAddProductOrder;