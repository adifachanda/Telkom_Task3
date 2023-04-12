fun generateData(): List<String> {
    val number = "0821234567"
    val customers = mutableListOf<String>()
    var mobileNumber: String

    for (i in 0..99) {
        mobileNumber = if (i < 10) {
            number + "0" + i
        } else {
            number + i
        }
        customers.add(mobileNumber)
    }
    return customers
}

fun sendPromoDiscount(array: List<String>) {
    for (i in array.indices) {
        println("Sending promo to ${array[i]}")
    }
    for (i in 1..10) {
        println("Sending Discount to Chosen Customer $i")
    }
}

fun main() {
    val customers = generateData()
    sendPromoDiscount(customers)
}
