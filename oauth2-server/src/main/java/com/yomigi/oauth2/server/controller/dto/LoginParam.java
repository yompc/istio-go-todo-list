package com.yomigi.oauth2.server.controller.dto;

import lombok.Data;

import java.io.Serializable;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/5
 * @time: 18:12
 */
@Data
public class LoginParam implements Serializable {
    private String username;
    private String password;
}
