import { cn } from "@/lib/utils";

const MovDay = () => {
	const movimientos = [
		{ producto: "Laptop Dell XPS", tipo: "Venta", precio: "$1200.00" },
		{ producto: "Monitor Samsung", tipo: "Compra", precio: "$300.00" },
		{ producto: "Teclado Mecánico", tipo: "Venta", precio: "$80.00" },
		{ producto: "Mouse Logitech", tipo: "Egreso", precio: "$25.00" },
	
	];

	return (
		<div className="max-w-4xl mx-auto p-6">
			<h1 className="text-xl font-bold mb-4 text-black">Movimientos del día</h1>
			<div className="overflow-x-auto">
				<table className="min-w-full bg-white shadow-md ">
					<thead className="bg-blue-gray-200">
						<tr>
							<th className="py-2 px-4 text-left text-xs font-semibold text-black">Productos</th>
							<th className="py-2 px-4 text-left text-xs font-semibold text-black">Tipo de Movimiento</th>
							<th className="py-2 px-4 text-left text-xs font-semibold text-black">Precio</th>
						</tr>
					</thead>
					<tbody>
						{movimientos.map((mov, index) => (
							<tr key={index} className="border-b border-blue-gray-100 hover:bg-blue-gray-50">
								<td className="py-3 px-4 text-sm text-black">{mov.producto}</td>
								<td className="py-3 px-4">
									<span
										className={cn(
											"inline-block px-2 py-1 text-xs font-medium text-black rounded",
											mov.tipo === "Venta" ? "bg-teal-500" : "bg-red-600"
										)}
									>
										{mov.tipo}
									</span>
								</td>
								<td className="py-3 px-4 text-sm text-black">{mov.precio}</td>
							</tr>
						))}
					</tbody>
				</table>
			</div>
		</div>
	);
};

export default MovDay;
