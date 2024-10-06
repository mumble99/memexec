#define _GNU_SOURCE
#include <stdio.h>
#include <unistd.h>
#include <sys/mman.h>

// EXAMPLE: gcc memexec-c.c && cat /bin/id | ./a.out

int
main(int argc, char* argv[])
{
    int fd;
    char buffer;
    ssize_t rc;

    fd = memfd_create("", 0);
    if (fd < 0)
    {
        printf("memfd_create failed\n");
        return 1;
    }

    while ((rc = read(0, &buffer, 1)) > 0)
    {
        write(fd, &buffer, rc);
    }

    execveat(fd, "", argv, NULL, 0x1000);
}