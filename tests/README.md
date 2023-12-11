# Retail Go Tests

## Description
This project is a Python script to generate dummy data for retail go stores. It allows users to create fake inventory items in various categories such as fruits, vegetables, sushi, and spices. Users can choose to create a new store or use an existing store ID to populate the data.

## Installation
1. Clone the repository:
    ```bash
    git clone [your-repository-link]
    ```

2. Navigate to the project directory:
    ```bash
    cd tests
    ```

3. Install the required dependencies:
    ```bash
    pip install -r requirements.txt
    ```

## Usage
    To run the script, execute the following command:
```bash
python retail_go_tests.py
```

    When prompted, choose to create a new store or use an existing store ID.

    If creating a new store, enter a valid email address.
    If using an existing store, provide the store ID.
    Select the type of data you wish to generate (Fruits, Vegetables, Sushi, Spices).

    Enter the number of items you want to generate for the chosen category.

    The script will then create these items and send them to the Retail Go server.

## Sample output:
    Welcome to the Fake Data Generator for RetailGo!
    Do you want to create a new store or use an existing store ID? (create/use): use
    Enter your existing store ID: {store id}
    Welcome to the Fake Data Generator for RetailGo!
    1. Generate Fruits
    2. Generate Vegetables
    3. Generate Sushi
    4. Generate Spices
    5. Exit
    Enter your choice (1-5): 4
    How many items do you want to create? 2
    Creating item: {'Name': 'Fines Herbes', 'Price': 9, 'Quantity': 659, 'category_name': 'Spices', 'Photo': 'https://picsum.photos/200'}