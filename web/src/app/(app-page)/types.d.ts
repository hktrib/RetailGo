type ItemWithoutId = {
  name: string;
  price: number;
  quantity: number;
  category: string;
};

type Category = {
  id: number;
  name: string;
  photo: string;
  store_id: number;
};

type Item = {
  category: string;
  category_id: number;
  category_name: string;
  date_last_sold: string;
  id: number;
  name: string;
  number_sold_since_update: number;
  photo: string;
  price: number;
  store_id: number;
  strip_price_id: string;
  stripe_product_id: string;
  weaviate_id: string;
};
