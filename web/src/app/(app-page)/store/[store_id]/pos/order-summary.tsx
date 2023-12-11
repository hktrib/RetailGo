import { useState, useEffect } from "react";
import { createCheckout } from "./actions";
import { CheckoutDialog } from "@/components/app-page/checkout-dialog";
import { Button } from "@/components/ui/button";

const POSOrderSummary = ({
  cart,
  TAX_RATE,
  storeId,
}: {
  cart: (Item & { quantityAdded: number })[];
  TAX_RATE: number;
  storeId: string;
}) => {
  const [checkoutOpen, setCheckoutOpen] = useState(false);
  const [clientSecret, setClientSecret] = useState("");

  const subtotal = cart.reduce(
    (acc, val) => acc + parseFloat((val.price * val.quantityAdded).toFixed(2)),
    0,
  );

  const total = parseFloat((subtotal * TAX_RATE).toFixed(2));

  const generateCheckoutEmbed = async () => {
    const lineItems = cart.map((item) => {
      return {
        id: item.id,
        quantity: item.quantityAdded,
      };
    });

    const res = await createCheckout({
      lineItems: lineItems,
      store_id: storeId,
    });

    if (!res) {
      console.error("something went wrong...");
      return;
    }

    const newClientSecret = JSON.parse(res).ClientSecret;

    setClientSecret(newClientSecret);
  };

  useEffect(() => {
    if (!clientSecret) {
      setCheckoutOpen(false);
      return;
    }

    setCheckoutOpen(true);
  }, [clientSecret]);

  return (
    <div className="h-full">
      {cart.length ? (
        <div className="space-y-2">
          {cart.map((cartItem, idx) => (
            <div
              key={`${cartItem.id}-x${cartItem.quantityAdded}`}
              className="flex items-center justify-between rounded-lg bg-gray-100 p-3 dark:bg-zinc-800"
            >
              <div className="flex items-center gap-x-2 text-sm">
                <div className="flex h-6 w-6 items-center justify-center rounded-full bg-black text-gray-50 dark:bg-zinc-700">
                  <span className="text-xs">{idx + 1}</span>
                </div>
                <span>{cartItem.name}</span>
                <span className="text-gray-600 dark:text-zinc-300">
                  x{cartItem.quantityAdded}
                </span>
              </div>

              <div>
                <span className="text-sm">
                  ${(cartItem.price * cartItem.quantityAdded).toFixed(2)}
                </span>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div className="rounded-lg bg-gray-100 p-3 dark:bg-zinc-800">
          No items added to cart.
        </div>
      )}

      <div className="mt-6 rounded-lg bg-gray-100 p-3 dark:bg-zinc-800">
        <div>
          <div className="flex items-center justify-between">
            <span className="text-sm text-gray-700 dark:text-zinc-300">
              Subtotal
            </span>
            <span className="text-sm dark:text-zinc-200">
              ${subtotal.toFixed(2)}
            </span>
          </div>

          <div className="mt-1 flex items-center justify-between">
            <span className="text-sm text-gray-700 dark:text-zinc-300">
              Tax
            </span>
            <span className="text-sm dark:text-zinc-200">
              $
              {(
                subtotal - parseFloat((subtotal * TAX_RATE).toFixed(2))
              ).toFixed(2)}
            </span>
          </div>

          <hr className="my-3" />

          <div className="mt-1 flex items-center justify-between">
            <span className="text-gray-700 dark:text-zinc-200">Total</span>
            <span className="font-medium">${total.toFixed(2)}</span>
          </div>

          <div className="mt-6">
            <Button
              type="submit"
              disabled={!cart.length}
              onClick={generateCheckoutEmbed}
              className="w-full rounded-full py-5"
            >
              Place order
            </Button>
          </div>
        </div>
      </div>

      <CheckoutDialog
        open={checkoutOpen}
        setOpen={setCheckoutOpen}
        clientSecret={clientSecret}
        setClientSecret={setClientSecret}
      />
    </div>
  );
};

export default POSOrderSummary;
