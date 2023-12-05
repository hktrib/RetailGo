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

const categories = [
  { id: 0, name: "Breakfast", items: 3 },
  { id: 4, name: "Lunch", items: 3 },
  { id: 1, name: "Vegetables", items: 3 },
  { id: 2, name: "Drinks", items: 3 },
  { id: 3, name: "Dairy", items: 3 },
  { id: 6, name: "Soups", items: 3 },
];

export const getProducts = () => {
  console.log("fetching products...");
  return products;
};

export const getCategories = () => {
  console.log("fetching categories...");
  return categories;
};
