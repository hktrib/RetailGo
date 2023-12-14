import { cx } from "class-variance-authority";

const POSCategories = ({
  items,
  categories,
  selectedCategory,
  setSelectedCategory,
}: {
  items: Item[];
  categories: Category[];
  selectedCategory: string;
  setSelectedCategory: React.Dispatch<React.SetStateAction<string>>;
}) => {
  const fetchNumProductsInCategory = (categoryName: string) => {
    return items.filter((item) => item.category_name === categoryName).length;
  };

  return (
    <section className="grid gap-3 lg:grid-cols-3">
      <CategoryCard
        name="-1"
        displayName="All"
        items={items.length}
        selected={selectedCategory === "-1"}
        setSelectedCategory={setSelectedCategory}
      />

      {categories.map((category) => (
        <CategoryCard
          key={category.id}
          displayName={category.name}
          name={category.name}
          items={fetchNumProductsInCategory(category.name)}
          selected={selectedCategory === category.name}
          setSelectedCategory={setSelectedCategory}
        />
      ))}
    </section>
  );
};

const CategoryCard = ({
  name,
  displayName,
  items,
  selected,
  setSelectedCategory,
}: {
  name: string;
  displayName: string;
  items: number;
  selected: boolean;
  setSelectedCategory: (name: string) => void;
}) => {
  return (
    <article
      className={cx(
        "cursor-pointer rounded-lg bg-white p-4 shadow-sm dark:bg-zinc-800",
        selected && "ring-2 ring-sky-300 dark:ring-blue-500",
      )}
      onClick={() => setSelectedCategory(name)}
    >
      <div className="text-xl font-medium">{displayName}</div>
      <div className="text-sm leading-6 text-gray-600 dark:text-zinc-300">
        {items} <span>{items === 1 ? "item" : "items"}</span>
      </div>
    </article>
  );
};

export default POSCategories;
