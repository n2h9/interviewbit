#include <iostream>
#include <unordered_map> 

struct cacheNode {
  int key;
  int value;
  cacheNode* prev;
  cacheNode* next;

  cacheNode(int key, int value) {
    this->key = key;
    this->value = value;
  }
  ~cacheNode() {
    if (this->prev != NULL) {
      delete this->prev;
    }
    if (this->next != NULL) {
      delete this->next;
    }
  }
};


class LRUCacheYellow {
   public:
      LRUCacheYellow(int capacity);
      int get(int key);
      void set(int key, int value);
      ~LRUCacheYellow();
 
   private:
      int capacity;
      std::unordered_map<int, cacheNode*> cacheMap;
      cacheNode* first;
      cacheNode* last;
};
 
// Member functions definitions including constructor
LRUCacheYellow::LRUCacheYellow( int capacity) {
   this->capacity = capacity;
}

LRUCacheYellow::~LRUCacheYellow() {
  delete &this->cacheMap;
}

int LRUCacheYellow::get(int key) {
  if (this->cacheMap.find(key) == this->cacheMap.end()) {
    return -1;
  }

 // printf("get key (exists) %d\n", key);

  if (this->cacheMap.size() == 1) {
    // do not process any changes, only one element in a map
    return this->cacheMap[key]->value;
  }
  // remove node from the list
  cacheNode* node = this->cacheMap[key];
  cacheNode* prev = node->prev;
  cacheNode* next = node->next;
  // set prev to next
  if (prev != NULL) {
    prev->next = next;
  }
  // move node to top of a list, because it was recently accessed
  node->prev = NULL;
  node->next = this->first;
  this->first = node;

  // if it was last node set last to prev
  if (this->last == node) {
    this->last = prev;
  }
  return node->value;
}

void LRUCacheYellow::set(int key, int value) {
  printf("set method start last %p\n", this->last);
  printf("%lu size ===\n", this->cacheMap.size());
  printf("%d capacity ===\n", this->capacity);
  if (this->cacheMap.find(key) != this->cacheMap.end()) {
    printf("setting exsisting value %d\n", key);
    // so we set value in exsisting index
    // it means that capacity will not changed
    // remove it from list and set it to the top of list
    // to move it to the top of list simply call get method :)
    this->get(key);
    // and than change top value
    this->first->value = value;
    return;
  }

  if (this->cacheMap.size() >= this->capacity) {
   printf("setting in full cache; key value %d %d\n", key, value);
   printf("last %p\n", this->last);
   printf("last key value %d %d\n", this->last->key, this->last->value);
    // delete last element to free space
    // if (this->last != NULL) {
    cacheNode* node = this->last;
    cacheNode* prev = node->prev;
      if (prev != NULL) {
        // destroy relation with last node
        prev->next = NULL;
      }
    this->last = prev;
    printf("deleting prev last node: key value === %d %d\n", node->key, node->value);
    printf("Value %d \n", this->cacheMap[node->key]->value);
    for (auto p : this->cacheMap) 
        std::cout << p.first << "==>" << ((cacheNode*)p.second)->value << "\n"; 
    // this->cacheMap.erase(this->cacheMap.find(node->key));
    std::cout << this->cacheMap.find(5)->first;
    this->cacheMap.erase(5);
    printf("AFTER ERASE");
    // if we deleteing last which is also first
    if (this->first == node) {
      this->first = NULL;
    }
    delete node;
    // }
  }

  // add a new value
  cacheNode* node = new cacheNode(key, value);
  printf("add key value %d %d\n", node->key, node->value);
  this->cacheMap.emplace(key, node);
  printf("This is value %d\n", this->cacheMap[key]->value);
  printf("map size %lu\n", this->cacheMap.size());
  if (this->first != NULL) {
    node->next = this->first;
    this->first->prev = node;
  }

  this->first = node;

  if (this->cacheMap.size() == 1) {
    // it was the first element inserted
    // set last to first
    this->last = node;
   printf("set last %p\n", this->last);
  }
}

LRUCacheYellow* cache;

class LRUCache {
   public:
      LRUCache(int capacity);
      int get(int key);
      void set(int key, int value);
      ~LRUCache();
};

LRUCache::~LRUCache() {
  delete &cache;
}

LRUCache::LRUCache(int capacity) {
  cache = new LRUCacheYellow(capacity);
}

int LRUCache::get(int key) {
  return (*cache).get(key);
}

void LRUCache::set(int key, int value) {
  (*cache).set(key, value);
}

void testCase1() {
  LRUCache* c = new LRUCache(2);
    (*c).set(1, 10);
    (*c).set(5, 12);
   printf("%d\n", (*c).get(5));
   printf("%d\n", (*c).get(1));
   printf("%d\n", (*c).get(10));
    (*c).set(6, 14);
   printf("%d\n", (*c).get(5));
}

void testCase2() {
  LRUCache* c = new LRUCache(1);
    (*c).set(2, 1);
    (*c).set(2, 2);
   printf("%d\n", (*c).get(2));
    (*c).set(1, 1);
    (*c).set(4, 1);
   printf("%d\n", (*c).get(2));
}

void testCase3() {
  LRUCache* c = new LRUCache(4);
  (*c).set(5, 13);
  (*c).set(9, 6);
  (*c).set(4, 1);
  printf("%d\n", (*c).get(4));
  printf("before set\n");
  (*c).set(6, 1);
  printf("before set\n");
  (*c).set(8, 11);
  printf("before get\n");
  printf("%d\n", (*c).get(13));
  printf("%d\n", (*c).get(1));
  (*c).set(12, 12);
  printf("%d\n", (*c).get(10));
  (*c).set(15, 13);
  (*c).set(2, 13);
  (*c).set(7, 5);
  (*c).set(10, 3);
  printf("%d\n", (*c).get(6));
  printf("%d\n", (*c).get(10));
  (*c).set(15, 14);
  (*c).set(5, 12);
  printf("%d\n", (*c).get(5));
  printf("%d\n", (*c).get(7));
  printf("%d\n", (*c).get(15));
  printf("%d\n", (*c).get(5));
  printf("%d\n", (*c).get(6));
  printf("%d\n", (*c).get(10));
  (*c).set(7, 13);
  printf("%d\n", (*c).get(14));
  (*c).set(8, 9);
  printf("%d\n", (*c).get(4));
  (*c).set(6, 11);
  printf("%d\n", (*c).get(9));
  (*c).set(6, 12);
  printf("%d\n", (*c).get(3));
  
}


int main()
{
  // testCase1();
  // printf("====\n");
  // testCase2();
  printf("====\n");
  testCase3();
  return 0;
}

