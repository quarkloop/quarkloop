import { Sheet } from "../Sheet";

const SheetView = () => {
    return (
        <div className="flex-1 flex flex-col bg-red-100">
            <div className="text-2xl font-semibold">Sheets</div>
            <div className="flex-1 bg-blue-300">
                <Sheet />
            </div>
        </div>
    );
};

export { SheetView };
