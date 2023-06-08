
#include <stdio.h> 
#include <stdlib.h> 

#define SIZE 8

typedef struct item_type {
	int value;
} ItemType;

// Binary Tree
//		Every Node Have At The Maximum 02 Childern 
//		[ Childerns Count Can Be 0, 1, 2 ]

// Binary Search Tree Is Binary Tree With Following Condition
//		Left Child Is Less Than Parent and Right Child Is Greater Than Parent
//			Or Vice-Versa
typedef struct node_type {
	ItemType item;
	struct node_type *left, *right;
} NodeType;

//_______________________________________________________________________

NodeType * createNode( ItemType item ) {
	NodeType * newNode = ( NodeType * ) malloc( sizeof( NodeType ) );

	newNode -> item 	= item;
	newNode -> left 	= NULL;
	newNode -> right 	= NULL;	
	return newNode;
}

//_______________________________________________________________________

NodeType * addNode( NodeType *root, NodeType * node ) {
	if ( root == NULL ) {
		root = node;
		return root;
	}

	if ( node -> item.value < root -> item.value ) {
		root -> left = addNode( root -> left, node ) ;
	} else if ( node -> item.value > root -> item.value ) {
		root -> right = addNode( root -> left, node ) ;
	} else {
		return NULL;
	}

	return node;
}

//_______________________________________________________________________

NodeType * createBinarySearchTree( NodeType * root ) {
	NodeType * newNode = NULL;

	ItemType items[SIZE] = { { 30 }, { 20 } , { 40 }, { 70 } , { 60 } , { 80 }  };

	for (int i = 0 ; i < SIZE ; i++ ) {
		newNode = createNode( items[i] );
		addNode( root, newNode );
	}

	return root;
}

//_______________________________________________________________________

void processNode( NodeType *node ) {
	printf("\nNode Value :: %d", node -> item.value );
}

void inorderTraversal( NodeType * root ) {
	if ( root != NULL ) {
		inorderTraversal( root -> left );
		processNode( root );
		inorderTraversal( root -> right );
	}
}

//_______________________________________________________________________

void testCodeDesign1() {
	NodeType * rootNode = NULL;
	NodeType *newNode = NULL ;
	NodeType *newNode1 = NULL, *newNode2 = NULL, *newNode3 = NULL ;

	ItemType item1 = { 88 };
	ItemType item2 = { 77 };
	ItemType item3 = { 99 };

	newNode1 = createNode( item1 );
	printf("\nNew Node: %d", newNode1 -> item.value );

	newNode2 = createNode( item2 );
	printf("\nNew Node: %d", newNode2 -> item.value );

	newNode3 = createNode( item3 );
	printf("\nNew Node: %d", newNode3 -> item.value );

	ItemType items1[SIZE] = { {10}, {22}, {30}, {20} , {40}, {70} , {60} , {80}  };

	for (int i = 0 ; i < SIZE ; i++ ) {
		newNode = createNode( items1[i] );
		printf("\nNew Node: %d", newNode -> item.value );
	}

	printf("\nBinary Search Tree Inorder Traversal: \n");
	inorderTraversal( rootNode );
}

void testCodeDesign2() {
	NodeType * rootNode = NULL;
	NodeType *newNode = NULL ;

	int nodeCount = 1;
	// ItemType items[SIZE] = { {20} } ; 
	
	nodeCount = 3;
	ItemType items[SIZE] = { {20}, {10}, {30} } ; 
	// ItemType items[SIZE] = { {10}, {20}, {30}, {9}, {80}, {15}, {14}, {16} } ;
	
	rootNode = createNode( items[0] );
	printf("\nRoot Node: %d", rootNode -> item.value );
	
	printf("\nAdding Nodes To Binary Search Tree: \n");
	for (int i = 1 ; i < nodeCount ; i++ ) {
		newNode = createNode( items[i] );
		printf("\nNew Node: %d", newNode -> item.value );
		addNode( rootNode, newNode );
	}
	
	printf("\nBinary Search Tree Inorder Traversal: \n");

	inorderTraversal( rootNode );
	// addNode( root, newNode2 );
}

//_______________________________________________________________________


void main() {
	// testCodeDesign1();
	testCodeDesign2();
	// root = createBinarySearchTree( root );
	// inorderTraversal( root );
}

//_______________________________________________________________________
