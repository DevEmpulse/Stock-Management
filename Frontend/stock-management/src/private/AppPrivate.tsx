import { BookOpenCheck, ChartLine, ClipboardList, DollarSign, Handshake, LayoutDashboard, PackageSearch, ShoppingCart, TrendingUp } from "lucide-react"
import { SidebarCategory, SidebarItem } from "./components"
import { Sidebar } from "./components/Sidebar"
import { Hedear } from "./components/Hedear"
import { useDispatch, useSelector } from "react-redux"
import { AppDispatch, RootState } from "@/store"
import Dashboard from "./components/Dashboard/Dashboard"

import { useEffect, useState } from "react"
import Productos from "./components/Productos/Productos"
import Proveedor from "./components/Proveedor/Proveedor"
import Categoria from "./components/Categoria/Categoria"
import { checkAuth } from "@/features/auth/authSlice"
import { Navigate } from "react-router-dom"



function AppPrivade() {
    const dispatch = useDispatch<AppDispatch>();
    const { user, isAuthenticated, loading } = useSelector((state: RootState) => state.auth);
    const [activePage, setActivePage] = useState("dashboard");
    
    useEffect(() => {
        dispatch(checkAuth());
    }, [dispatch]);

    if (loading) {
        return <div>Cargando...</div>;
    }

    // Redirigir a login si no está autenticado
    if (!isAuthenticated) {
        return <Navigate to="/login" />;
    }
   
    return (


        <div className="flex min-h-screen z-10">
            <Sidebar>
                <SidebarItem icon={<LayoutDashboard size={20} />} text="Dashboard" active={activePage === "dashboard"} onClick={() => setActivePage("dashboard")} />
                <SidebarCategory category="Movimientos" />
                <SidebarItem icon={<TrendingUp size={20} />} text="Ventas" />
                <SidebarItem icon={<ShoppingCart size={20} />} text="Compras" />
                <SidebarCategory category="Finanzas" />
                <SidebarItem icon={<ChartLine size={20} />} text="Ingresos" />
                <SidebarItem icon={<ClipboardList size={20} />} text="Gastos" />
                <SidebarItem icon={<DollarSign size={20} />} text="Balances" />
                <SidebarCategory category="Invetario" />
                <SidebarItem icon={<PackageSearch size={20} />} text="Productos" active={activePage === "productos"} onClick={() => setActivePage("productos")} />
                <SidebarItem icon={<Handshake size={20} />} text="Proveedores" active={activePage === "proveedores"} onClick={() => setActivePage("proveedores")} />
                <SidebarItem icon={<BookOpenCheck size={20} />} text="Categorias" active={activePage === "categorias"} onClick={() => setActivePage("categorias")} />
            </Sidebar>
            <div className="flex flex-col flex-1">
                <Hedear title="Dashboard" nombre={user?.nombre || 'Usuario'} email={user?.email || "name@gmail.com"} />
                <hr className="" />
                <div className="bg-white flex-1">
                    {activePage === "dashboard" && <Dashboard />}
                    {activePage === "productos" && <Productos />}
                    {activePage === "proveedores" && <Proveedor />}
                    {activePage === "categorias" && <Categoria />}
                </div>
            </div>
        </div>

    )
}

export default AppPrivade
