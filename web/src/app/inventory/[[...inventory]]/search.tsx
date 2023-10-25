import React from "react";
import styles from "./Search.module.css";
import { search } from "lucide-react";

const Search = ({ value, onChange }) => {
  return (
    <div className={styles.search}>
      <search  className={styles.icon} />
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