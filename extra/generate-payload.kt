import java.io.ByteArrayOutputStream
import java.io.ObjectOutputStream
import java.util.Base64

fun main() {
		// Payload is simply a serialized Java string
    val serialized = ByteArrayOutputStream()
        .apply {
            use { baos ->
                ObjectOutputStream(baos).use { oos ->
                    oos.writeObject("!!! VULNERABLE !!!")
                }
            }
        }
        .toByteArray()

    println(Base64.getEncoder().encodeToString(serialized))
}
