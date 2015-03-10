#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define NDEBUG
#include <assert.h>

#define STACK_SIZE 64

#define stack_push(zz_s, zz_sp, zz_n, zz_d) \
	assert(zz_sp >= &zz_s[0]); \
	assert(zz_sp < &zz_s[STACK_SIZE - 1]); \
	zz_sp->n = zz_n; \
	zz_sp->d = zz_d; \
	zz_sp++;

#define stack_pop(zz_s, zz_sp, zz_n, zz_d) \
	assert(zz_sp > &zz_s[0]); \
	assert(zz_sp <= &zz_s[STACK_SIZE - 1]); \
	zz_sp--; \
	zz_n = zz_sp->n; \
	zz_d = zz_sp->d;

typedef struct node node_t;
typedef struct pair pair_t;

struct node
{
	int x;
	int r;
	struct node *l[2];
};

struct pair
{
	struct node *n;
	int d;
};

int is_red(node_t *root)
{
	return ((root != NULL) && (root->r == 1));
}

int is_valid(node_t *root)
{
	int lh, rh;
	if (root == NULL)
	{
		return 1;
	}
	else
	{
		node_t *ln = root->l[0];
		node_t *rn = root->l[1];
		/* Consecutive red links */
		if (is_red(root))
		{
			if (is_red(ln) || is_red(rn))
			{
				puts("Red violation");
				return 0;
			}
		}
		lh = is_valid(ln);
		rh = is_valid(rn);
		/* Invalid binary search tree */
		if (
		    (ln != NULL && ln->x >= root->x)
		    || (rn != NULL && rn->x <= root->x))
		{
			puts("Binary tree violation");
			return 0;
		}
		/* Black height mismatch */
		if (lh != 0 && rh != 0 && lh != rh)
		{
			puts("Black violation");
			return 0;
		}
		/* Only count black links */
		if (lh != 0 && rh != 0)
		{
			return is_red(root) ? lh : lh + 1;
		}
		else
		{
			return 0;
		}
	}
}

node_t *new_node(int x, int r)
{
	node_t *n;

	n = (node_t *)malloc(sizeof(node_t));
	if (n == NULL)
	{
		return NULL;
	}
	n->x = x;
	n->r = r;
	n->l[0] = NULL;
	n->l[1] = NULL;
	return n;
}

node_t *rotate(node_t *t, int d)
{
	node_t *n, *z;
	int f;

	f = !d;
	n = t->l[d];
	z = n->l[f];
	n->l[f] = t;
	t->l[d] = z;

	return n;
}

node_t *insert(node_t *r, int x)
{
	int pd, gd;
	pair_t s[STACK_SIZE];
	pair_t *sp;
	node_t *n, *p, *u, *g;

	assert(r != NULL);
	sp = &s[0];
	n = r;
	while (n != NULL)
	{
		if (n->x > x)
		{
			stack_push(s, sp, n, 0);
			n = n->l[0];
		}
		else if (n->x < x)
		{
			stack_push(s, sp, n, 1);
			n = n->l[1];
		}
		else
		{
			return r;
		}
	}
	assert(n == NULL);
	n = new_node(x, 1);
	stack_pop(s, sp, p, pd);
	p->l[pd] = n;
	while (p->r == 1)
	{
		stack_pop(s, sp, g, gd);
		u = g->l[!gd];
		if ((u != NULL) && (u->r == 1))
		{
			u->r = 0;
			p->r = 0;
			if (g != r)
			{
				g->r = 1;
				n = g;
				stack_pop(s, sp, p, pd);
				continue;
			}
			break;
		}
		/* Rotations */
		if (gd != pd)
		{
			p = rotate(p, pd);
			g->l[gd] = p;
			//pd = !pd;
			//n = p->l[pd];
		}
		if (g == r)
		{
			g->r = 1;
			p->r = 0;
			r = rotate(g, gd);
			assert(r == p);
		}
		else
		{
			g->r = 1;
			p->r = 0;
			g = rotate(g, gd);
			assert(g == p);
			n = g;
			stack_pop(s, sp, g, gd);
			g->l[gd] = n;
		}
		break;
	}
	assert(is_valid(r));
	return r;
}

int main()
{
	int i, u;
	node_t *r;
	clock_t estimate;

	estimate = clock();
	r = new_node(10, 0);
	for (i = 0; i < 20000000; i++)
	{
		//r = insert(r, rand());
		r = insert(r, i);
	}
	assert(is_valid(r));
	estimate = clock() - estimate;
	printf("Time taken: %.2fs\n", ((double)estimate) / CLOCKS_PER_SEC);
	return 0;
}
