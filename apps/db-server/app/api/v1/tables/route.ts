import { NextResponse } from "next/server";

export async function GET(request: Request, { params }: { params: any }) {
  return NextResponse.json(
    {
      tables: [
        "App",
        "AppInstance",
        "AppThread",
        "AppThreadSettings",
        "AppFile",
        "AppFileSettings",
        "AppForm",
        "AppFormSettings",
        "AppPage",
        "AppPageSettings",
      ],
    },
    { status: 200 }
  );
}
