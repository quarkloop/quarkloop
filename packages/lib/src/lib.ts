import { customAlphabet } from "nanoid";

export const generateId = customAlphabet(
  "abcdefghijklmnopqrstuvwxyz0123456789",
  10
);
