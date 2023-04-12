val results = mutableListOf<List<String>>()
val candidates = listOf("Baswedan", "Subianto", "Maharani")

fun arrange(array: MutableList<String>, memory: List<String> = emptyList()): List<List<String>> {
    var current: String
    for (i in 0 until array.size) {
        current = array.removeAt(i)
        if (array.size == 0) {
            results.add(memory.plus(current))
        }
        arrange(array.toMutableList(), memory.plus(current))
        array.add(i, current)
    }
    return results
}
fun main(){
	val chairs = arrange(candidates.toMutableList())
	for(c in chairs){
        println(c)
    }
}