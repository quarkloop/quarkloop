import { Metadata } from "next";
import { Layout, Providers } from "@quarkloop/components";

import "@/styles/globals.css";

export const metadata: Metadata = {
    title: {
        default: "Quarkloop",
        template: "%s | Quarkloop",
    },
    description: "Quarkloop - redefining industry",
    applicationName: "Quarkloop",
    keywords: ["Quarkloop", "technology", "platform"],
    creator: "Quarkloop",
    publisher: "Quarkloop",
    metadataBase: new URL("https://quarkloop.com/org"),
    // robots: {
    //     index: true,
    //     follow: true,
    //     nocache: true,
    //     googleBot: {
    //         index: true,
    //         follow: true,
    //         noimageindex: true,
    //         'max-video-preview': -1,
    //         'max-image-preview': 'large',
    //         'max-snippet': -1,
    //     },
    // },
    // verification: {
    //     google: 'google',
    //     yandex: 'yandex',
    //     yahoo: 'yahoo',
    // },
};

// type Props = {
//     params: { orgId: string };
//     //searchParams: { [key: string]: string | string[] | undefined };
// };

// export async function generateMetadata({ params }: Props): Promise<Metadata> {
//     const org = await findUniqueOrganization({ orgId: params.orgId });

//     return {
//         title: {
//             absolute: org ? org.name : "",
//         },
//     };
// }

interface RootLayoutProps {
    children: React.ReactNode;
    params: any;
}

const RootLayout = (props: RootLayoutProps) => {
    const {
        children,
        params: { orgId, workspaceId, projectId },
    } = props;

    return (
        <html lang="en">
            <body className="h-auto w-full">
                <Providers>
                    <Layout params={props.params}>{props.children}</Layout>
                </Providers>
            </body>
        </html>
    );
};

export default RootLayout;
