package com.yomigi.oauth2.server.domain;

import lombok.Data;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

/**
 * 用户和角色关联表
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 13:55
 */
@Data
@Table(name = "tb_user_role")
@Entity
public class UserRoleEntity {
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Id
    private long id;
    /**
     * 用户ID
     */
    @Column(name = "user_id")
    private long userId;
    /**
     * 角色ID
     */
    @Column(name = "role_id")
    private long roleId;
}
