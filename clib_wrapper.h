#ifndef _C_LIB_WRAPPER_H_
#define _C_LIB_WRAPPER_H_

#include "clib.h"

void hello_wrapper(void *userdata, int a, int b, void *f);
int hello_cb_wrapper(void *userdata, int a);

#endif
