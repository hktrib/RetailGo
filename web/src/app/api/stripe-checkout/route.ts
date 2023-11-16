import { NextApiRequest, NextApiResponse } from "next";
import Stripe from "stripe";

const stripe = new Stripe(process.env.STRIPE_SECRET_KEY!);

export default async function POST(
  req: NextApiRequest,
  res: NextApiResponse
) {
  try {
    const { lineItems } = req.body;

    if (!lineItems) {
      throw new Error("No items in cart");
    }

    const session = await stripe.checkout.sessions.create({
      line_items: [
        {
          quantity: lineItems.quantity,
          price_data: {
            unit_amount: lineItems.price * 100,
            currency: "usd",
            product_data: {
              name: "name",
            },
          },
        },
      ],
      mode: "payment",
      success_url: `${req.headers.origin}`,
      cancel_url: `${req.headers.origin}`,
    });

    res.status(200).json({ sessionId: session.id });
  } catch (err) {
    res.status(500).json({
      statusCode: 500,
      message: err instanceof Error ? err.message : "Internal server error",
    });
  }
}