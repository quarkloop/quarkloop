import { ReactNode } from "react";
import { NavLink } from "@quarkloop/components";
import {
    IconAntenna,
    IconCirclesRelation,
    IconSquare,
    IconUser,
    IconUserCircle,
    IconWallet,
} from "@tabler/icons-react";

interface UserSettingsLayoutProps {
    children: ReactNode;
}

const UserSettingsLayout = (props: UserSettingsLayoutProps) => {
    const { children } = props;

    return (
        <div className="py-20 px-3 w-full self-center flex flex-col gap-10 md:w-2/3">
            <div className="px-7 pb-2 flex items-center text-2xl font-medium border-b">
                User Settings
            </div>
            <div className="flex flex-col gap-14 md:flex-row md:items-start">
                <div className="px-7 py-5 basis-[28%] flex flex-col gap-1 rounded-lg border">
                    <NavLink
                        href="/user/settings/profile"
                        label="Public profile"
                        icon={
                            <IconUser
                                size="1.5rem"
                                stroke="1.5"
                            />
                        }
                    />
                    <NavLink
                        href="/user/settings/account"
                        label="Account"
                        icon={
                            <IconUserCircle
                                size="1.5rem"
                                stroke="1.5"
                            />
                        }
                    />
                    <NavLink
                        href="/user/settings/linked-accounts"
                        label="Linked accounts"
                        icon={
                            <IconCirclesRelation
                                size="1.5rem"
                                stroke="1.5"
                            />
                        }
                    />
                    <NavLink
                        href="/user/settings/sessions"
                        label="Sessions"
                        icon={
                            <IconAntenna
                                size="1.5rem"
                                stroke="1.5"
                            />
                        }
                    />
                    <NavLink
                        href="/user/settings/plan"
                        label="Plan"
                        icon={
                            <IconSquare
                                size="1.5rem"
                                stroke="1.5"
                            />
                        }
                    />
                    {/* <NavLink
            href="/user/settings/billing"
            label="Billing"
            icon={
              <IconWallet
                size="1.5rem"
                stroke="1.5"
              />
            }
          /> */}
                </div>
                <div className="flex-1 flex items-start">{children}</div>
            </div>
        </div>
    );
};

export default UserSettingsLayout;
