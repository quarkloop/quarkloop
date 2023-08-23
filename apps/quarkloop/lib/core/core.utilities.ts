import { PipelineArgs, PipelineState } from "@/lib/core/pipeline";
import { customAlphabet } from "nanoid";

export const generateId = customAlphabet(
  "abcdefghijklmnopqrstuvwxyz0123456789",
  10
);

// https://stackoverflow.com/a/14919494/1789729
/**
 * Format bytes as human-readable text.
 *
 * @param bytes Number of bytes.
 * @param si True to use metric (SI) units, aka powers of 1000. False to use
 *           binary (IEC), aka powers of 1024.
 * @param dp Number of decimal places to display.
 *
 * @return Formatted string.
 */
export function humanFileSize(bytes: number, si = false, dp = 1) {
  const thresh = si ? 1000 : 1024;

  if (Math.abs(bytes) < thresh) {
    return bytes + " B";
  }

  const units = ["KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

  // const units = si
  //   ? ["kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"]
  //   : ["KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"];
  let u = -1;
  const r = 10 ** dp;

  do {
    bytes /= thresh;
    ++u;
  } while (
    Math.round(Math.abs(bytes) * r) / r >= thresh &&
    u < units.length - 1
  );

  return bytes.toFixed(dp) + " " + units[u];
}

export type Optional<T, K extends keyof T> = Pick<Partial<T>, K> & Omit<T, K>;

export type ToString<T> = T extends Date
  ? string
  : T extends object
  ? TypePropertiesToString<T>
  : string;

export type TypePropertiesToString<T> = {
  [K in keyof T]: ToString<T[K]>;
};

// https://stackoverflow.com/a/1349426/1789729
export function makeUid(length: number) {
  let result = "";
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  const charactersLength = characters.length;
  let counter = 0;
  while (counter < length) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
    counter += 1;
  }
  return result;
}
