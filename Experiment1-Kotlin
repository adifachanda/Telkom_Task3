fun main() {
    val company = "telkom"
    val largeCompanyName = Array(10) { company }
    findCompany(largeCompanyName)
}

fun findCompany(array: Array<String>) {
    val tx = System.currentTimeMillis()
    for (i in array.indices) {
        if (array[i] == "telkom") {
            println("Company Found!")
        }
    }
    val ty = System.currentTimeMillis()
    println("The Performance is " + (ty - tx) + "ms")
}
