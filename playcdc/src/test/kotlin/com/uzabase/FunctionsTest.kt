package com.uzabase

import com.github.tomakehurst.wiremock.client.WireMock
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe
import io.mockk.clearAllMocks

class FunctionsTest : FreeSpec({
    afterTest {
        clearAllMocks()
    }

    "MappingBuilderをRequestJsonに変換する" - {
        "URL" {
            val mappingBuilder = WireMock.get("/test")
            toRequestJson(mappingBuilder).url shouldBe "/test"
        }

        "METHOD" - {
            "GET" {
                val mappingBuilder = WireMock.get("/test")
                toRequestJson(mappingBuilder).method shouldBe "GET"
            }
            "POST" {
                val mappingBuilder = WireMock.post("/test")
                toRequestJson(mappingBuilder).method shouldBe "POST"
            }
        }
    }
})
