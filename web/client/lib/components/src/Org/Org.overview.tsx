"use client";

import { useMemo } from "react";
import { useParams } from "next/navigation";
import Link from "next/link";

import { useGetOrgByIdQuery } from "./Org.net.client";
import { useOrgData } from "./Org.util";

const OrgDashboard = () => {
    const { orgSid }: { orgSid: string } = useParams();
    const { data: orgData } = useGetOrgByIdQuery({ orgSid });

    const rowData = useOrgData(orgData?.data);
    const renderData = useMemo(() => {
        const d = orgData?.data;
        const org = [
            {
                label: "Name",
                value: d?.name,
            },
            {
                label: "Scope ID",
                value: d?.sid,
            },
            {
                label: "Description",
                value: d?.description,
            },
            {
                label: "Visibility",
                value: d?.visibility,
            },
            {
                label: "Created By",
                value: d?.createdBy,
            },
            {
                label: "Updated By",
                value: d?.updatedBy,
            },
        ];
        return org;
    }, [orgData]);

    return (
        <div className="p-10 flex-1 flex flex-col gap-3">
            <div className="basis-[300px] p-5 flex flex-col gap-3 rounded-lg border">
                {renderData.map((d, idx) => (
                    <div
                        key={idx}
                        className="flex flex-col">
                        <p className="">{d.label}</p>
                        <p className="truncate font-medium">{d.value ?? ""}</p>
                    </div>
                ))}
            </div>
            <div className="basis-[300px] p-5 flex flex-col gap-3 rounded-lg border">
                <div className="flex flex-col">
                    <div className="flex items-center gap-4">
                        <Link
                            href={`/manage/${rowData.sid}`}
                            className="flex items-center text-base font-semibold hover:text-neutral-700">
                            {rowData.name}
                        </Link>
                        <div className="px-2 flex items-center bg-neutral-200 text-neutral-500 text-sm rounded-lg">
                            {rowData.visibility}
                        </div>
                        <Link
                            href={rowData.path ?? ""}
                            rel="noopener noreferrer"
                            target="_blank"
                            className="flex items-center underline hover:text-neutral-700">
                            view profile
                        </Link>
                    </div>
                    <div className="w-[300px] md:w-[600px] inline-block truncate">
                        {rowData.description}
                    </div>
                    <div className="flex items-center gap-1">
                        <Link
                            href={rowData.path ?? ""}
                            className="hover:text-neutral-700"
                            rel="noopener noreferrer"
                            target="_blank">
                            @{rowData.sid}
                        </Link>
                        {rowData.updatedBy ? (
                            <>
                                <p>updated {rowData.updatedAt} by</p>
                                <Link
                                    href={`/users/${rowData.updatedBy}`}
                                    className="hover:text-neutral-700">
                                    {rowData.updatedBy}
                                </Link>
                            </>
                        ) : (
                            <>
                                <p>created {rowData.createdAt} by</p>
                                <Link
                                    href={`/users/${rowData.createdBy}`}
                                    className="hover:text-neutral-700">
                                    {rowData.createdBy}
                                </Link>
                            </>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
};

export { OrgDashboard };
