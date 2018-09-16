#include <stdio.h>

main()
{

int sum, i;

sum = 0;

for (i=0; i<1000; i=i+1){
    if((i % 3 == 0) || (i % 5 == 0) ){
      sum = sum + i;
    }
}

printf("%d\n", sum);

}
