import React from "react";
// import styles from "./Search.module.css";

interface SearchProps {
  value: string;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

const Search = ({ value, onChange }: SearchProps) => {
  return (
    <div>
      <search />
      <input
        type="text"
        placeholder="Search products"
        value={value}
        onChange={onChange}
      />
    </div>
  );
};

export default Search;
