import { redirect } from "next/navigation";

const AccountSetup = (props: any) => {
  const {
    searchParams: { callbackUrl },
  } = props;

  if (callbackUrl == null) {
    redirect("/");
  }

  const { searchParams } = new URL(callbackUrl);
  const plan = searchParams.get("plan"); // standard, pro, enterprise
  const next = searchParams.get("next"); // setup, upgrade, downgrade

  if (plan && next) {
    if (plan === "free") {
      redirect(`/user/account-setup/free?next=${next}`);
    } else if (plan === "standard") {
      redirect(`/user/account-setup/standard?next=${next}`);
    } else if (plan === "pro") {
      redirect(`/user/account-setup/pro?next=${next}`);
    } else if (plan === "enterprise") {
      redirect(`/user/account-setup/enterprise?next=${next}`);
    }
  }

  redirect("/");
};

export default AccountSetup;
