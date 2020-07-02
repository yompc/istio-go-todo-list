package com.yomigi.oauth2.server.repository;

import com.yomigi.oauth2.server.domain.PermissionEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 14:06
 */
@Repository
public interface PermissionRepository extends JpaRepository<PermissionEntity,Long> {
    /**
     * 多表联查出 对应用户的权限
     * 一个用户有多个角色（role） 一个角色有多个权限(permission)
     * @param username 用户名
     * @return 权限集合
     */
    @Query(nativeQuery=true,value="SELECT tb_permission.id, tb_permission.created, tb_permission.description, tb_permission.epermission_name, tb_permission.parent_id , tb_permission.permission_name, tb_permission.updated, tb_permission.url FROM tb_permission JOIN tb_role_permission ON tb_permission.id = tb_role_permission.permission_id JOIN tb_role ON tb_role_permission.role_id = tb_role.id JOIN tb_user_role ON tb_role.id = tb_user_role.role_id JOIN tb_user ON tb_user_role.user_id = tb_user.id WHERE tb_user.username = ?1")
    List<PermissionEntity> selectRbacPermission(String username);
}
