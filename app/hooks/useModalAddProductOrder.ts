import { create } from 'zustand';

interface ModalAddProductOrder {
    isOpen: boolean;
    onOpen: () => void;
    onClose: () => void;
}

const useModalAddProductOrder = create<ModalAddProductOrder>((set) => ({
    isOpen: false,
    onOpen: () => set({ isOpen: true}),
    onClose: () => set({ isOpen: false }),
}));

export default useModalAddProductOrder;