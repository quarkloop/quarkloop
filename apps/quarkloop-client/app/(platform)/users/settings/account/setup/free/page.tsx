import { redirect } from "next/navigation";

const FreeAccountSetup = async (props: any) => {
    const {
        searchParams: { plan, next },
    } = props;

    // const pipeline = createPipeline<PipelineState, PipelineArgs[]>({
    //     initialState: {},
    // });

    // const planState = await pipeline
    //     .use(GetPlansPlugin)
    //     .use(GetApiResponsePlugin)
    //     .onError(DefaultErrorHandler)
    //     .execute({
    //         plan: { planType: "Free" } as GetPlansPluginArgs,
    //     });

    // const plans = planState.database?.plan?.records as Plan[] | undefined;

    // const finalState = await pipeline
    //     .use(GetServerSessionPlugin)
    //     .use(GetUserPlugin)
    //     .use(CreatePlanSubscriptionPlugin)
    //     .use(GetApiResponsePlugin)
    //     .onError(DefaultErrorHandler)
    //     .execute({
    //         planSubscription: {
    //             startDate: new Date(),
    //             status: "Active",
    //             planId: plans?.find((plan) => plan)?.id,
    //         } as CreatePlanSubscriptionPluginArgs,
    //     });

    // const planSubscription = finalState.database?.planSubscription?.records as
    //     | PlanSubscription
    //     | undefined;

    redirect("/");
};

export default FreeAccountSetup;
