import tensorflow as tf
import numpy as np
import sys
import json
from tensorflow.keras.preprocessing import image

def predict(image_path):
    model = tf.keras.models.load_model('../../../pkg/model/best_plant_disease.h5')
    
    img = image.load_img(image_path, target_size=(150, 150))
    img_array = image.img_to_array(img)
    img_array = np.expand_dims(img_array, axis=0)  # Create batch axis
    
    predictions = model.predict(img_array)
    predicted_class = np.argmax(predictions, axis=1)
    
    return predicted_class.tolist()

if __name__ == "__main__":
    image_path = sys.argv[1]
    predictions = predict(image_path)
    print(json.dumps(predictions))
