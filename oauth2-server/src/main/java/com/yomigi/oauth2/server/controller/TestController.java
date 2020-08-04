package com.yomigi.oauth2.server.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author 刘锋嘉
 * @since 2020/8/4 9:42
 */
@RestController
public class TestController {
    @GetMapping("/test")
    public String test(){
        return "OK";
    }
}
