package com.yomigi.oauth2.server.domain;


import lombok.Data;

import javax.persistence.Column;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import javax.persistence.Entity;
import java.time.LocalDate;

/**
 * 用户模型
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 13:22
 */
@Data
@Entity
@Table(name = "tb_user")
public class UserEntity {
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Id
    private long id;
    /**
     * 用户名
     */
    private String username ;
    /**
     * 密码
     */
    private String password;
    /**
     * 备注
     */
    private String description;
    /**
     * 刷新令牌
     */
    @Column(name = "refresh_token")
    private String refreshToken;
    /**
     * 手机号码
     */
    private String phone;
    /**
     * 创建时间
     */
    private LocalDate created;
    /**
     * 更新时间
     */
    private LocalDate updated;
}
