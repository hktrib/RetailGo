const POSCategories = ({
  items,
  categories,
  selectedCategory,
  setSelectedCategory,
}: {
  items: Item[];
  categories: Category[];
  selectedCategory: number;
  setSelectedCategory: React.Dispatch<React.SetStateAction<number>>;
}) => {
  const fetchNumProductsInCategory = (categoryId: number) => {
    return items.filter((item) => item.category_id === categoryId).length;
  };

  return (
    <section className="grid lg:grid-cols-3 gap-3">
      <CategoryCard
        id={-1}
        name="All"
        items={items.length}
        selected={selectedCategory === -1}
        setSelectedCategory={setSelectedCategory}
      />

      {categories.map((category) => (
        <CategoryCard
          key={category.id}
          id={category.id}
          name={category.name}
          items={fetchNumProductsInCategory(category.id)}
          selected={selectedCategory === category.id}
          setSelectedCategory={setSelectedCategory}
        />
      ))}
    </section>
  );
};

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

export default POSCategories;
