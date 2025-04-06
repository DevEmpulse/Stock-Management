"use client";

import { useFetch } from "@/private/hooks";
import { DataTable } from "../DataTable";
import { columns, Payment } from ".";


const url = 'http://127.0.0.1:3000/api/proveedor/3'


const Proveedor  = () => {
	const { data, error, loading } = useFetch<Payment[]>(url);
	
	  console.log(data);
	
	  if (loading) {
		return <h1>Cargando...</h1>; // Puedes mejorar esto con un componente de carga
	  }
	
	  if (error) {
		return <div>Error</div>;
	  }
	  if (!data) {
		return <div>No hay datos disponibles</div>;
	  }
	  // Aseguramos que data sea un array, usando un valor por defecto si es null o undefined
	  const safeData = data ?? [];
	return (
		 <div className="container mx-auto py-4">
			  <h1 className="font-bold text-xl ml-4">Lista de Proveedores</h1>
			  <DataTable
			  columns={columns}
			  data={safeData}
			  filters={[
				{ columnKey: "Empresa/Empresa", placeholder: "Filtrar por nombre/empresa" },
			  ]}
			  buttonText="Agregar Proveedor"
			  buttonClassName="bg-green-500 text-white cursor-pointer"
			/>
			</div>
	);
};

export default Proveedor;
