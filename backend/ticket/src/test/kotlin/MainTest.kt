package net.mioyi.railway.ticket

import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured.`when`
import org.hamcrest.Matchers.equalTo
import kotlin.test.Test

@QuarkusTest
class MainTest {
    @Test
    fun f() {
        `when`().get("/").then()
            .statusCode(200)
            .body(equalTo("ticket"))
    }
}