#include <stdio.h>

const char* enc_data="\x00\x00\x00\x00\x18\x11\xb\xa\x01\x08";
#define enc_data_size 10
const char* enc_key="\x63\x6f\x6f\x6c\x6b\x65\x79";
#define enc_key_size 7

void decrypt() {
  char output[enc_data_size] = { 0 };

  for (int i = 0; i < enc_data_size; ++i) {
    output[i] = enc_data[i] ^ enc_key[i % enc_key_size];
  }

  // print the call arguments
  puts(output);
}


void print_something() {
  puts("Something nice");
}

int main() {

  decrypt();

  int x = 100;
  int y = x + y;
  
  print_something();

  return 0;
}
