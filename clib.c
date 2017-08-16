#include "clib.h"

void hello(void *userdata, int a, int b, hello_cb cb) {
	int c = a + b;
	if(cb) {
		cb(userdata, c);
	}
}

