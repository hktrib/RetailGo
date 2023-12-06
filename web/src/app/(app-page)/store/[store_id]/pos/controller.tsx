"use client";

import { useState, useEffect } from "react";

import POSHeader from "./header";
import POSCategories from "./categories";
import POSProducts from "./products";
import POSOrderSummary from "./order-summary";

const TAX_RATE = 1;

const POSController = ({
  categories,
  items,
}: {
  categories: Category[];
  items: Item[];
}) => {
  const [selectedCategory, setSelectedCategory] = useState(-1);
  const [cart, setCart] = useState<(Item & { quantity: number })[]>([]);
  const [visibleProducts, setVisibleProducts] = useState(items ?? []);
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    if (!items) return;
    if (selectedCategory === -2) return;
    if (selectedCategory === -1) {
      setVisibleProducts(items);
      return;
    }

    const filteredProducts = items.filter(
      (product) => product.category_id === selectedCategory
    );

    setVisibleProducts(filteredProducts);
  }, [selectedCategory, items]);

  useEffect(() => {
    if (searchTerm.length === 0) {
      setSelectedCategory(-1);
      return;
    }

    const query = searchTerm.toLowerCase();
    const filteredProducts = items.filter(
      (product) =>
        product.name.toLowerCase().includes(query) ||
        product.id.toString().toLowerCase().includes(query)
    );

    setSelectedCategory(-2);
    setVisibleProducts(filteredProducts);
  }, [searchTerm, items]);

  const fetchCategoryById = (categoryId: number) => {
    if (!categories) return "";

    return categories.filter((category) => category.id === categoryId)[0].name;
  };

  return (
    <div>
      <div className="py-5 px-8 h-full flex-grow flex flex-col w-full mx-auto max-w-2xl lg:max-w-7xl items-center lg:items-start">
        <POSHeader
          searchTerm={searchTerm}
          setSearchTerm={setSearchTerm}
          selectedCategory={selectedCategory}
        />

        <div className="flex-grow mt-6 mx-auto grid max-w-2xl w-full grid-cols-1 grid-rows-1 items-start gap-x-8 gap-y-8 lg:mx-0 lg:max-w-none lg:grid-cols-3">
          <div className="lg:col-start-3 lg:row-end-1">
            <POSOrderSummary cart={cart} TAX_RATE={TAX_RATE} />
          </div>

          <div className="lg:col-span-2 lg:row-span-2">
            <POSCategories
              items={items}
              categories={categories}
              selectedCategory={selectedCategory}
              setSelectedCategory={setSelectedCategory}
            />

            <hr className="my-6" />

            <POSProducts
              products={items}
              visibleProducts={visibleProducts}
              cart={cart}
              setCart={setCart}
              fetchCategoryById={fetchCategoryById}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default POSController;
