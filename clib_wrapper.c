#include "clib_wrapper.h"

extern int hello_cb_go(void *userdata, int a);

void hello_wrapper(void *userdata, int a, int b, void *f) {
	hello(userdata, a, b, (hello_cb)f);
}

int hello_cb_wrapper(void *userdata, int a) {
	int result = hello_cb_go(userdata, a);
	return result;
}

