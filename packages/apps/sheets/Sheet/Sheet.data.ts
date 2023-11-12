import {
    ArrowDownIcon,
    ArrowRightIcon,
    ArrowUpIcon,
    CircleIcon,
} from "lucide-react";
import { SheetRow } from "./Sheet.type";

export const labels = [
    {
        value: "bug",
        label: "Bug",
    },
    {
        value: "feature",
        label: "Feature",
    },
    {
        value: "documentation",
        label: "Documentation",
    },
];

export const statuses = [
    {
        value: "backlog",
        label: "Backlog",
        icon: CircleIcon,
    },
    {
        value: "todo",
        label: "Todo",
        icon: CircleIcon,
    },
    {
        value: "in progress",
        label: "In Progress",
        icon: CircleIcon,
    },
    {
        value: "done",
        label: "Done",
        icon: CircleIcon,
    },
    {
        value: "canceled",
        label: "Canceled",
        icon: CircleIcon,
    },
];

export const priorities = [
    {
        label: "Low",
        value: "low",
        icon: ArrowDownIcon,
    },
    {
        label: "Medium",
        value: "medium",
        icon: ArrowRightIcon,
    },
    {
        label: "High",
        value: "high",
        icon: ArrowUpIcon,
    },
];

export const data: SheetRow[] = [
    {
        id: "TASK-8782",
        name: "You can't compress the program without quantifying the open-source SSD pixel!",
        path: "",
        description: "documentation",
    },
    {
        id: "TASK-7878",
        name: "Try to calculate the EXE feed, maybe it will index the multi-byte pixel!",
        path: "",
        description: "documentation",
    },
    {
        id: "TASK-7839",
        name: "We need to bypass the neural TCP card!",
        path: "",
        description: "bug",
    },
];
