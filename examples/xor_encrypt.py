# Generate c arrays with encrypted xor data

import sys


def to_c_bytearray(s, var_name):
    out = f'const char* {var_name}="'
    size = f'#define {var_name}_size {len(s)}'

    for b in s:
        out += f"\\x0{hex(b)[2:]}" if b < 10 else f"\\{hex(b)[1:]}"
    out += '"'

    code = f"{out};\n{size}"
    return code


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} string key")
        sys.exit(1)

    s = sys.argv[1]
    k = sys.argv[2]
    size_k = len(k)
    out = []

    for i, c in enumerate(s):
        out.append((ord(c) ^ ord(k[i % size_k])))

    enc_code = to_c_bytearray(out, "enc_data")
    key_code = to_c_bytearray([ord(c) for c in k], "enc_key")

    print(f"{enc_code}\n{key_code}")
