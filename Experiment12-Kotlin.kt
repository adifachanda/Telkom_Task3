fun generateData(n: Int): List<String> {
    val baseNumber = "082"
    val customers = mutableListOf<String>()

    for (i in 0 until n) {
        val mobileNumber = baseNumber + "%09d".format(i)
        customers.add(mobileNumber)
    }

    return customers
}

fun sendPromoDiscount(input: List<String>) {
    for (mobileNumber in input) {
        println("Sending Promo to $mobileNumber")
    }
    println("It's Finished to Send Promo to ${input.size} Customers")
    for (mobileNumber in input) {
        println("Sending Discount to $mobileNumber")
    }
    println("It's Finished to Send Discount to ${input.size} Customers")
}

fun main() {
    val data = generateData(1000)
    sendPromoDiscount(data)
}
