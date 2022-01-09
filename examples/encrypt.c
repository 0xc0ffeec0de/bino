#include <stdio.h>


#define enc_data_size 9
#define enc_key_size 4

const char* enc_data="\x10\x00\x02\x09\x17\x07\x06\x02\x04";
const char* enc_key="\x63\x6f\x6f\x6c";



void decrypt() {
  char output[enc_data_size] = { 0 };

  for (int i = 0; i < enc_data_size; ++i) {
    output[i] = enc_data[i] ^ enc_key[i % enc_key_size];
  }

  // print the call arguments
  puts(output);
}

int main() {

  decrypt();

  return 0;
}
