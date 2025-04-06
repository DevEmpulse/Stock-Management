"use client"

import { ColumnDef } from "@tanstack/react-table"

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type Payment = {
	id_prod: number
	nombre: string
	codigo_de_barras: string
	imagen?: string
	stock: number
	descripcion: string
	precio_compra: number
	precio_venta: number
	fecha_creacion: string
	fecha_actualizacion: string
}

export const columns: ColumnDef<Payment>[] = [
  {
    accessorKey: "nombre",
    header: "Nombre",
  },
  {
    accessorKey: "stock",
    header: "Stock",
  },
  {
    accessorKey: "descripcion",
    header: "Descripcion",
  },
  {
    accessorKey: "precio_compra",
    header: "Precio de compra",
  },
  {
    accessorKey: "precio_venta",
    header: "Precio de venta",
  },
]
