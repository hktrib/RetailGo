import torch
import json

image_processor = AutoProcessor.from_pretrained("openai/clip-vit-base-patch32")

def input_fn(request_body, request_content_type, context):
    # The request should be a JSON
    if request_content_type != "application/json":
        raise Exception("Expect Request Content Type application/json, received", request_content_type)
    
    else:
        items = json.loads(request_body)
        # Format items as fetched and preprocessed image and tokenized title
        items  = list(
            map(
                lambda item: {
                # inputs = image_processor(images = Image.open(requests.get("http://ecx.images-amazon.com/images/I/51fAmVkTbyL._SY300_.jpg", stream = True).raw), return_tensors="pt")
                    "image": image_processor(images = Image.open(requests.get(item["imageURL"]), stream = True).raw, return_tensors = "pt")
                },
                items
            )
        )


def predict_fn(input_object, model, context):
    pass

def output_fn(prediction, response_content_type, context):
    pass
