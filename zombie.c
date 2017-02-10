#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main(int argc, char ** argv) {
    pid_t pid = fork();
    if (pid == 0) {
        pid = fork();
        if (pid == 0) {
            exit(0);
        }
        sleep(3);
        exit(0);
    }
    printf("got pid %d and exited\n", pid);
    // sleep(20);
}
