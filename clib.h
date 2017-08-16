#ifndef _C_LIB_H_
#define _C_LIB_H_

typedef int (*hello_cb)(void *userdata, int c);
void hello(void *userdata, int a, int b, hello_cb cb);

#endif
