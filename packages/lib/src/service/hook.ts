import { useDispatch, useSelector, TypedUseSelectorHook } from "react-redux";
import type { AppState, AppDispatch } from "./store";

// Use throughout your project instead of plain `useDispatch` and `useSelector`
export const useAppSelector: TypedUseSelectorHook<AppState> = useSelector;
export const useAppDispatch: () => AppDispatch = useDispatch;
