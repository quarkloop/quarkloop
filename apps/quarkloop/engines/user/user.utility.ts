import { UserAccountStatus } from "./user.type";

export function getStatusString(status: UserAccountStatus): string {
  switch (status) {
    case UserAccountStatus.Inactive:
      return "Inactive";
    case UserAccountStatus.Active:
      return "Active";
    case UserAccountStatus.Suspended:
      return "Suspended";
    case UserAccountStatus.Deleted:
      return "Deleted";
    default:
      throw new Error("Invalid account status");
  }
}
