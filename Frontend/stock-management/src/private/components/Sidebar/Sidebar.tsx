import { ChevronFirst, ChevronLast } from "lucide-react";
import logo from "/src/assets/logo.png";
import { createContext, useContext, useState, ReactNode } from "react";
import { Link } from "react-router-dom";

interface SidebarContextType {
  expanded: boolean;
}

const SidebarContext = createContext<SidebarContextType | undefined>(undefined);

interface SidebarProps {
  children?: ReactNode;
}

export default function Sidebar({ children}: SidebarProps) {
  const [expanded, setExpanded] = useState(true);

  return (
    <aside className="h-auto">
      <nav className="h-full flex flex-col bg-white dark:bg-white border-r shadow-sm">
        <div className="p-4 pb-2 flex justify-between items-center">
          <img
            src={logo}
            className={`overflow-hidden transition-all ${expanded ? "w-12" : "w-0"}`}
            alt="Logo"
          />
          <span className={`text-black font-bold overflow-hidden transition-all ${expanded ? "w-auto ml-2" : "w-0"}`}>
            Sistema de Gestión
          </span>
          <button
            onClick={() => setExpanded((curr) => !curr)}
            className="p-1.5 rounded-lg bg-gray-100 hover:bg-gray-100"
          >
            {expanded ? <ChevronFirst className="cursor-pointer" /> : <ChevronLast className="cursor-pointer" />}
          </button>
        </div>
    
        <SidebarContext.Provider value={{ expanded }}>
          <ul className="flex-1 px-3">{children}</ul>
        </SidebarContext.Provider>
      </nav>
    </aside>
  );
}

interface SidebarItemProps {
  icon: ReactNode;
  text: string;
  to?: string
  active?: boolean;
  alert?: boolean;
  onClick?: () => void
}

export function SidebarItem({ icon, text, active, alert, to, onClick}: SidebarItemProps) {
  const context = useContext(SidebarContext);
  if (!context) {
    throw new Error("SidebarItem must be used within a Sidebar");
  }
  const { expanded } = context;

  const content = (
   
    <li
      className={`relative flex items-center py-2 px-3 my-1 font-medium rounded-md cursor-pointer transition-colors group ${active ? "bg-gradient-to-tr from-indigo-600 to-indigo-400 text-black" : "hover:bg-indigo-500 text-black"
        }`} onClick={onClick}
    >
      {icon}
      <span className={`overflow-hidden transition-all ${expanded ? "w-52 ml-3" : "w-0"}`}>{text}</span>
      {alert && <div className={`absolute right-2 w-2 h-2 rounded bg-white ${expanded ? "" : "top-2"}`} />}
      {!expanded && (
        <div
          className="absolute left-full rounded-md z-1 px-2 py-1 ml-6 bg-indigo-100 text-indigo-800 text-sm invisible opacity-20 -translate-x-3 transition-all group-hover:visible group-hover:opacity-100 group-hover:translate-x-0"
        >
          {text}
        </div>
      )}
    </li>

  )

  return to ? <Link to={to}>{content}</Link> : content;
} 

interface Props{
  category: string
}


export function SidebarCategory({category}:Props){
  const context = useContext(SidebarContext);
  if (!context) {
    throw new Error("SidebarItem must be used within a Sidebar");
  }
  const { expanded } = context;

  return (
    <div className="mb-2 mt-4">
    {expanded && <h4 className="text-black text-sm px-2">{category}</h4>}
    <hr className="border-white/20 mt-1 mb-2" />
  </div>
  )
}