"use client"

import { Button } from "@/components/ui/button"
import { ColumnDef } from "@tanstack/react-table"

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type Payment = {
	id_prod: number
	cuit_cuil: string
	nombre_apellido: string
	email: string
	num_contacto: string
	comentario: string
}

export const columns: ColumnDef<Payment>[] = [
  {
	accessorKey: "nombre_apellido",
	header: "Nombre y Apellido / Empresa",
  },
  {
	accessorKey:"email",
	header: "Email"
  },
  {
	accessorKey: "num_contacto",
	header: "Número de contacto",
  },
  {
	accessorKey: "comentario",
	header: "Comentario",
  },

  {
	id: "actions", // No usa accessorKey porque no mapea un campo directo
	header: "Acciones",
	cell: ({ row }) => {
	  const payment = row.original

	  return (
		<div className="flex space-x-2">
		  <Button
		 
			variant="outline"
			size="sm"
			 className="dark:bg-blue-500 bg-blue-500 text-white dark:hover:bg-blue-600 hover:bg-blue-600  cursor-pointer"
			onClick={() => console.log("Modificar", payment.id_prod)} // Aquí va tu lógica de modificar
		  >
			Modificar
		  </Button>
		  <Button
			className="cursor-pointer"
			variant="destructive"
			size="sm"
			onClick={() => console.log("Borrar", payment.id_prod)} // Aquí va tu lógica de borrar
		  >
			Borrar
		  </Button>
		</div>
	  )
	},
  },
]
