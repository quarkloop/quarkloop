import Image from "next/image";
import Link from "next/link";

const QuarkloopLogo = () => (
    <Link
        href="/"
        className="w-10 h-10 relative rounded-full">
        <Image
            fill
            priority
            unoptimized
            src="/logo.jpeg"
            alt="Quarkloop"
            sizes="(max-width: 40px, max-height: 40px)"
            className="object-cover rounded-full"
        />
    </Link>
);

export { QuarkloopLogo };
