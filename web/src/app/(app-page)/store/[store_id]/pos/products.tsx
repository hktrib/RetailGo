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
    <section className="grid xl:grid-cols-3 gap-3">
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

export default POSProducts;
