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
 * 角色模型
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 13:41
 */
@Data
@Entity
@Table(name = "tb_role")
public class RoleEntity {
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Id
    private long id;
    /**
     * 父级ID
     */
    @Column(name = "parent_id")
    private String parentId;
    /**
     * 角色名称
     */
    @Column(name = "role_name")
    private String roleName;
    /**
     * 角色别名
     */
    @Column(name = "erole_name")
    private String eroleName;
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
