import { Providers } from "@quarkloop/components";

import "@/styles/globals.css";

interface RootAuthLayoutProps {
    children: React.ReactElement;
}

const RootAuthLayout = async ({ children }: RootAuthLayoutProps) => {
    return (
        <html lang="en">
            <body className="h-auto w-full flex flex-col">
                <Providers>
                    <main className="flex-1 flex">{children}</main>
                </Providers>
            </body>
        </html>
    );
};

export default RootAuthLayout;
