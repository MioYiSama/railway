package net.mioyi.railway.user

import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured.`when`
import org.hamcrest.Matchers.equalTo
import kotlin.test.Test
import kotlin.test.assertEquals

@QuarkusTest
class MainTest {
    @Test
    fun f() {
        `when`().get("/").then()
            .statusCode(200)
            .body(equalTo("user"))
    }
}