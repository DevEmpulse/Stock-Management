
import { logout } from "@/features/auth/authSlice"
import { AppDispatch} from "@/store"
import { Avatar, Dropdown, DropdownHeader, DropdownItem, Navbar, NavbarBrand } from "flowbite-react"
import { useDispatch } from "react-redux"



interface Props {
    title: string
    nombre: string
    email: string
}

export default function Hedear({ title, nombre, email }: Props) {
    
    const dispatch = useDispatch<AppDispatch>();

    return (

        <Navbar fluid className="bg-blue-800 dark:bg-blue-800">
            <NavbarBrand>
                <span className="self-center whitespace-nowrap text-xl font-bold dark:text-white text-white">{title}</span>
            </NavbarBrand>
            <div className="flex md-order-2 mr-6">
                <h2 className="self-center whitespace-normal text-xl font-medium mr-2.5 text-white">Hola, {nombre}</h2>
                <Dropdown
                    arrowIcon={false}
                    inline
                    label={
                        <Avatar 
                        alt="User" 
                        img="https://flowbite.com/docs/images/people/profile-picture-5.jpg" 
                        rounded 
                        className="cursor-pointer hover:scale-105" />
                    }
                >
                    <DropdownHeader>
                        <span className="block truncate text-sm font-medium">{email}</span>
                    </DropdownHeader>
                    <DropdownItem onClick={()=>dispatch(logout())}>Cerrar Sesión</DropdownItem>
                </Dropdown>

            </div>
        </Navbar>


    )
}