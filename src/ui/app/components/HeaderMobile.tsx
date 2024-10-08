"use client";

import React, { ReactNode, useEffect, useRef, useState } from "react";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";

import { SIDENAV_ITEMS } from "../constants";
import { SideNavItem } from "../types";
import { Icon } from "@iconify/react";
import { motion, useCycle } from "framer-motion";
import deleteUserSession from "../actions/deleteUserSession";
import { CurrentUser } from "../types/auth";
import getCurrentUser from "../actions/getCurrentUser";
import { validateUserInRoles } from "../utils/Utils";

type MenuItemWithSubMenuProps = {
  item: SideNavItem;
  toggleOpen: () => void;
};

const sidebar = {
  open: (height = 1000) => ({
    clipPath: `circle(${height * 2 + 200}px at 100% 0)`,
    transition: {
      type: "spring",
      stiffness: 20,
      restDelta: 2,
    },
  }),
  closed: {
    clipPath: "circle(0px at 100% 0)",
    transition: {
      type: "spring",
      stiffness: 400,
      damping: 40,
    },
  },
};

const HeaderMobile = () => {
  const { push } = useRouter();
  const pathname = usePathname();
  const containerRef = useRef(null);
  const { height } = useDimensions(containerRef);
  const [isOpen, toggleOpen] = useCycle(false, true);
  const [currentUser, setCurrentUser] = useState<CurrentUser>();

  const onCloseSession = () => {
    toggleOpen();
    deleteUserSession();
    push("/login");
  };

  const onAccount = () => {
    toggleOpen();
    push("/management/account");
  };

  useEffect(() => {
    const user: CurrentUser | null = getCurrentUser();
    if (user) {
      setCurrentUser(user);
    }
  }, []);

  return (
    <motion.nav
      initial={false}
      animate={isOpen ? "open" : "closed"}
      custom={height}
      className={`fixed inset-0 z-50 w-full md:hidden ${
        isOpen ? "" : "pointer-events-none"
      }`}
      ref={containerRef}
    >
      <motion.div
        className="absolute inset-0 right-0 w-full bg-white"
        variants={sidebar}
      />
      <motion.ul
        variants={variants}
        className="absolute grid w-full gap-3 px-10 py-16"
      >
        {SIDENAV_ITEMS.map((item, idx) => {
          if (
            currentUser &&
            validateUserInRoles(item.roleIds, currentUser.roleId)
          ) {
            return (
              <div key={idx}>
                {item.submenu ? (
                  <MenuItemWithSubMenu item={item} toggleOpen={toggleOpen} />
                ) : (
                  <MenuItem>
                    <Link
                      href={item.path}
                      onClick={() => toggleOpen()}
                      className={`flex items-center w-full text-lg text-black ${
                        item.path === pathname ? "font-bold" : ""
                      }`}
                    >
                      {item.icon}
                      <span className="ml-2">{item.title}</span>
                    </Link>
                  </MenuItem>
                )}

                <MenuItem className="my-3 h-px w-full bg-gray-300" />
              </div>
            );
          } else {
            return null;
          }
        })}
        <MenuItem>
          <button
            onClick={onAccount}
            className={`flex items-center w-full text-lg text-black`}
          >
            <Icon icon="mdi:user" />
            <span className="ml-2">Cuenta</span>
          </button>
        </MenuItem>
        <MenuItem className="my-3 h-px w-full bg-gray-300" />
        <MenuItem>
          <button
            onClick={onCloseSession}
            className={`flex items-center w-full text-lg text-black`}
          >
            <Icon icon="ci:exit" />
            <span className="ml-2">Cerrar Sesión</span>
          </button>
        </MenuItem>
      </motion.ul>
      <MenuToggle toggle={toggleOpen} />
    </motion.nav>
  );
};

export default HeaderMobile;

const MenuToggle = ({ toggle }: { toggle: any }) => (
  <button
    onClick={toggle}
    className="pointer-events-auto absolute right-4 top-[14px] z-30 pl-[3.5px]"
  >
    <svg width="23" height="23" viewBox="0 0 23 23">
      <Path
        variants={{
          closed: { d: "M 2 2.5 L 20 2.5" },
          open: { d: "M 3 16.5 L 17 2.5" },
        }}
      />
      <Path
        d="M 2 9.423 L 20 9.423"
        variants={{
          closed: { opacity: 1 },
          open: { opacity: 0 },
        }}
        transition={{ duration: 0.1 }}
      />
      <Path
        variants={{
          closed: { d: "M 2 16.346 L 20 16.346" },
          open: { d: "M 3 2.5 L 17 16.346" },
        }}
      />
    </svg>
  </button>
);

const Path = (props: any) => (
  <motion.path
    fill="transparent"
    strokeWidth="2"
    stroke="hsl(0, 0%, 18%)"
    strokeLinecap="round"
    {...props}
  />
);

const MenuItem = ({
  className,
  children,
}: {
  className?: string;
  children?: ReactNode;
}) => {
  return (
    <motion.li variants={MenuItemVariants} className={className}>
      {children}
    </motion.li>
  );
};

const MenuItemWithSubMenu: React.FC<MenuItemWithSubMenuProps> = ({
  item,
  toggleOpen,
}) => {
  const pathname = usePathname();
  const [subMenuOpen, setSubMenuOpen] = useState(false);

  return (
    <>
      <MenuItem>
        <button
          className="flex w-full text-lg"
          onClick={() => setSubMenuOpen(!subMenuOpen)}
        >
          <div className="flex flex-row justify-between w-full items-center">
            <span
              className={`flex flex-row items-center text-black ${
                pathname.includes(item.path) ? "font-bold" : ""
              }`}
            >
              {item.icon}
              <span className="ml-2">{item.title}</span>
            </span>
            <div className={`text-black ${subMenuOpen && "rotate-180"}`}>
              <Icon icon="lucide:chevron-down" width="24" height="24" />
            </div>
          </div>
        </button>
      </MenuItem>
      <div className="mt-2 ml-2 flex flex-col space-y-2 text-black">
        {subMenuOpen && (
          <>
            {item.subMenuItems?.map((subItem, subIdx) => {
              return (
                <MenuItem key={subIdx}>
                  <Link
                    href={subItem.path}
                    onClick={() => toggleOpen()}
                    className={` ${
                      subItem.path === pathname ? "font-bold" : ""
                    }`}
                  >
                    {subItem.title}
                  </Link>
                </MenuItem>
              );
            })}
          </>
        )}
      </div>
    </>
  );
};

const MenuItemVariants = {
  open: {
    y: 0,
    opacity: 1,
    transition: {
      y: { stiffness: 1000, velocity: -100 },
    },
  },
  closed: {
    y: 50,
    opacity: 0,
    transition: {
      y: { stiffness: 1000 },
      duration: 0.02,
    },
  },
};

const variants = {
  open: {
    transition: { staggerChildren: 0.02, delayChildren: 0.15 },
  },
  closed: {
    transition: { staggerChildren: 0.01, staggerDirection: -1 },
  },
};

const useDimensions = (ref: any) => {
  const dimensions = useRef({ width: 0, height: 0 });

  useEffect(() => {
    if (ref.current) {
      dimensions.current.width = ref.current.offsetWidth;
      dimensions.current.height = ref.current.offsetHeight;
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ref]);

  return dimensions.current;
};
