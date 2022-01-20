# TREE (DEPRECATED)

The main idea was to use go to manipulate the DOM using wasm. While posible, is not efficient at all
since wasm has no direct access to the DOM instead it works as a messenger that tells it what
actions to perform.

For Example:

creating 1000 node element + appending to a container took about 70ms on avg while plain JS took 8ms

However, algorithmic computation is the real benefit of wasm, computing a 30 digit fibonacci sequence took
about 38ms on JS but the same algorithm took 22ms when executed in wasm.

Go library to access and manipulate DOM Elements using wasm

# Testing

To run test suite follow these steps:

-   go install github.com/agnivade/wasmbrowsertest@latest
-   chmod +x test.sh
-   ./test.sh
