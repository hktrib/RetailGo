from faker_music import MusicProvider
from faker_food import FoodProvider
import requests
from faker import Faker
import re

#initialize faker the random data class
fake = Faker()
fake.add_provider(MusicProvider)
fake.add_provider(FoodProvider)

#Urls for local and production
url = 'http://localhost:8080'
#url = 'https://retailgo-production.up.railway.app'

def is_valid_email(email):
    email_regex = r'^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$'
    return re.match(email_regex, email) is not None

# Function to create a new store
def create_store(email):
    data = {
        "store_name": "Global Food Co.",
        "store_phone": "9162762606",
        "store_address": "123 Main St, Sacramento, CA 95814",
        "store_type": "Grocery Store",
        "owner_email": email
    }
    response = requests.post(f'{url}/create/store', json=data)
    assert response.status_code in [200, 201, 202]
    response_body = response.json()
    return response_body
    


def create_item(storeId, item_name,CategoryName):
    data = {
        "Name": item_name,
        "Price": fake.random_int(1, 50),
        "Quantity": fake.random_int(20, 1000),
        "category_name": CategoryName,
        "Photo": "https://picsum.photos/200",
    }
    print("Creating item: {}".format(data))
    print("Sending to url: {}/store/{}/inventory/create".format(url,storeId))
    response = requests.post('{}/store/{}/inventory/create'.format(url,storeId), json=data)
    assert response.status_code == 200 or response.status_code == 201 or response.status_code == 202
    print("Item \'{}\'created successfully!".format(item_name))

#Faker genreation methods
def generate_fruit(count):
    #Dont want to add any duplicate values and an easy way to check is just to add them to a list
    added_fruit = []
    for i in range(count):
        fake_item = fake.fruit()
        added_fruit.append(fake_item)
        #create item on server
        create_item(storeId,fake_item, "Fruit")

def generate_veg(count):
    added_veggies = []
    for i in range(count):
        fake_item = fake.vegetable()
        added_veggies.append(fake_item)
        #create item on server
        create_item(storeId,fake_item, "Vegetable")

def generate_sushi(count):
    added_sushi = []
    for i in range(count):
        fake_item = fake.sushi()
        added_sushi.append(fake_item)
        create_item(storeId,fake_item, "Sushi")

def generate_spices(count):
    added_spice = []
    for i in range(count):
        fake_item = fake.spice()
        added_spice.append(fake_item)
        create_item(storeId,fake_item, "Spices")

def main_menu():
    print("Welcome to the Fake Data Generator for RetailGo!")
    print("1. Generate Fruits")
    print("2. Generate Vegetables")
    print("3. Generate Sushi")
    print("4. Generate Spices")
    print("5. Exit")
    choice = input("Enter your choice (1-5): ")
    return choice

def get_count():
    count = int(input("How many items do you want to create? "))
    return count

def main():
    print("Welcome to the Fake Food Data Generator for RetailGo!")
    store_choice = input("Do you want to create a new store or use an existing store ID? (create/use): ")
    global storeId

    if store_choice.lower() == 'create':
        while True:
            email = input("Enter your email to create a store: ")
            if is_valid_email(email):
                storeId = create_store(email)
                print(f"New store created with ID: {storeId}")
                break
            else:
                print("Invalid email format. Please try again.")
    elif store_choice.lower() == 'use':
        storeId = input("Enter your existing store ID: ")
    else:
        print("Invalid choice, exiting.")
        return

    while True:
        choice = main_menu()
        if choice == '1':
            count = get_count()
            generate_fruit(count)
        elif choice == '2':
            count = get_count()
            generate_veg(count)
        elif choice == '3':
            count = get_count()
            generate_sushi(count)
        elif choice == '4':
            count = get_count()
            generate_spices(count)
        elif choice == '5':
            print("Exiting the program.")
            break
        else:
            print("Invalid choice, please try again.")

if __name__ == "__main__":
    main()