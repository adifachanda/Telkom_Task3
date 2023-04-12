fun generateData(n: Int): List<String> {
    val baseNumber = "082"
    val customers = mutableListOf<String>()

    for (i in 0 until n) {
        val mobileNumber = baseNumber + "%09d".format(i)
        customers.add(mobileNumber)
    }
    return customers
}

fun sendPromoDiscount(input1: List<String>, input2: List<String>) {
    for (mobileNumber in input1) {
        println("Sending Promo to $mobileNumber")
    }
    println("It's Finished to Send Promo to ${input1.size} Customers")
    for (mobileNumber in input2) {
        println("Sending Discount to $mobileNumber")
    }
    println("It's Finished to Send Discount to ${input2.size} Customers")
}

fun main() {
    val dataPromo = generateData(100000000)
    val dataDiscount = generateData(1000)
    sendPromoDiscount(dataPromo, dataDiscount)
}
