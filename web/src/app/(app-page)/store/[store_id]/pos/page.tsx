import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";
import { getCategories, getProducts } from "./actions";
import POSController from "./controller";

export default async function POSPage() {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery({
    queryKey: ["items"],
    queryFn: getProducts,
  });
  await queryClient.prefetchQuery({
    queryKey: ["categories"],
    queryFn: getCategories,
  });

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      <POSController />
    </HydrationBoundary>
  );
}
