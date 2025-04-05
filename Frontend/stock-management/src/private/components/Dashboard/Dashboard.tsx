"use client";

import { MovDay } from "./MovDay";
import { SectionCards } from "./SectionCards";



const Dashboard = () => {
	return (
		<div className="flex flex-1 flex-col h-screen">
			<div className="@container/main flex flex-1 flex-row gap-2">
				{/* Sección de Cards */}
				<div className="flex flex-[3] flex-col gap-4 py-4 md:gap-6 md:py-6">
					<SectionCards />
				</div>

				{/* Línea divisoria + Movimientos */}
				<section className="flex flex-[1] p-4 border-l border-gray-300">
					<MovDay/>
				</section>
			</div>
		</div>

	);
};

export default Dashboard;
