package com.yomigi.oauth2.server.domain;

import lombok.Data;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

/**
 * 角色和权限关联表
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 13:58
 */
@Data
@Entity
@Table(name = "tb_role_permission")
public class RolePermissionEntity {
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Id
    private long id;
    /**
     * 角色ID
     */
    @Column(name = "role_id")
    private long roleId;
    /**
     * 权限ID
     */
    @Column(name = "permission_id")
    private long permissionId;
}
