#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define NDEBUG
#include <assert.h>


struct jsw_node
{
    int red;
    int data;
    struct jsw_node *link[2];
};

struct jsw_tree
{
    struct jsw_node *root;
};

int is_red(struct jsw_node *root)
{
    return root != NULL && root->red == 1;
}

struct jsw_node *jsw_single(struct jsw_node *root, int dir)
{
    struct jsw_node *save = root->link[!dir];

    root->link[!dir] = save->link[dir];
    save->link[dir] = root;

    root->red = 1;
    save->red = 0;

    return save;
}

struct jsw_node *jsw_double(struct jsw_node *root, int dir)
{
    root->link[!dir] = jsw_single(root->link[!dir], !dir);

    return jsw_single(root, dir);
}

struct jsw_tree *make_tree()
{
	struct jsw_tree *t = malloc(sizeof *t);
	t->root = NULL;
	return t;
}

struct jsw_node *make_node(int data)
{
    struct jsw_node *rn = malloc(sizeof *rn);

    if (rn != NULL)
    {
        rn->data = data;
        rn->red = 1; /* 1 is red, 0 is black */
        rn->link[0] = NULL;
        rn->link[1] = NULL;
    }

    return rn;
}

int jsw_insert(struct jsw_tree *tree, int data)
{
    if (tree->root == NULL)
    {
        /* Empty tree case */
        tree->root = make_node(data);

        if (tree->root == NULL)
        {
            return 0;
        }
    }
    else
    {
        struct jsw_node head = { 0 }; /* False tree root */

        struct jsw_node *g, *t;     /* Grandparent & parent */
        struct jsw_node *p, *q;     /* Iterator & parent */
        int dir = 0, last;

        /* Set up helpers */
        t = &head;
        g = p = NULL;
        q = t->link[1] = tree->root;

        /* Search down the tree */
        for (;;)
        {
            if (q == NULL)
            {
                /* Insert new node at the bottom */
                p->link[dir] = q = make_node(data);

                if (q == NULL)
                {
                    return 0;
                }
            }
            else if (is_red(q->link[0]) && is_red(q->link[1]))
            {
                /* Color flip */
                q->red = 1;
                q->link[0]->red = 0;
                q->link[1]->red = 0;
            }

            /* Fix red violation */
            if (is_red(q) && is_red(p))
            {
                int dir2 = t->link[1] == g;

                if (q == p->link[last])
                {
                    t->link[dir2] = jsw_single(g, !last);
                }
                else
                {
                    t->link[dir2] = jsw_double(g, !last);
                }
            }

            /* Stop if found */
            if (q->data == data)
            {
                break;
            }

            last = dir;
            dir = q->data < data;

            /* Update helpers */
            if (g != NULL)
            {
                t = g;
            }

            g = p, p = q;
            q = q->link[dir];
        }

        /* Update root */
        tree->root = head.link[1];
    }

    /* Make root black */
    tree->root->red = 0;

    return 1;
}

int main()
{
        int i;
        struct jsw_tree *t;
        clock_t estimate;

        estimate = clock();
	t = make_tree();	
        for (i = 0; i < 20000000; i++)
        {
                jsw_insert(t, i);
        }
        estimate = clock() - estimate;
        printf("Time taken: %.2fs\n", ((double)estimate) / CLOCKS_PER_SEC);
        return 0;
}

