package task19

// Task 19: Concurrent Cache with TTL
// Реализуйте потокобезопасный кэш с TTL (Time To Live).

// "fmt"
// "sync"
// "time"

// CacheItem представляет элемент кэша
type CacheItem struct {
	Value     interface{}
	ExpiresAt int64 // Unix timestamp в миллисекундах
}

// Cache представляет потокобезопасный кэш
type Cache struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map для хранения элементов
	// - mutex для синхронизации
	// - канал для очистки истекших элементов
}

// CacheConfig содержит настройки кэша
type CacheConfig struct {
	DefaultTTL      int64 // TTL по умолчанию в миллисекундах
	CleanupInterval int64 // Интервал очистки в миллисекундах
	MaxSize         int   // Максимальный размер кэша
}

// NewCache создает новый кэш
func NewCache(config CacheConfig) *Cache {
	// TODO: Реализуйте конструктор
	// 1. Инициализируйте поля
	// 2. Запустите горутину для очистки истекших элементов
	return nil
}

// Set сохраняет значение в кэш
func (c *Cache) Set(key string, value interface{}) {
	// TODO: Реализуйте метод
	// 1. Создайте CacheItem с TTL
	// 2. Сохраните в map
	// 3. Проверьте максимальный размер
}

// SetWithTTL сохраняет значение с кастомным TTL
func (c *Cache) SetWithTTL(key string, value interface{}, ttl int64) {
	// TODO: Реализуйте метод
}

// Get получает значение из кэша
func (c *Cache) Get(key string) (interface{}, bool) {
	// TODO: Реализуйте метод
	// 1. Проверьте наличие ключа
	// 2. Проверьте, не истек ли TTL
	// 3. Верните значение и флаг успеха
	return nil, false
}

// Delete удаляет значение из кэша
func (c *Cache) Delete(key string) {
	// TODO: Реализуйте метод
}

// Clear очищает весь кэш
func (c *Cache) Clear() {
	// TODO: Реализуйте метод
}

// Size возвращает размер кэша
func (c *Cache) Size() int {
	// TODO: Реализуйте метод
	return 0
}

// GetStats возвращает статистику кэша
func (c *Cache) GetStats() map[string]int {
	// TODO: Реализуйте метод
	// Верните map с ключами: "size", "hits", "misses", "expired"
	return nil
}

// Close закрывает кэш
func (c *Cache) Close() {
	// TODO: Реализуйте метод
	// Остановите горутину очистки
}

// LRUCache представляет кэш с политикой LRU
type LRUCache struct {
	// TODO: Добавьте необходимые поля
	// Вам понадобятся:
	// - map для быстрого доступа
	// - двусвязный список для порядка
	// - mutex для синхронизации
	// - максимальный размер
}

// NewLRUCache создает новый LRU кэш
func NewLRUCache(maxSize int) *LRUCache {
	// TODO: Реализуйте конструктор
	return nil
}

// Get получает значение из LRU кэша
func (lru *LRUCache) Get(key string) (interface{}, bool) {
	// TODO: Реализуйте метод
	// 1. Найдите элемент в map
	// 2. Переместите в начало списка
	// 3. Верните значение
	return nil, false
}

// Set сохраняет значение в LRU кэш
func (lru *LRUCache) Set(key string, value interface{}) {
	// TODO: Реализуйте метод
	// 1. Проверьте, есть ли элемент
	// 2. Если есть - обновите и переместите в начало
	// 3. Если нет - добавьте новый
	// 4. Проверьте максимальный размер
}

