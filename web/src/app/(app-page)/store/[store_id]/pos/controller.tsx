"use client";

import { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import { getProducts, getCategories } from "./actions";

import POSCategories from "./categories";
import POSHeader from "./header";
import POSProducts from "./products";
import POSOrderSummary from "./order-summary";

export type Category = { id: number; name: string; items: number };
export type Item = {
  id: number;
  name: string;
  price: number;
  category: number;
};

const TAX_RATE = 1;

const POSController = () => {
  const {
    data: products,
    isPending: isProductsPending,
    isError: isProductsError,
  } = useQuery({
    queryKey: ["items"],
    queryFn: getProducts,
  });
  const {
    data: categories,
    isPending: isCategoriesPending,
    isError: isCategoriesError,
  } = useQuery({
    queryKey: ["categories"],
    queryFn: getCategories,
  });

  const [selectedCategory, setSelectedCategory] = useState(-1);
  const [cart, setCart] = useState<(Item & { quantity: number })[]>([]);
  const [visibleProducts, setVisibleProducts] = useState(products ?? []);
  const [searchTerm, setSearchTerm] = useState("");

  useEffect(() => {
    if (!products) return;

    let filteredProducts = products;

    if (selectedCategory !== -1) {
      filteredProducts = filteredProducts.filter(
        (product) => product.category === selectedCategory
      );
    }

    if (searchTerm) {
      const query = searchTerm.toLowerCase();

      filteredProducts = filteredProducts.filter(
        (product) =>
          product.name.toLowerCase().includes(query) ||
          product.id.toString().toLowerCase().includes(query)
      );
    }

    setVisibleProducts(filteredProducts);
  }, [selectedCategory, searchTerm]);

  const fetchCategoryById = (categoryId: number) => {
    if (!categories) return "";

    return categories.filter((category) => category.id === categoryId)[0].name;
  };

  if (isProductsPending || isCategoriesPending) {
    return <span>loading...</span>;
  }

  if (isProductsError || isCategoriesError) {
    return <span>error...</span>;
  }

  return (
    <div>
      <div className="py-5 px-8 h-full flex-grow flex flex-col w-full mx-auto max-w-2xl lg:max-w-7xl items-center lg:items-start">
        <POSHeader
          searchTerm={searchTerm}
          setSearchTerm={setSearchTerm}
          categories={categories}
          selectedCategory={selectedCategory}
        />

        <div className="flex-grow mt-6 mx-auto grid max-w-2xl w-full grid-cols-1 grid-rows-1 items-start gap-x-8 gap-y-8 lg:mx-0 lg:max-w-none lg:grid-cols-3">
          <div className="lg:col-start-3 lg:row-end-1">
            <POSOrderSummary cart={cart} TAX_RATE={TAX_RATE} />
          </div>

          <div className="lg:col-span-2 lg:row-span-2">
            <POSCategories
              numProducts={products.length}
              categories={categories}
              selectedCategory={selectedCategory}
              setSelectedCategory={setSelectedCategory}
            />

            <hr className="my-6" />

            <POSProducts
              products={products}
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
