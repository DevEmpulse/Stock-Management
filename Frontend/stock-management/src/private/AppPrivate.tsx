import { BookOpenCheck, ChartLine, ClipboardList, DollarSign, Handshake, LayoutDashboard, PackageSearch, ShoppingCart, TrendingUp } from "lucide-react"
import { SidebarCategory, SidebarItem } from "./components"
import { Sidebar } from "./components/Sidebar"
import { Hedear } from "./components/Hedear"
import { useSelector } from "react-redux"
import { RootState } from "@/store"
import Dashboard from "./components/Dashboard/Dashboard"

import { useState } from "react"
import Productos from "./components/Productos/Productos"



function AppPrivade() {
    const { user } = useSelector((state: RootState) => state.auth)
    console.log({ user })
    const [activePage, setActivePage] = useState("dashboard");
    return (


        <div className="flex min-h-screen z-10">
            <Sidebar>
            <SidebarItem icon={<LayoutDashboard size={30} />} text="Dashboard" active={activePage === "dashboard"} onClick={() => setActivePage("dashboard")} />
                <SidebarCategory category="Movimientos"/>
                <SidebarItem icon={<TrendingUp size={20} />} text="Ventas"/>
                <SidebarItem icon={<ShoppingCart size={20} />} text="Compras" />
                <SidebarCategory category="Finanzas"/>
                <SidebarItem icon={<ChartLine size={20}/>} text="Ingresos"/> 
                <SidebarItem icon={<ClipboardList size={20}/>} text="Gastos"/>
                <SidebarItem icon={<DollarSign size={20}/>} text="Balances"/>
                <SidebarCategory category="Invetario"/>
                <SidebarItem icon={<PackageSearch size={20} />} text="Productos" active={activePage ==="productos"} onClick={() => setActivePage("productos")} />
                <SidebarItem icon={<Handshake size={20} />} text="Proveedores" />
                <SidebarItem icon={<BookOpenCheck size={20} />} text="Categorias" />
            </Sidebar>
            <div className="flex flex-col flex-1">
                <Hedear title="Dashboard" nombre={user?.nombre || 'Usuario'} email={user?.email || "name@gmail.com"} />
                <hr className="" />
                <div className="bg-white flex-1">
                    {activePage === "dashboard" && <Dashboard />}
                    {activePage === "productos" && <Productos />}
                </div>
            </div>
        </div>

    )
}

export default AppPrivade
