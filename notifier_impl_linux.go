//+build linux,cgo android,cgo

package networkchangenotifier

/*
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <memory.h>
#include <unistd.h>
#include <errno.h>
#include <net/if.h>
#include <netinet/in.h>
#include <linux/netlink.h>
#include <linux/rtnetlink.h>
#include <time.h>
#include <pthread.h>

extern void callback_cgo(uint64_t dataNLMSG);

typedef void (*callback_t)(uint64_t);

callback_t callback;

pthread_mutex_t lock;

void init(){
   pthread_mutex_init(&lock, NULL);
}

void cleanup(){
   pthread_mutex_destroy(&lock);
}

int sock = 0;

bool readevent = false;

int registerNetworkChangeEvent()
{
   sock = socket(AF_NETLINK,SOCK_RAW,NETLINK_ROUTE);
   struct sockaddr_nl addr;
   memset((void *)&addr, 0, sizeof(addr));
   if (sock<0) {
      return errno;
   }
   addr.nl_family = AF_NETLINK;
   addr.nl_pid = getpid();
   // addr.nl_groups = RTMGRP_LINK|RTMGRP_IPV4_IFADDR|RTMGRP_IPV6_IFADDR|RTMGRP_IPV4_ROUTE|RTMGRP_IPV6_ROUTE;
   addr.nl_groups = RTMGRP_IPV4_ROUTE|RTMGRP_IPV6_ROUTE;

   struct timeval timeout;
   timeout.tv_sec = 2;
   timeout.tv_usec = 0;
   setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO, (char*)&timeout, sizeof(timeout));

   int res = bind(sock, (struct sockaddr *)&addr, sizeof(addr));
   if (res<0) {
      int err = errno;
      close(sock);
      return errno;
   }
   return 0;
}

int read_event(int (*msg_handler)(struct sockaddr_nl *, struct nlmsghdr *))
{
    int ret = 0;
    char buf[4096];
    struct iovec iov = {
        .iov_base = buf,
        .iov_len = sizeof(buf)
    };
    struct sockaddr_nl snl;
    struct msghdr msg = {
        .msg_name =(void*)&snl,
        .msg_namelen=sizeof(snl),
        .msg_iov = &iov,
        .msg_iovlen = 1,
        .msg_control = NULL,
        .msg_controllen = 0,
        .msg_flags = 0
    };
    struct nlmsghdr *h;
    int status = recvmsg(sock, &msg, 0);
    if(status < 0)
    {
     return errno;
    }
    for(h = (struct nlmsghdr *)buf; NLMSG_OK(h, status); h = NLMSG_NEXT (h, status))
    {
      if (h->nlmsg_type == NLMSG_DONE)
      {
         return ret;
      }
      if (h->nlmsg_type == NLMSG_ERROR)
      {
         return errno;
      }
      status = (*msg_handler)(&snl, h);
      if(status != 0)
      {
         return status;
      }
    }
    return ret;
}

int msg_handler(struct sockaddr_nl *nl, struct nlmsghdr *msg)
{
   int type = msg->nlmsg_type;
   char* data = NLMSG_DATA(msg);
   pthread_mutex_lock(&lock);
   if (callback != NULL) {
      callback((uint64_t)data);
   }
   pthread_mutex_unlock(&lock);
   return 0;
}

void regCallback(callback_t cb){
   int ret = 0;
   pthread_mutex_lock(&lock);
   callback = cb;
   readevent = true;
   pthread_mutex_unlock(&lock);

   while(readevent){
      ret = read_event(msg_handler);
      if (ret != 0 && ret != EAGAIN ) {
         break;
      }
   }
}

void unregCallback(){
   pthread_mutex_lock(&lock);
   readevent = false;
   callback = NULL;
   pthread_mutex_unlock(&lock);
}

int unregisterNetworkChangeEvent()
{
   int err = 0;
   unregCallback();
   if (close(sock) != 0)
   {
      err = errno;
   }
   sock = 0;
   sleep(2);
   return err;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func ncnInit() error {
	var err error
	C.init()
	ret := int32(C.registerNetworkChangeEvent())
	if ret != 0 {
		return fmt.Errorf("registerNetworkChangeEvent failed, err code %d", ret)
	}
	return err
}

func ncnRegisterCallback() {
	go C.regCallback((C.callback_t)(unsafe.Pointer(C.callback_cgo)))
}

func ncnUnregisterCallback() {
	C.unregCallback()
}

func ncnCleanup() error {
	var err error
	defer C.cleanup()
	ret := int32(C.unregisterNetworkChangeEvent())
	if ret != 0 {
		return fmt.Errorf("unregisterNetworkChangeEvent failed, err code %d", ret)
	}
	return err
}
