/*
* Example Singleton C
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
*/

#include <stdlib.h>
#include <stdio.h>
#include <pthread.h>
#include <unistd.h> 

struct Conn
{   
    int conn_pg;
    int thread;
};

struct Conn* Connect()
{
    static struct Conn *instance = NULL;

    // entra somente
    // uma unica vez
    if(instance == NULL)
    {
        instance = malloc(sizeof(*instance));
        instance->conn_pg = 34498;
    }
    
    return instance;
};

void *PConn(void *tmp){

   struct Conn *v = (struct Conn *) tmp;
   
   // aguarde ...
   sleep(rand()%10)+1;

   printf("Oi eu sou a thread %d\n", v->thread);
   printf("Connect em C: %d\n", Connect()->conn_pg);
}

int main() {

    pthread_t linhas[10];
    int execute,i;

    struct Conn *conn;


    //// thread 1
    printf("Criei a thread 1\n");
    conn = (struct Conn *) malloc(sizeof(struct Conn *));
    conn->thread = 1;
    execute = pthread_create(&linhas[0],NULL,PConn,(void *)conn);


    //// thread 2
    printf("Criei a thread 2\n");
    conn = (struct Conn *) malloc(sizeof(struct Conn *));
    conn->thread = 2;
    execute = pthread_create(&linhas[1],NULL,PConn,(void *)conn);

    // aguarde ...
    sleep(rand()%10)+2;
    pthread_exit(NULL);

    return 0;
}