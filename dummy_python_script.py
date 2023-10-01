def add_numbers(a, b):
    return a + b

def subtract_numbers(a, b):
    return a - b

def multiply_numbers(a, b):
    return a * b

def divide_numbers(a, b):
    if b == 0:
        raise ValueError("Division by zero is not allowed")
    return a / b

if __name__ == "__main__":
    result = add_numbers(5, 3)
    print(f"Addition result: {result}")
    
    result = subtract_numbers(8, 2)
    print(f"Subtraction result: {result}")
    
    result = multiply_numbers(4, 7)
    print(f"Multiplication result: {result}")
    
    try:
        result = divide_numbers(6, 0)
    except ValueError as e:
        print(f"Error: {e}")
