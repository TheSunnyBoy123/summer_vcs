# read the contents of a file line by line
# then rewrite the file in the same way 

if __name__ == '__main__':
    with open('test.mdx', 'r') as f:
        lines = f.readlines()
    with open('test.mdx', 'w') as f:
        for line in lines:
            f.write(line)