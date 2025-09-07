package task19

import (
	"sync"
	"testing"
	"time"
)

func TestCacheBasic(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 50,
		MaxSize:         10,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элемент
	cache.Set("key1", "value1")

	// Проверяем получение
	value, ok := cache.Get("key1")
	if !ok {
		t.Error("Should be able to get value")
	}
	if value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}

	// Проверяем размер
	if cache.Size() != 1 {
		t.Errorf("Expected size 1, got %d", cache.Size())
	}
}

func TestCacheTTL(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      50,
		CleanupInterval: 10,
		MaxSize:         10,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элемент
	cache.Set("key1", "value1")

	// Проверяем, что элемент есть
	value, ok := cache.Get("key1")
	if !ok {
		t.Error("Should be able to get value")
	}
	if value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}

	// Ждем истечения TTL
	time.Sleep(60 * time.Millisecond)

	// Проверяем, что элемент истек
	_, ok = cache.Get("key1")
	if ok {
		t.Error("Value should be expired")
	}
}

func TestCacheCustomTTL(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 10,
		MaxSize:         10,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элемент с кастомным TTL
	cache.SetWithTTL("key1", "value1", 50)

	// Проверяем, что элемент есть
	value, ok := cache.Get("key1")
	if !ok {
		t.Error("Should be able to get value")
	}
	if value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}

	// Ждем истечения TTL
	time.Sleep(60 * time.Millisecond)

	// Проверяем, что элемент истек
	_, ok = cache.Get("key1")
	if ok {
		t.Error("Value should be expired")
	}
}

func TestCacheDelete(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 50,
		MaxSize:         10,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элемент
	cache.Set("key1", "value1")

	// Проверяем, что элемент есть
	_, ok := cache.Get("key1")
	if !ok {
		t.Error("Should be able to get value")
	}

	// Удаляем элемент
	cache.Delete("key1")

	// Проверяем, что элемент удален
	_, ok = cache.Get("key1")
	if ok {
		t.Error("Value should be deleted")
	}
}

func TestCacheClear(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 50,
		MaxSize:         10,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элементы
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Проверяем размер
	if cache.Size() != 2 {
		t.Errorf("Expected size 2, got %d", cache.Size())
	}

	// Очищаем кэш
	cache.Clear()

	// Проверяем, что кэш пуст
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}

	_, ok := cache.Get("key1")
	if ok {
		t.Error("Value should be cleared")
	}
}

func TestCacheMaxSize(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 50,
		MaxSize:         2,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элементы до превышения лимита
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	// Проверяем, что размер не превышает лимит
	if cache.Size() > 2 {
		t.Errorf("Size should not exceed max size, got %d", cache.Size())
	}
}

func TestCacheStats(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 50,
		MaxSize:         10,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Добавляем элементы
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Получаем статистику
	stats := cache.GetStats()
	if stats == nil {
		t.Error("Stats should not be nil")
	}

	// Проверяем размер в статистике
	if stats["size"] != 2 {
		t.Errorf("Expected size 2 in stats, got %d", stats["size"])
	}
}

func TestCacheConcurrent(t *testing.T) {
	config := CacheConfig{
		DefaultTTL:      100,
		CleanupInterval: 50,
		MaxSize:         100,
	}

	cache := NewCache(config)
	defer cache.Close()

	var wg sync.WaitGroup

	// Запускаем горутины, которые добавляют элементы
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := "key" + string(rune(index))
			value := "value" + string(rune(index))
			cache.Set(key, value)
		}(i)
	}

	wg.Wait()

	// Проверяем, что все элементы добавлены
	if cache.Size() != 10 {
		t.Errorf("Expected size 10, got %d", cache.Size())
	}
}

func TestLRUCacheBasic(t *testing.T) {
	lru := NewLRUCache(3)

	// Добавляем элементы
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")

	// Проверяем получение
	value, ok := lru.Get("key1")
	if !ok {
		t.Error("Should be able to get value")
	}
	if value != "value1" {
		t.Errorf("Expected 'value1', got %v", value)
	}

	// Добавляем еще один элемент (должен вытеснить key2)
	lru.Set("key4", "value4")

	// Проверяем, что key2 вытеснен
	_, ok = lru.Get("key2")
	if ok {
		t.Error("key2 should be evicted")
	}

	// Проверяем, что key1 все еще есть
	_, ok = lru.Get("key1")
	if !ok {
		t.Error("key1 should still be in cache")
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	lru := NewLRUCache(3)

	// Добавляем элемент
	lru.Set("key1", "value1")

	// Обновляем элемент
	lru.Set("key1", "value1_updated")

	// Проверяем обновленное значение
	value, ok := lru.Get("key1")
	if !ok {
		t.Error("Should be able to get value")
	}
	if value != "value1_updated" {
		t.Errorf("Expected 'value1_updated', got %v", value)
	}
}

func TestLRUCacheAccessOrder(t *testing.T) {
	lru := NewLRUCache(3)

	// Добавляем элементы
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")

	// Обращаемся к key1 (должен стать самым недавно использованным)
	lru.Get("key1")

	// Добавляем новый элемент (должен вытеснить key2)
	lru.Set("key4", "value4")

	// Проверяем, что key2 вытеснен
	_, ok := lru.Get("key2")
	if ok {
		t.Error("key2 should be evicted")
	}

	// Проверяем, что key1 все еще есть
	_, ok = lru.Get("key1")
	if !ok {
		t.Error("key1 should still be in cache")
	}
}

