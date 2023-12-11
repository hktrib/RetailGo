"use client";

import { useState, useEffect } from "react";

import POSCategories from "./pos-categories";
import POSProducts from "./pos-products";
import POSOrderSummary from "./pos-order-summary";
import { Input } from "@/components/ui/input";

const TAX_RATE = 1;

const POSController = ({
  categories,
  items,
  storeId,
}: {
  categories: Category[];
  items: Item[];
  storeId: string;
}) => {
  const [selectedCategory, setSelectedCategory] = useState("-1");
  const [cart, setCart] = useState<(Item & { quantityAdded: number })[]>([]);
  const [visibleProducts, setVisibleProducts] = useState(items ?? []);
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    if (!items) return;
    if (selectedCategory === "-2") return;
    if (selectedCategory === "-1") {
      setVisibleProducts(items);
      return;
    }

    const filteredProducts = items.filter(
      (product) => product.category_name === selectedCategory,
    );

    setVisibleProducts(filteredProducts);
  }, [selectedCategory, items]);

  useEffect(() => {
    if (searchTerm.length === 0) {
      setSelectedCategory("-1");
      return;
    }

    const query = searchTerm.toLowerCase();
    const filteredProducts = items.filter(
      (product) =>
        product.name.toLowerCase().includes(query) ||
        product.id.toString().toLowerCase().includes(query),
    );

    setSelectedCategory("-2");
    setVisibleProducts(filteredProducts);
  }, [searchTerm, items]);

  const fetchCategoryById = (categoryId: number) => {
    if (!categories) return "";

    return categories.filter((category) => category.id === categoryId)[0].name;
  };

  return (
    <div className="mx-auto flex h-full w-full max-w-2xl flex-grow flex-col items-center px-4 md:px-6 lg:max-w-7xl lg:items-start xl:px-8">
      <div className="w-full lg:w-auto">
        <Input
          placeholder="Search items..."
          className="mt-2 w-96 dark:border-zinc-800 dark:focus:ring-zinc-700"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>

      <div className="mx-auto mt-6 grid w-full max-w-2xl flex-grow grid-cols-1 grid-rows-1 items-start gap-x-8 gap-y-8 lg:mx-0 lg:max-w-none lg:grid-cols-3">
        <div className="lg:col-start-3 lg:row-end-1">
          <POSOrderSummary cart={cart} TAX_RATE={TAX_RATE} storeId={storeId} />
        </div>

        <div className="lg:col-span-2 lg:row-span-2">
          <POSCategories
            items={items}
            categories={categories}
            selectedCategory={selectedCategory}
            setSelectedCategory={setSelectedCategory}
          />

          <hr className="my-6 dark:border-zinc-800" />

          <POSProducts
            products={items}
            visibleProducts={visibleProducts}
            cart={cart}
            setCart={setCart}
          />
        </div>
      </div>
    </div>
  );
};

export default POSController;
