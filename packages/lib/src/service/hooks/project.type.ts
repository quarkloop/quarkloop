import { z } from "zod";
import { apiResponseV2Schema, projectSchema, Project } from "@quarkloop/types";

/// GetProjectList
export const getProjectListSchema = apiResponseV2Schema.merge(
    z.object({
        data: z.array(projectSchema),
    })
);
export type GetProjectListApiResponse = z.infer<typeof getProjectListSchema>;
export type GetProjectListApiArgs = {
    orgId: string[];
    workspaceId: string[];
};

/// GetProjectById
export const getProjectByIdSchema = apiResponseV2Schema.merge(
    z.object({
        data: projectSchema,
    })
);
export type GetProjectByIdApiResponse = z.infer<typeof getProjectByIdSchema>;
export type GetProjectByIdApiArgs = {
    id: string;
};

/// CreateProject
export const createProjectSchema = apiResponseV2Schema.merge(
    z.object({
        data: projectSchema,
    })
);
export type CreateProjectApiResponse = z.infer<typeof createProjectSchema>;
export type CreateProjectApiArgs = {
    orgId: string;
    workspaceId: string;
    project: Partial<Project>;
};

/// UpdateProject
export const updateProjectSchema = apiResponseV2Schema.merge(
    z.object({
        data: projectSchema,
    })
);
export type UpdateProjectApiResponse = z.infer<typeof updateProjectSchema>;
export type UpdateProjectApiArgs = {
    id: string;
    project: Partial<Project>;
};

/// DeleteProject
export const deleteProjectSchema = apiResponseV2Schema;
export type DeleteProjectApiResponse = z.infer<typeof deleteProjectSchema>;
export type DeleteProjectApiArgs = {
    id: string;
};
