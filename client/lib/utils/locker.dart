typedef VoidCallback = void Function();

class Locker<T> {
  final Map<T, bool> _lockMap = {};

  void lock(T key) {
    _lockMap[key] = true;
  }

  void unlock(T key) {
    _lockMap[key] = false;
  }

  bool isLocked(T key) {
    return _lockMap[key] ?? false;
  }

  void execute(T key, VoidCallback action) {
    if (!isLocked(key)) {
      lock(key);
      action();
      unlock(key);
    }
  }
}
