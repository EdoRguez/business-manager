"use client";

import { MenuElement } from "@/app/common/types";
import Link from "next/link";
import { AiOutlineDollar, AiOutlineHome } from "react-icons/ai";
import { BsBoxSeamFill } from "react-icons/bs";
import { BiSolidPlaneAlt} from "react-icons/bi";
import { FcShop } from "react-icons/fc";
import { usePathname } from 'next/navigation';

interface SidebarProps {
  children: React.ReactNode;
}

const Siderbar: React.FC<SidebarProps> = ({ children }) => {
  const pathname = usePathname();

  const menuElements: MenuElement[] = [
    { title: "Inicio", path: '/', icon: AiOutlineHome, gap: true },
    { title: "Pedidos", path: '/orders/1', icon: BiSolidPlaneAlt },
    { title: "Ventas", path: '', icon: AiOutlineDollar },
    { title: "Productos", path: '', icon: BsBoxSeamFill },
  ];

  const renderIcon = (icon: any, index: number) => {
    const Icon = icon;
    return (
      <div key={index}>
        <Icon />
      </div>
    );
  };

  return (
    <div className="flex">
      <div
        className="
          w-55 
          bg-dark-purple 
          h-screen 
          p-5  
          pt-8 
          relative 
          duration-300 
          hidden 
          lg:block"
      >
        <div
          className="
            flex 
            gap-x-4 
            items-center"
        >
          <div className="text-2xl">
          <FcShop />
          </div>
          <h1
            className="
              origin-left 
              font-bold
              duration-200 
              text-green-500
              text-xl
              select-none"
          >
            Business Manager
          </h1>
        </div>
        <ul className="pt-6">
          { menuElements.map((element, index) => (
            <Link 
              href={element.path} 
              key={index}
            >
              <li
                className={`
                  flex 
                  items-center 
                  rounded-md  
                  p-2 
                  cursor-pointer 
                  gap-x-4 
                  select-none 
                  border-2 
                  hover:bg-zinc-200 
                  hover:text-green-500 
                  ${ element.gap ? "mt-9" : "mt-2" } 
                  ${ index === 0 && "bg-white" } 
                  ${ pathname == element.path ? "bg-zinc-200 text-green-500 border-green-500" : "border-white"}
                  `}
              >
                <div className="text-xl duration-200">
                  { renderIcon(element.icon, index) }
                </div>
                <span className="origin-left duration-200 text-xl">{ element.title }</span>
              </li>
            </Link>
          ))}
        </ul>
      </div>
      <div className="flex-1 p-7 bg-zinc-50">
        <h1 className="text-2xl font-semibold ">Home Page</h1>
        <div>{ children }</div>
      </div>
    </div>
  );
};

export default Siderbar;
