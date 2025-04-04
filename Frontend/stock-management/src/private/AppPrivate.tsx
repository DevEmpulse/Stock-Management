import { LayoutDashboard, LifeBuoy, PackageSearch, Settings, TrendingDown, TrendingUp, Truck } from "lucide-react"
import { SidebarItem } from "./components"
import { Sidebar } from "./components/Sidebar"
import { Hedear } from "./components/Hedear"
import { useSelector } from "react-redux"
import { RootState } from "@/store"
import Dashboard from "./components/Dashboard/Dashboard"




function AppPrivade() {
        const {user}= useSelector((state:RootState)=>state.auth)
        console.log({user})
    return (
        
            
            <div className="flex z-10">
            <Sidebar>
                <SidebarItem icon={<LayoutDashboard size={30} />} text="Dashboard" active />
                <SidebarItem icon={<TrendingUp size={20} />} text="Ingresos" />
                <SidebarItem icon={<TrendingDown size={20} />} text="Egresos" />
                <SidebarItem icon={<PackageSearch size={20} />} text="Productos" />
                <SidebarItem icon={<Truck size={20} />} text="Proveedores" />
                <hr className="my-3" />
                <SidebarItem icon={<Settings size={20} />} text="Settings" />
                <SidebarItem icon={<LifeBuoy size={20} />} text="Help" />
            </Sidebar>
            <div className="flex flex-col flex-1">
                <Hedear title="Dashboard" nombre={user?.nombre || 'Usuario'} email={user?.email||"name@gmail.com"}/>
                <hr className=""/>
                <div className="bg-blue-800 flex-1">
                    {/* Aquí irá el contenido de cada página */}
                    <Dashboard/>
                </div>
            </div>
            </div>

    )
}

export default AppPrivade
