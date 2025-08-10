class Solution {
  bool reorderedPowerOf2(int n) {
    Map<String, int> mapper = {};
    int number = 1;
    String nString = n.toString();
    while (number <= 1e9 + 1) {
      String numberString = number.toString();
      if (numberString.length == nString.length) {
        for (var i = 0; i < numberString.length; i++) {
          mapper[numberString[i]] = (mapper[numberString[i]] ?? 0) + 1;
          mapper[nString[i]] = (mapper[nString[i]] ?? 0) - 1;
        }
        if (!mapper.values.any((ele) => ele != 0)) {
          return true;
        }
        mapper.clear();
      }
      number *= 2;
    }
    return false;
  }
}