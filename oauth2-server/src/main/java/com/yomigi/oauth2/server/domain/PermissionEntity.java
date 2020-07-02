package com.yomigi.oauth2.server.domain;

import lombok.Data;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;
import java.time.LocalDate;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 13:50
 */
@Data
@Entity
@Table(name = "tb_permission")
public class PermissionEntity {

    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Id
    private long id;
    /**
     * 父级ID
     */
    @Column(name = "parent_id")
    private String parentId;
    /**
     * 权限名称
     */
    @Column(name = "permission_name")
    private String permissionName;
    /**
     * 权限别名
     */
    @Column(name = "epermission_name")
    private String epermissionName;
    /**
     * 授权路径
     */
    private String url;
    /**
     * 备注
     */
    private String description;
    /**
     * 创建时间
     */
    private LocalDate created;
    /**
     * 更新时间
     */
    private LocalDate updated;
}
