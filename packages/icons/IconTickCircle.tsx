
interface IconTickCircleProps {
    size: string;
}

const IconTickCircle = (props: IconTickCircleProps) => {
    const { size } = props;
    return (<div className={`h-[${size}] w-[${size}] flex shrink-0 grow-0 items-center justify-center rounded-full bg-green-300 text-green-700`}>
        <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            strokeWidth="2">
            <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" />
        </svg>
    </div>);
}

export { IconTickCircle };