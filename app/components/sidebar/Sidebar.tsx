"use client";

interface SidebarProps {
  children: React.ReactNode;
}

const Siderbar: React.FC<SidebarProps> = ({ children }) => {
  const Menus = [
    { title: "Dashboard", src: "Chart_fill" },
    { title: "Inbox", src: "Chat" },
    { title: "Accounts", src: "User", gap: true },
    { title: "Schedule ", src: "Calendar" },
    { title: "Search", src: "Search" },
    { title: "Analytics", src: "Chart" },
    { title: "Files ", src: "Folder", gap: true },
    { title: "Setting", src: "Setting" },
  ];

  return (
    <div className="flex">
      <div
        className="
          w-72 
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
          <img
            src="/images/logo.png"
            className={"cursor-pointer duration-500"}
          />
          <h1
            className="
              text-white 
              origin-left 
              font-medium 
              text-xl 
              duration-200 
              scale-0"
          >
            Designer
          </h1>
        </div>
        <ul className="pt-6">
          {Menus.map((Menu, index) => (
            <li
              key={index}
              className={`flex rounded-md p-2 cursor-pointer hover:bg-light-white text-gray-300 text-sm items-center gap-x-4 ${
                Menu.gap ? "mt-9" : "mt-2"
              } ${index === 0 && "bg-light-white"} `}
            >
              <img src={`/images/${Menu.src}.png`} />
              <span className="origin-left duration-200">{Menu.title}</span>
            </li>
          ))}
        </ul>
      </div>
      <div className="flex-1 p-7 bg-zinc-50">
        <h1 className="text-2xl font-semibold ">Home Page</h1>
        <div>{children}</div>
      </div>
    </div>
  );
};

export default Siderbar;
