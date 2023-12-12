
// https://www.svgrepo.com/svg/448951/caret-down-filled

interface IconCaretDownFilledProps {
    size: string;
}

const IconCaretDownFilled = (props: IconCaretDownFilledProps) => {
    const { size } = props;
    return (<div className={`h-[${size}] w-[${size}] flex-no-shrink fill-current`}>
        <svg
            fill="#000000"
            height={size}
            width={size}
            version="1.1"
            id="Layer_1"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24">
            <polygon className="fill-rule:evenodd;clip-rule:evenodd;" points="3,6 21,6 12,19 " />
        </svg>
    </div>);
}

export { IconCaretDownFilled };