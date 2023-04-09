fun main() {
    val address = "DKI Jakarta"
    val addresses = Array(10) { address }
    findAddress(addresses)
}

fun findAddress(addresses: Array<String>) {
    val tx = System.currentTimeMillis()
    println("The Default Address is " + addresses[0])
    val ty = System.currentTimeMillis()
    println("The performance is " + (ty - tx) + "ms")
}
