def read_file(filename):
    """Read a file from the files directory and return lines as a list."""
    with open(f"../go/files/{filename}", 'r') as f:
        return [line.strip() for line in f.readlines()]
