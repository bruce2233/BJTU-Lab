#include <stdio.h>
#include <conio.h>
#include <stdlib.h>
#include <string.h>

static int blockaddress[20005]; //存储了所需数据的块地址
static int before[500];         //存放了从cache中被替换出的数据块的tag位，用于判断失效类型。
static int t = 0;

int misstype(int ba, int nb, int l);
int mycache();

int main()
{
    int abc;
    abc = mycache();
    FILE *fp1;
    int cachesize; 
    int blocksize;
    int assoc;
    int blockinbyte;
    int NOofblock;
    int NOofset;
    int choice;
    float misscount, accesscount, hitcount;
    int index, byte, tag, ii;
    int i = 0, j, x, y, z, cc, c, m;
    int bytearray[20005], wordaddress[20005]; //代表所需数据的字节地址与字地址

    // newarray中存储cache中各块的valid位与tag位；lru中存储最近使用情况，用来实现LRU替换策略（对应位数越大，代表越久未被使用）
    int newarray[300][300] = {0}, lru[300][300] = {0};

    char ans = 'y';
    int c1c = 0, c2c = 0, c3c = 0;
    float missrate = 0, hitrate = 0;

    do
    {
        printf("Cache Simulation Project:");
        printf("\n\n1. Direct_mapped\n2. Set_associate\n3. Fully_associate\n\n: ");
        // choice from direct mapped, set associate and fully associate
        scanf("%d", &choice);
        if (choice == 1 || choice == 2 || choice == 3)
            break;
        printf("Incorrect input.");

    } while (1);

    do
    {
        printf("\n\nCache Size from range[64/128/256]: ");
        scanf("%d", &cachesize);
        if (cachesize == 64 || cachesize == 128 || cachesize == 256 || cachesize == 16) // choose cache size
            break;
        printf("Incorrect input.");
    } while (1);

    do
    {
        printf("\n\nBlock Size from range[1/2/4]: ");
        scanf("%d", &blocksize);
        if (blocksize == 1 || blocksize == 2 || blocksize == 4) // choose block size
            break;
        printf("Incorrect input.");
    } while (1);

    do
    {
        printf("\n\nEnter the value for n-way Set value from[1/2/4/8/16]: ");
        scanf("%d", &assoc);
        if (assoc == 1 || assoc == 2 || assoc == 4 || assoc == 8 || assoc == 16)
            break;
        printf("Incorrect input.\n");
    } while (1);
    printf("123\n");
    for (ii = 0; ii < 500; ii++)
        before[ii] = -1;
    blockinbyte = blocksize * 4;         //以字节计算出的块体积calculate the block size in bit
    NOofblock = cachesize / blockinbyte; // calculate the # of blocks
    NOofset = NOofblock / assoc;
    // calculate the # of sets
    fp1 = fopen("project.txt", "r"); // open the file
    printf("123\n");
    while (fscanf(fp1, "%d", &byte) != EOF)
    {
        bytearray[i] = byte;
        i++;
    }
    printf("i is %d\n", i);
    fclose(fp1);

    misscount = 0; // initial miss counter=0
    hitcount = 0;  // initial hit counter=0
    accesscount = 0;
    // initial access counter=0
    for (j = 0; j < i; j++)
    {
        accesscount++;                                // increase the access counter
        wordaddress[j] = bytearray[j] / 4;            //计算数据的字地址            //calculate the word address
        blockaddress[j] = wordaddress[j] / blocksize; //计算数据的块地址 //calculate the block address
        index = blockaddress[j] % NOofset;            //计算数据所在的set地址            //calculate the index
        tag = blockaddress[j] / NOofset;              //计算数据所在的tag            //calculate the tag
        x = y = z = 0;
        if (choice == 1 || choice == 2) // 1、当为直接映射或集相联映射时
        //此时，映射出的cache中地址（index）为固定值，则循环看"该映射行"中是否有匹配的数据块即可，这通过访问newarray[index][]来实现。
        //又因该cache表中每一个cache块都需要保存valid与tag位，则每块需要占用两位的newarray（newarray[index][z]中存的是valid，
        // newarray[index][z+1]中存的是tag），所以循环的增量条件是z = z + 2；
        //当是n路组相联时，代表"该映射行"中包含n个（assoc个）cache块，则共占用2n位newarray，则循环结束的条件是z<(assoc*2) （数组地址从0开始，则小于即可）
        {
            while (z < (assoc * 2)) //循环结束条件为assoc*2，这是因为每一个cache块都需要保存valid与tag位，则每块需要占用两位的newarray
            {
                cc = 0;
                c = 0;
                if (newarray[index][z] == 0) // 1.1、当cache 中对应位置上的有效位无效，数据不在cache中，失效
                //此时能直接判断失效的原因：在多路组相联中，是从左到右依次存数据的，则遇到的第一个valid为0的数据时，后面就不可能再有数据了。
                //此时，能够确定是失效，则将该块存入该cache块中（填写tag与valid位）、判断失效类型、更新LRU数组。
                {
                    newarray[index][z + 1] = tag; //填写tag位
                    newarray[index][z] = 1;       //填写有效位为1
                    misscount++;
                    c = misstype(blockaddress[j], NOofblock, j); //判断失效类型
                    cc = 1;                                      //该cc值其实没用

                    //以下为替换策略（近期最少使用策略LRU）的实现方案，主要通过lru[index][]数组实现（代表该块多久未使用，越大代表近期越少使用）。
                    //先将index这一行的所有块的lru[index][m]都加1，代表此次未被访问，然后把此次用到的这一个块（z块）的lru[index][z]至0，代表刚刚被访问。
                    for (m = 0; m < (assoc * 2); m = m + 2)
                        lru[index][m]++; // increase the value of lru[index][m] to 1 in the same index
                    lru[index][z] = 0;   // set the  recent use value to 0

                    z = (assoc * 2); //退出循环
                }
                else // 1.2、当cache 中对应位置上的有效位有效，继续判断tag位是否一致
                {
                    if (newarray[index][z + 1] == tag) // 1.2.1、并且tag位一致，则数据在cache中，命中
                    {
                        hitcount++; // hit counter increase

                        //同上面的原因，为了实现LRU的替换策略，更新lru[index][]数组
                        for (m = 0; m < (assoc * 2); m = m + 2)
                            lru[index][m]++;
                        lru[index][z] = 0;

                        z = (assoc * 2); //退出循环
                    }
                    else // 1.2.2、但是tag位不一致
                    {
                        if (assoc < 2) // 1.2.2.1、直接映射时，则肯定是失效了（不存在其他可放的块了）
                        {
                            newarray[index][z + 1] = tag; //直接替换（替换为当前所需要的数据块）
                            misscount++;
                            c = misstype(blockaddress[j], NOofblock, j); // decide which miss type this miss belong to
                            cc = 1;
                            z = (assoc * 2); //退出循环
                        }
                        else // 1.2.2.2、组相联映射时，则可能还有其他块（其他路），循环继续检测
                        {
                            if (x < lru[index][z]) //该if语句块，保证了y中存放的是当前index这一行中最久未被使用的块的路数（组数），也就是当全部块都被放满，将要被替换的那个块。
                            {
                                x = lru[index][z];
                                y = z;
                            }

                            if (z == ((assoc * 2) - 2)) //如果该index中的所有块都已经循环检测过，均无所需的块，则失效，要进行LRU替换，直接替换到y的位置即可。
                            {
                                newarray[index][y + 1] = tag; // y处即为近期最久未被使用的块，因为其lru[index][]的值最大
                                misscount++;
                                c = misstype(blockaddress[j], NOofblock, j);
                                cc = 1;

                                for (m = 0; m < (assoc * 2); m = m + 2) //更新lru[index][]数组
                                    lru[index][m]++;
                                lru[index][y] = 0;
                            }
                            z = z + 2; //继续循环，去检测该index的下一组数据（下一个块）
                        }
                    }
                }
            }

            if (c == 1 && cc == 1) //此处cc其实无用
                c1c++;
            if (c == 2 && cc == 1)
                c2c++;
            if (c == 3 && cc == 1)
                c3c++;
        }
        else // 2、当为全相联映射时
        {
            while (z <= NOofblock) //由于是全相联映射，则所有的从上到下所有的cache块均可存该数据，则从上到下循环即可。
            {
                cc = 0;
                if (newarray[z][0] == 0) // 2.1、如果该cache的valid位为0，则为失效。因为：数据是从上往下扫描的，所以遇到的第一个valid为0的，下面肯定不会有数据的。
                {
                    //由于失效，将该数据块更新到cache中
                    newarray[z][1] = blockaddress[j]; // staore the block address in the new array
                    newarray[z][0] = 1;               // valid bit set to 1
                    misscount++;
                    c = misstype(blockaddress[j], NOofblock, j);
                    cc = 1;
                    for (m = 0; m <= NOofblock; m++) //更新lru[][]的值，这样当所有cache都放满时，会利用此数组的值进行LRU替换
                        lru[m][1]++;
                    lru[z][1] = 0;
                    z = (NOofblock + 1); //退出循环
                }
                else // 2.2、如果遇到某块的valid位为1
                {
                    if (newarray[z][1] == blockaddress[j]) // 2.2.1、tag位也相符，证明命中。
                    {
                        hitcount++;
                        for (m = 0; m <= NOofblock; m++) //更新lru数组
                            lru[m][1]++;
                        lru[z][1] = 0;
                        z = NOofblock + 1; //退出循环
                    }
                    else // 2.2.2、tag位不相符相符，则更新lru数组后继续循环（如果没到cache的最后）。
                    {
                        if (x < lru[z][1]) //该if语句块，保证了y中存放的是当前index这一行中最久未被使用的块的路数（组数），也就是当全部块都被放满，将要被替换的那个块。
                        {
                            x = lru[z][1];
                            y = z;
                        }
                        if (z == NOofblock) //如果该cache中的所有块都已经循环检测过，均无所需的块，则失效，要进行LRU替换，直接替换到y的位置即可。
                        {
                            newarray[y][1] = blockaddress[j]; // y处即为近期最久未被使用的块，因为其lru[][1]的值最大
                            misscount++;
                            c = misstype(blockaddress[j], NOofblock, j);
                            cc = 1;
                            for (m = 0; m <= NOofblock; m++) //更新lru数组
                                lru[m][1]++;
                            lru[y][1] = 0;
                        }
                        z++; //继续检查下一cache块
                    }
                }
            }
            if (c == 1 && cc == 1)
                c1c++;
            if (c == 2 && cc == 1)
                c2c++;
            if (c == 3 && cc == 1)
                c3c++;
        }
    }

    missrate = (misscount / accesscount);
    hitrate = (hitcount / accesscount);
    printf("\n        Miss Rate = %3f \n", missrate);
    printf("\n         Hit Rate = %3f \n", hitrate);
    printf("\n  Compulsory Miss = %3d \n", c1c);
    printf("\n    Capacity Miss = %3d \n", c3c);
    printf("\n    Conflict Miss = %3d \n", c2c);
    printf("\n       hit Number = %3f \n", hitcount);
    printf("\n      miss Number = %3f \n", misscount);
    printf("\n    Access Number = %3f \n", accesscount);
}

int mycache()
{
    FILE *fp = fopen("project.txt", "w");
    int i, j, c, stride, array[256];
    int abc = 0;
    stride = 131;
    for (i = 0; i < 10000; i++)
    {
        for (j = 0; j < 256; j = j + stride)
        {
            // c = array[j] + 5;
            fprintf(fp, "%d ", j * 4); //字地址转换成字节地址
            abc++;
        }
    }
    fclose(fp);
    return abc;
}

int misstype(int ba, int nb, int l) // this function is used to decide which miss type it's belong to
{
    int u, k = 0, b = 0, n = 0, m, ii;
    int blarray[500];
    int type;
    for (ii = 0; ii < 500; ii++) // initila the array
        blarray[ii] = 9999;
    for (u = 0; u <= t; u++) // check if the block address already in the array
    {
        if (before[u] == ba) // if the block address in the before array
        {
            k = 0;
            break;
        }
        else
            k = 1;
    }
    if (k == 1) // compulsory miss
    {
        type = 1;
        before[t] = ba;
        t++;
    }
    if (k == 0) // capacity or conflict miss
    {
        for (u = (l - 1); u >= 0; u--)
        {
            if (blockaddress[u] == ba) //检查所有已经访问过的数据blockaddress[u]，看当前访问数据是否已经被访问过。例如以前访问过202，现在有访问202.
            {
                break;
            }
            else
            {
                n = 0;
                for (m = 0; m <= b; m++)
                {
                    n++;
                    if (blarray[m] == blockaddress[u]) // it it's not a distinct address, then break
                        break;
                }

                if (n == (b + 1))
                {
                    blarray[b] = blockaddress[u]; // store the address in the blarray
                    b++;                          // blarray[b+1] to store next address
                }
            }
        }
        if ((b) < nb) // conflict miss
            type = 2;
        else // capacity miss
            type = 3;
    }
    return type;
}
