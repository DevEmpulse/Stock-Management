import { Badge } from "@/components/ui/badge";
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { TrendingDownIcon, TrendingUpIcon } from "lucide-react";


const SectionCards = () => {
	return (
		<div>
			<div className=" *:data-[slot=card]:shadow-xs @xl/main:grid-cols-2 @5xl/main:grid-cols-2 grid grid-cols-1 gap-4 px-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card dark:*:data-[slot=card]:bg-card lg:px-6">
				<Card className="bg-blue-400 @container/card">
					<CardHeader className="relative">
						<CardDescription>Ventas</CardDescription>
						<CardTitle className="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums">
							$1,250.00
						</CardTitle>
						<div className="absolute right-4 top-4">
							<Badge variant="outline" className="flex gap-1 rounded-lg text-xs">
								<TrendingUpIcon className="size-3" />
								+12.5%
							</Badge>
						</div>
					</CardHeader>
				</Card>
				<Card className=" bg-blue-400 @container/card">
					<CardHeader className="relative">
						<CardDescription>Compras</CardDescription>
						<CardTitle className="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums">
							$1,234.00
						</CardTitle>
						<div className="absolute right-4 top-4">
							<Badge variant="outline" className="flex gap-1 rounded-lg text-xs">
								<TrendingDownIcon className="size-3" />
								-20%
							</Badge>
						</div>
					</CardHeader>
				</Card>
			</div>
		</div>
	);
};

export default SectionCards;
