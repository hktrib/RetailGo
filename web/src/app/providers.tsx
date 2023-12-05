import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { createContext, useContext, useState } from 'react';
import SelectedStoreProvider  from '../components/storeprovider';

export default function Providers({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000,
          },
        },
      }),
  )


  return (
    <QueryClientProvider client={queryClient}>
      <SelectedStoreProvider>
      {children}
      </SelectedStoreProvider>
    </QueryClientProvider>
  )
}

