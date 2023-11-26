"use client";

import { SetStateAction, useEffect, useState } from "react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Minus, MoveRight, Plus } from "lucide-react";

type Product = {
  id: number;
  name: string;
  price: number;
};

const categories = [
  { id: 0, name: "Breakfast", items: 3 },
  { id: 4, name: "Lunch", items: 3 },
  { id: 1, name: "Vegetables", items: 3 },
  { id: 2, name: "Drinks", items: 3 },
  { id: 3, name: "Dairy", items: 3 },
  { id: 6, name: "Soups", items: 3 },
];

const products = [
  { id: 0, name: "Burger", price: 5.75, category: 4 },
  { id: 1, name: "Hamburger", price: 6, category: 4 },
  { id: 2, name: "Cheeseburger", price: 6.25, category: 4 },
  { id: 3, name: "Biscuits", price: 3.5, category: 0 },
  { id: 4, name: "Pancakes (x6)", price: 11, category: 0 },
  { id: 5, name: "Bacon (x4)", price: 6, category: 0 },
  { id: 6, name: "Broccoli", price: 0.5, category: 1 },
  { id: 7, name: "Brussel Sprouts", price: 2, category: 1 },
  { id: 8, name: "Cabbage", price: 20, category: 1 },
  { id: 9, name: "Coke", price: 2, category: 2 },
  { id: 10, name: "Sprite", price: 2, category: 2 },
  { id: 11, name: "Root Bear", price: 2, category: 2 },
  { id: 12, name: "Milk", price: 1, category: 3 },
  { id: 13, name: "Almond Milk", price: 2, category: 3 },
  { id: 14, name: "Oat Milk", price: 2, category: 3 },
  { id: 15, name: "Chicken Soup", price: 2, category: 6 },
  { id: 16, name: "Tomato Soup", price: 2, category: 6 },
  { id: 17, name: "Lentil Soup", price: 2, category: 6 },
];

const TAX_RATE = 1;

export default function POSPage() {
  const [selectedCategory, setSelectedCategory] = useState(-1);
  const [cart, setCart] = useState<(Product & { quantity: number })[]>([]);
  const [visibleProducts, setVisibleProducts] = useState(products);
  const [searchTerm, setSearchTerm] = useState("");

  const addItemToCart = (pid: number) => {
    console.log(pid, cart);
    // update qty
    if (cart.some((product) => product.id === pid)) {
      const i = cart.findIndex((i) => i.id === pid);
      const updatedProduct = { ...cart[i], quantity: cart[i].quantity + 1 };

      const newCart = [...cart];
      newCart[i] = updatedProduct;
      setCart(newCart);

      return;
    }

    // filter product by id
    const product = {
      ...products.filter((product) => product.id === pid)[0],
      quantity: 1,
    };
    setCart([...cart, product]);
  };

  const removeItemFromCart = (pid: number) => {
    // item not in cart, return
    if (!cart.some((product) => product.id === pid)) return;

    // delete item from cart
    const product = cart.filter((product) => product.id === pid)[0];
    if (product.quantity === 1) {
      setCart(cart.filter((product) => product.id !== pid));

      return;
    }

    // update qty
    const i = cart.findIndex((i) => i.id === pid);
    const updatedProduct = { ...cart[i], quantity: cart[i].quantity - 1 };

    const newCart = [...cart];
    newCart[i] = updatedProduct;
    setCart(newCart);
  };

  const fetchProductsByCategory = (categoryId: number) => {
    const productsToShow = products.filter(
      (product) => product.category === categoryId
    );

    setVisibleProducts([...productsToShow]);
  };

  useEffect(() => {
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
    return categories.filter((category) => category.id === categoryId)[0].name;
  };

  const fetchProductCartQty = (pid: number) => {
    if (!cart.some((product) => product.id === pid)) return;

    const i = cart.findIndex((i) => i.id === pid);
    return cart[i].quantity;
  };

  const handleSearchChange = (e: {
    target: { value: SetStateAction<string> };
  }) => {
    setSearchTerm(e.target.value);
  };

  return (
    <div className="py-5 px-8 h-full flex-grow flex flex-col w-full mx-auto max-w-7xl">
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
        <Input
          placeholder="Search items..."
          className="w-96 mt-2"
          value={searchTerm}
          onChange={handleSearchChange}
        />
      </div>

      <div className="flex-grow mt-6 mx-auto grid max-w-2xl w-full grid-cols-1 grid-rows-1 items-start gap-x-8 gap-y-8 lg:mx-0 lg:max-w-none lg:grid-cols-3">
        <div className="lg:col-start-3 lg:row-end-1">
          <OrderSummary cart={cart} />
        </div>

        <div className="lg:col-span-2 lg:row-span-2">
          <section className="grid lg:grid-cols-3 gap-3">
            <CategoryCard
              id={-1}
              name="All"
              items={products.length}
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

          <section className="grid xl:grid-cols-3 gap-3">
            {visibleProducts.map((product) => (
              <ProductCard
                key={product.id}
                productData={{
                  id: product.id,
                  name: product.name,
                  price: product.price,
                  category: fetchCategoryById(product.category),
                }}
                qty={fetchProductCartQty(product.id) ?? 0}
                addItem={addItemToCart}
                removeItem={removeItemFromCart}
              />
            ))}
          </section>
        </div>
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
        selected && "ring ring-blue-200"
      }`}
      onClick={() => setSelectedCategory(id)}
    >
      <div className="text-xl font-medium">{name}</div>
      <div className="text-sm text-gray-600 leading-6">
        {items} <span>{items === 1 ? "item" : "items"}</span>
      </div>
    </article>
  );
};

const ProductCard = ({
  productData,
  qty,
  addItem,
  removeItem,
}: {
  productData: Product & { category: string };
  qty: number;
  addItem: (pid: number) => void;
  removeItem: (pid: number) => void;
}) => {
  const { id, name, price, category } = productData;

  return (
    <article className="bg-gray-100 p-4 rounded-lg">
      <div className="text-xs text-gray-500">{category}</div>
      <div className="text-lg font-medium mt-1 leading-6">{name}</div>
      <div className="text-sm text-gray-600">${price}</div>

      <div className="flex items-center justify-end mt-3 gap-x-2">
        <Button
          variant="outline"
          className="px-2 h-8"
          onClick={() => removeItem(id)}
        >
          <Minus className="w-4 h-4" />
        </Button>
        <span>{qty}</span>
        <Button
          variant="outline"
          className="px-2 h-8"
          onClick={() => addItem(id)}
        >
          <Plus className="w-4 h-4" />
        </Button>
      </div>
    </article>
  );
};

const OrderSummary = ({
  cart,
}: {
  cart: (Product & { quantity: number })[];
}) => {
  const subtotal = cart.reduce(
    (acc, val) => acc + parseFloat((val.price * val.quantity).toFixed(2)),
    0
  );

  const total = parseFloat((subtotal * TAX_RATE).toFixed(2));

  return (
    <div className="h-full">
      {cart.length ? (
        <div className="space-y-2">
          {cart.map((cartItem, idx) => (
            <div
              key={cartItem.id}
              className="flex items-center justify-between bg-gray-100 rounded-lg p-3"
            >
              <div className="flex items-center gap-x-2 text-sm">
                <div className="rounded-full bg-black text-gray-50 w-5 h-5 flex items-center justify-center">
                  <span className="text-xs">{idx + 1}</span>
                </div>
                <span>{cartItem.name}</span>
                <span className="text-gray-600">x{cartItem.quantity}</span>
              </div>

              <div>
                <span className="text-sm">
                  ${(cartItem.price * cartItem.quantity).toFixed(2)}
                </span>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div className="bg-gray-100 rounded-lg p-3">
          No items added to cart.
        </div>
      )}

      <div className="mt-6 p-3 bg-gray-100 rounded-lg">
        <div>
          <div className="flex items-center justify-between">
            <span className="text-sm text-gray-700">Subtotal</span>
            <span className="text-sm">${subtotal.toFixed(2)}</span>
          </div>

          <div className="flex items-center justify-between mt-1">
            <span className="text-sm text-gray-700">Tax</span>
            <span className="text-sm">
              $
              {(
                subtotal - parseFloat((subtotal * TAX_RATE).toFixed(2))
              ).toFixed(2)}
            </span>
          </div>

          <hr className="my-3" />

          <div className="flex items-center justify-between mt-1">
            <span className="text-gray-700">Total</span>
            <span className="font-medium">${total.toFixed(2)}</span>
          </div>

          <div className="mt-6">
            <Button className="rounded-full w-full py-5">Place order</Button>
          </div>
        </div>
      </div>
    </div>
  );
};
