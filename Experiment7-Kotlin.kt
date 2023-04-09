fun main() {
    val foods = arrayOf("sate", "bakso", "dimsum", "rames")
    val drinks = arrayOf("jeruk", "teh", "kelapa", "cendol")
    logPairs(foods, drinks, "Menu")
}

fun logPairs(array1: Array<String>, array2: Array<String>, words: String) {
    var counter = 0
    for (i in array1.indices) {
        for (j in array2.indices) {
            counter++
            println("$words ${array1[i]} dan ${array2[j]}")
        }
    }
}
