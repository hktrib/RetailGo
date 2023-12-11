import { Button } from "@/components/ui/button";
import { Minus, Plus } from "lucide-react";

const POSProducts = ({
  products,
  visibleProducts,
  cart,
  setCart,
  fetchCategoryById,
}: {
  products: Item[];
  visibleProducts: Item[];
  cart: (Item & { quantityAdded: number })[];
  setCart: React.Dispatch<
    React.SetStateAction<(Item & { quantityAdded: number })[]>
  >;
  fetchCategoryById: (categoryId: number) => string;
}) => {
  const addItemToCart = (pid: number) => {
    console.log(pid, cart);
    // update qty
    if (cart.some((product) => product.id === pid)) {
      const i = cart.findIndex((i) => i.id === pid);
      const updatedProduct = {
        ...cart[i],
        quantityAdded: cart[i].quantityAdded + 1,
      };

      const newCart = [...cart];
      newCart[i] = updatedProduct;
      setCart(newCart);

      return;
    }

    // filter product by id
    const product = {
      ...products.filter((product) => product.id === pid)[0],
      quantityAdded: 1,
    };
    setCart([...cart, product]);
  };

  const removeItemFromCart = (pid: number) => {
    // item not in cart, return
    if (!cart.some((product) => product.id === pid)) return;

    // delete item from cart
    const product = cart.filter((product) => product.id === pid)[0];
    if (product.quantityAdded === 1) {
      setCart(cart.filter((product) => product.id !== pid));

      return;
    }

    // update qty
    const i = cart.findIndex((i) => i.id === pid);
    const updatedProduct = {
      ...cart[i],
      quantityAdded: cart[i].quantityAdded - 1,
    };

    const newCart = [...cart];
    newCart[i] = updatedProduct;
    setCart(newCart);
  };

  const fetchProductCartQty = (pid: number) => {
    if (!cart.some((product) => product.id === pid)) return;

    const i = cart.findIndex((i) => i.id === pid);
    return cart[i].quantityAdded;
  };

  return (
    <section className="grid gap-3 xl:grid-cols-3">
      {visibleProducts.map((product) => (
        <ProductCard
          key={product.id}
          productData={{
            id: product.id,
            name: product.name,
            price: product.price,
            category: product.category_name,
          }}
          qty={fetchProductCartQty(product.id) ?? 0}
          addItem={addItemToCart}
          removeItem={removeItemFromCart}
        />
      ))}
    </section>
  );
};

const ProductCard = ({
  productData,
  qty,
  addItem,
  removeItem,
}: {
  productData: { id: number; name: string; price: number; category: string };
  qty: number;
  addItem: (pid: number) => void;
  removeItem: (pid: number) => void;
}) => {
  const { id, name, price, category } = productData;

  return (
    <article className="rounded-lg bg-gray-100 p-4 shadow-sm dark:bg-zinc-800">
      <div className="text-xs text-gray-500 dark:text-zinc-400">{category}</div>
      <div className="mt-0.5 text-lg font-medium leading-6">{name}</div>
      <div className="text-sm leading-6 text-gray-600 dark:text-zinc-300">
        ${price}
      </div>

      <div className="mt-3 flex items-center justify-end gap-x-2">
        <Button
          variant="outline"
          className="h-8 px-2 dark:border-zinc-700 dark:hover:bg-zinc-700"
          onClick={() => removeItem(id)}
        >
          <Minus className="h-4 w-4" />
        </Button>
        <span>{qty}</span>
        <Button
          variant="outline"
          className="h-8 px-2 dark:border-zinc-700 dark:hover:bg-zinc-700"
          onClick={() => addItem(id)}
        >
          <Plus className="h-4 w-4" />
        </Button>
      </div>
    </article>
  );
};

export default POSProducts;
