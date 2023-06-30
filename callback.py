def process_callbacks(callbacks):
    for callback in callbacks:
        callback()

# Example usage
def callback1():
    print("Callback 1")

def callback2():
    print("Callback 2")

# Create a list of 100,000 callbacks
callbacks = [callback1] * 50000 + [callback2] * 50000

# Process the callbacks
process_callbacks(callbacks)