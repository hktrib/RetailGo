import { Input } from "@/components/ui/input";
import { MoveRight } from "lucide-react";

import type { Category } from "./controller";

const POSHeader = ({
  searchTerm,
  setSearchTerm,
  categories,
  selectedCategory,
}: {
  searchTerm: string;
  setSearchTerm: React.Dispatch<React.SetStateAction<string>>;
  categories: Category[];
  selectedCategory: number;
}) => {
  return (
    <div className="w-full lg:w-auto">
      <div>
        {selectedCategory === -1 ? (
          <span className="text-gray-700 flex items-center text-sm">
            All products
          </span>
        ) : (
          <div className="text-gray-700 flex items-center gap-x-2 text-sm">
            <span>Category1</span>
            <MoveRight className="w-4 h-4" />
            <span>Category2</span>
          </div>
        )}
      </div>
      <Input
        placeholder="Search items..."
        className="w-96 mt-2"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
    </div>
  );
};

export default POSHeader;
