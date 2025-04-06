"use client";

import { useFetch } from "@/private/hooks";
import { DataTable } from "../DataTable";
import { columns } from "./";

const url = "http://127.0.0.1:3000/api/producto/3";

interface DataProductos {
  id_prod: number;
  nombre: string;
  codigo_de_barras: string;
  imagen?: string;
  stock: number;
  descripcion: string;
  precio_compra: number;
  precio_venta: number;
  fecha_creacion: string;
  fecha_actualizacion: string;
}

const Productos = () => {
  const { data, error, loading } = useFetch<DataProductos[]>(url);

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
    <div className="container mx-auto py-10">
      <DataTable columns={columns} data={safeData} />
    </div>
  );
};

export default Productos;