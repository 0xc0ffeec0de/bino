#include <stdio.h>

const char* enc_data="\x00\x00\x00\x00\x00\x00\x17\x15\x1a\x06\x1c\x06";
#define enc_data_size 12
const char* enc_key="\x73\x74\x72\x69\x6e\x67";
#define enc_key_size 6


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
