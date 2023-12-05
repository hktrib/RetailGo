// SelectedStoreContext.tsx

import { IStore } from '@/models/store';
import React, { createContext, useContext, useState, ReactNode } from 'react';
import isAuth from './isAuth';

interface SelectedStoreContextType {
  selectedStore: IStore | null;
  selectStore: (store: IStore) => void;
}

const SelectedStoreContext = createContext<SelectedStoreContextType | undefined>(
  undefined
);

export const useSelectedStore = (): SelectedStoreContextType => {
  const context = useContext(SelectedStoreContext);
  if (context === undefined) {
    throw new Error('useSelectedStore must be used within a SelectedStoreProvider');
  }
  return context;
};

interface SelectedStoreProviderProps {
  children: ReactNode;
}

const SelectedStoreProvider = ({
  children,
}: SelectedStoreProviderProps): JSX.Element => {
  const [selectedStore, setSelectedStore] = useState<IStore | null>(null);

  const selectStore = (store: IStore) => {
    setSelectedStore(store);
  };

  return (
    <SelectedStoreContext.Provider value={{ selectedStore, selectStore }}>
      {children}
    </SelectedStoreContext.Provider>
  );
};

export default isAuth(SelectedStoreProvider);
