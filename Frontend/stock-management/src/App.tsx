import { LayoutDashboard, LifeBuoy, PackageSearch, Settings, TrendingDown, TrendingUp, Truck} from "lucide-react"
import { SidebarItem } from "./private/components"
import { Sidebar } from "./private/components/Sidebar/Sidebar"


function App() {


  return (
    <div className="flex">
        <Sidebar>
          <SidebarItem icon={<LayoutDashboard size={30} />} text="Dashboard" active />
          <SidebarItem icon={<TrendingUp size={20} />} text="Ingresos"  />
          <SidebarItem icon={<TrendingDown size={20} />} text="Egresos"  />
          <SidebarItem icon={<PackageSearch size={20} />} text="Productos" />
          <SidebarItem icon={<Truck  size={20} />} text="Proveedores"  />
          <hr className="my-3" />
          <SidebarItem icon={<Settings size={20} />} text="Settings" />
          <SidebarItem icon={<LifeBuoy size={20} />} text="Help" />
          
        </Sidebar>
      </div>
  )
}

export default App
