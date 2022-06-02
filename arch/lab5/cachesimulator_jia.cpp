#include <stdio.h>
#include <conio.h>
#include <stdlib.h>
#include <string.h>

static int blockaddress[20005]; //�洢���������ݵĿ��ַ
static int before[500];         //����˴�cache�б��滻�������ݿ��tagλ�������ж�ʧЧ���͡�
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
    int bytearray[20005], wordaddress[20005]; //�����������ݵ��ֽڵ�ַ���ֵ�ַ

    // newarray�д洢cache�и����validλ��tagλ��lru�д洢���ʹ�����������ʵ��LRU�滻���ԣ���Ӧλ��Խ�󣬴���Խ��δ��ʹ�ã�
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
    blockinbyte = blocksize * 4;         //���ֽڼ�����Ŀ����calculate the block size in bit
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
        wordaddress[j] = bytearray[j] / 4;            //�������ݵ��ֵ�ַ            //calculate the word address
        blockaddress[j] = wordaddress[j] / blocksize; //�������ݵĿ��ַ //calculate the block address
        index = blockaddress[j] % NOofset;            //�����������ڵ�set��ַ            //calculate the index
        tag = blockaddress[j] / NOofset;              //�����������ڵ�tag            //calculate the tag
        x = y = z = 0;
        if (choice == 1 || choice == 2) // 1����Ϊֱ��ӳ�������ӳ��ʱ
        //��ʱ��ӳ�����cache�е�ַ��index��Ϊ�̶�ֵ����ѭ����"��ӳ����"���Ƿ���ƥ������ݿ鼴�ɣ���ͨ������newarray[index][]��ʵ�֡�
        //�����cache����ÿһ��cache�鶼��Ҫ����valid��tagλ����ÿ����Ҫռ����λ��newarray��newarray[index][z]�д����valid��
        // newarray[index][z+1]�д����tag��������ѭ��������������z = z + 2��
        //����n·������ʱ������"��ӳ����"�а���n����assoc����cache�飬��ռ��2nλnewarray����ѭ��������������z<(assoc*2) �������ַ��0��ʼ����С�ڼ��ɣ�
        {
            while (z < (assoc * 2)) //ѭ����������Ϊassoc*2��������Ϊÿһ��cache�鶼��Ҫ����valid��tagλ����ÿ����Ҫռ����λ��newarray
            {
                cc = 0;
                c = 0;
                if (newarray[index][z] == 0) // 1.1����cache �ж�Ӧλ���ϵ���Чλ��Ч�����ݲ���cache�У�ʧЧ
                //��ʱ��ֱ���ж�ʧЧ��ԭ���ڶ�·�������У��Ǵ��������δ����ݵģ��������ĵ�һ��validΪ0������ʱ������Ͳ��������������ˡ�
                //��ʱ���ܹ�ȷ����ʧЧ���򽫸ÿ�����cache���У���дtag��validλ�����ж�ʧЧ���͡�����LRU���顣
                {
                    newarray[index][z + 1] = tag; //��дtagλ
                    newarray[index][z] = 1;       //��д��ЧλΪ1
                    misscount++;
                    c = misstype(blockaddress[j], NOofblock, j); //�ж�ʧЧ����
                    cc = 1;                                      //��ccֵ��ʵû��

                    //����Ϊ�滻���ԣ���������ʹ�ò���LRU����ʵ�ַ�������Ҫͨ��lru[index][]����ʵ�֣�����ÿ���δʹ�ã�Խ��������Խ��ʹ�ã���
                    //�Ƚ�index��һ�е����п��lru[index][m]����1������˴�δ�����ʣ�Ȼ��Ѵ˴��õ�����һ���飨z�飩��lru[index][z]��0������ոձ����ʡ�
                    for (m = 0; m < (assoc * 2); m = m + 2)
                        lru[index][m]++; // increase the value of lru[index][m] to 1 in the same index
                    lru[index][z] = 0;   // set the  recent use value to 0

                    z = (assoc * 2); //�˳�ѭ��
                }
                else // 1.2����cache �ж�Ӧλ���ϵ���Чλ��Ч�������ж�tagλ�Ƿ�һ��
                {
                    if (newarray[index][z + 1] == tag) // 1.2.1������tagλһ�£���������cache�У�����
                    {
                        hitcount++; // hit counter increase

                        //ͬ�����ԭ��Ϊ��ʵ��LRU���滻���ԣ�����lru[index][]����
                        for (m = 0; m < (assoc * 2); m = m + 2)
                            lru[index][m]++;
                        lru[index][z] = 0;

                        z = (assoc * 2); //�˳�ѭ��
                    }
                    else // 1.2.2������tagλ��һ��
                    {
                        if (assoc < 2) // 1.2.2.1��ֱ��ӳ��ʱ����϶���ʧЧ�ˣ������������ɷŵĿ��ˣ�
                        {
                            newarray[index][z + 1] = tag; //ֱ���滻���滻Ϊ��ǰ����Ҫ�����ݿ飩
                            misscount++;
                            c = misstype(blockaddress[j], NOofblock, j); // decide which miss type this miss belong to
                            cc = 1;
                            z = (assoc * 2); //�˳�ѭ��
                        }
                        else // 1.2.2.2��������ӳ��ʱ������ܻ��������飨����·����ѭ���������
                        {
                            if (x < lru[index][z]) //��if���飬��֤��y�д�ŵ��ǵ�ǰindex��һ�������δ��ʹ�õĿ��·������������Ҳ���ǵ�ȫ���鶼����������Ҫ���滻���Ǹ��顣
                            {
                                x = lru[index][z];
                                y = z;
                            }

                            if (z == ((assoc * 2) - 2)) //�����index�е����п鶼�Ѿ�ѭ����������������Ŀ飬��ʧЧ��Ҫ����LRU�滻��ֱ���滻��y��λ�ü��ɡ�
                            {
                                newarray[index][y + 1] = tag; // y����Ϊ�������δ��ʹ�õĿ飬��Ϊ��lru[index][]��ֵ���
                                misscount++;
                                c = misstype(blockaddress[j], NOofblock, j);
                                cc = 1;

                                for (m = 0; m < (assoc * 2); m = m + 2) //����lru[index][]����
                                    lru[index][m]++;
                                lru[index][y] = 0;
                            }
                            z = z + 2; //����ѭ����ȥ����index����һ�����ݣ���һ���飩
                        }
                    }
                }
            }

            if (c == 1 && cc == 1) //�˴�cc��ʵ����
                c1c++;
            if (c == 2 && cc == 1)
                c2c++;
            if (c == 3 && cc == 1)
                c3c++;
        }
        else // 2����Ϊȫ����ӳ��ʱ
        {
            while (z <= NOofblock) //������ȫ����ӳ�䣬�����еĴ��ϵ������е�cache����ɴ�����ݣ�����ϵ���ѭ�����ɡ�
            {
                cc = 0;
                if (newarray[z][0] == 0) // 2.1�������cache��validλΪ0����ΪʧЧ����Ϊ�������Ǵ�������ɨ��ģ����������ĵ�һ��validΪ0�ģ�����϶����������ݵġ�
                {
                    //����ʧЧ���������ݿ���µ�cache��
                    newarray[z][1] = blockaddress[j]; // staore the block address in the new array
                    newarray[z][0] = 1;               // valid bit set to 1
                    misscount++;
                    c = misstype(blockaddress[j], NOofblock, j);
                    cc = 1;
                    for (m = 0; m <= NOofblock; m++) //����lru[][]��ֵ������������cache������ʱ�������ô������ֵ����LRU�滻
                        lru[m][1]++;
                    lru[z][1] = 0;
                    z = (NOofblock + 1); //�˳�ѭ��
                }
                else // 2.2���������ĳ���validλΪ1
                {
                    if (newarray[z][1] == blockaddress[j]) // 2.2.1��tagλҲ�����֤�����С�
                    {
                        hitcount++;
                        for (m = 0; m <= NOofblock; m++) //����lru����
                            lru[m][1]++;
                        lru[z][1] = 0;
                        z = NOofblock + 1; //�˳�ѭ��
                    }
                    else // 2.2.2��tagλ���������������lru��������ѭ�������û��cache����󣩡�
                    {
                        if (x < lru[z][1]) //��if���飬��֤��y�д�ŵ��ǵ�ǰindex��һ�������δ��ʹ�õĿ��·������������Ҳ���ǵ�ȫ���鶼����������Ҫ���滻���Ǹ��顣
                        {
                            x = lru[z][1];
                            y = z;
                        }
                        if (z == NOofblock) //�����cache�е����п鶼�Ѿ�ѭ����������������Ŀ飬��ʧЧ��Ҫ����LRU�滻��ֱ���滻��y��λ�ü��ɡ�
                        {
                            newarray[y][1] = blockaddress[j]; // y����Ϊ�������δ��ʹ�õĿ飬��Ϊ��lru[][1]��ֵ���
                            misscount++;
                            c = misstype(blockaddress[j], NOofblock, j);
                            cc = 1;
                            for (m = 0; m <= NOofblock; m++) //����lru����
                                lru[m][1]++;
                            lru[y][1] = 0;
                        }
                        z++; //���������һcache��
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
            fprintf(fp, "%d ", j * 4); //�ֵ�ַת�����ֽڵ�ַ
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
            if (blockaddress[u] == ba) //��������Ѿ����ʹ�������blockaddress[u]������ǰ���������Ƿ��Ѿ������ʹ���������ǰ���ʹ�202�������з���202.
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
