"use client";

import { useState } from "react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Minus, MoveRight, Plus } from "lucide-react";

const categories = [
  { id: 0, name: "Breakfast", items: 13 },
  { id: 1, name: "Breakfast", items: 13 },
  { id: 2, name: "Breakfast", items: 13 },
  { id: 3, name: "Breakfast", items: 13 },
  { id: 4, name: "Breakfast", items: 13 },
  { id: 5, name: "Breakfast", items: 13 },
  { id: 6, name: "Breakfast", items: 13 },
];

export default function POSPage() {
  const [selectedCategory, setSelectedCategory] = useState(-1);

  return (
    <div className="py-5 px-8">
      <div>
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
        <Input placeholder="Search items..." className="w-96 mt-2" />
      </div>

      <section className="grid xl:grid-cols-4 gap-3 mt-6">
        <CategoryCard
          id={-1}
          name="All"
          items={13}
          selected={selectedCategory === -1}
          setSelectedCategory={setSelectedCategory}
        />
        {categories.map((category) => (
          <CategoryCard
            key={category.id}
            id={category.id}
            name={category.name}
            items={category.items}
            selected={selectedCategory === category.id}
            setSelectedCategory={setSelectedCategory}
          />
        ))}
      </section>

      <hr className="my-6" />

      <div>
        <section className="grid xl:grid-cols-4 gap-3">
          <ProductCard name="Burger" price={5.75} category="Breakfast" />
          <ProductCard name="Burger" price={5.75} category="Breakfast" />
          <ProductCard name="Burger" price={5.75} category="Breakfast" />
        </section>
      </div>
    </div>
  );
}

const CategoryCard = ({
  id,
  name,
  items,
  selected,
  setSelectedCategory,
}: {
  id: number;
  name: string;
  items: number;
  selected: boolean;
  setSelectedCategory: (id: number) => void;
}) => {
  return (
    <article
      className={`bg-gray-100 p-4 rounded-lg cursor-pointer ${
        selected && "ring ring-orange-400"
      }`}
      onClick={() => setSelectedCategory(id)}
    >
      <div className="text-2xl font-medium">{name}</div>
      <div className="text-sm text-gray-600 leading-6">
        {items} <span>{items === 1 ? "item" : "items"}</span>
      </div>
    </article>
  );
};

const ProductCard = ({
  name,
  price,
  category,
}: {
  name: string;
  price: number;
  category: string;
}) => {
  return (
    <article className="bg-gray-100 p-4 rounded-lg">
      <div className="text-xs text-gray-500">{category}</div>
      <div className="text-xl font-medium mt-1">{name}</div>
      <div className="text-sm text-gray-600 leading-6">${price}</div>

      <div className="flex items-center justify-end mt-2 gap-x-2">
        <Button variant="outline" className="px-2 h-8">
          <Minus className="w-4 h-4" />
        </Button>
        <span>0</span>
        <Button variant="outline" className="px-2 h-8">
          <Plus className="w-4 h-4" />
        </Button>
      </div>
    </article>
  );
};
