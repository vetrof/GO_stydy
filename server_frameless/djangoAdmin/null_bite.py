import os

def find_null_bytes(directory):
    for root, _, files in os.walk(directory):
        for file in files:
            filepath = os.path.join(root, file)
            try:
                with open(filepath, 'rb') as f:
                    content = f.read()
                    if b'\x00' in content:
                        print(f"Null byte found in: {filepath}")
            except Exception as e:
                print(f"Could not read file {filepath}: {e}")

find_null_bytes('.')
