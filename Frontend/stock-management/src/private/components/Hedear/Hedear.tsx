


interface Props {
    title: string
}

export default function Hedear({title}:Props) {
    return (
        <nav className="bg-blue-800 px-4 py-3 flex justify-center">
            <div className="text-white me-4 cursor-pointer">
                <span className="font-bold text-white">{title}</span>
            </div>

        </nav>
    )
}