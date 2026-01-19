import type { User } from "../types/user";

export function canEditResource(
  user: User | null,
  creatorId: string | number
) {
  return user !== null && String(user.id) === String(creatorId);
}
