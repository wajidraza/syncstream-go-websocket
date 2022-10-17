# High-Throughput Memory & Redis Cache Manager
import time

class CacheManager:
    def __init__(self):
        self._cache = {}
        
    def get(self, key: str):
        if key in self._cache:
            val, exp = self._cache[key]
            if time.time() < exp:
                return val
            del self._cache[key]
        return None
        
    def set(self, key: str, val, ttl: int = 300):
        self._cache[key] = (val, time.time() + ttl)

cache_manager = CacheManager()
