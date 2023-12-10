type ItemWithoutId = {
  name: string;
  price: number;
  quantity: number;
  category: string;
};

type Category = {
  id: number;
  name: string;
};

type Item = {
  category_name: string;
  id: number;
  name: string;
  photo?: string;
  price: number;
  quantity: number;
  stripe_price_id?: string;
  stripe_product_id?: string;
};

type Employee = {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  name: string;
  employeeId: string;
  position: string;
  department: string;
  hireDate: string;
  location: string;
};
