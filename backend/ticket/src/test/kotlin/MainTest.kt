package net.mioyi.railway.ticket

import io.quarkus.test.junit.QuarkusTest
import io.restassured.module.kotlin.extensions.Then
import io.restassured.module.kotlin.extensions.When
import org.hamcrest.Matchers.equalTo
import kotlin.test.Test

@QuarkusTest
class MainTest {
    @Test
    fun f() {
        When {
            get("/")
        }.Then {
            statusCode(200)
            body(equalTo("ticket"))
        }
    }
}