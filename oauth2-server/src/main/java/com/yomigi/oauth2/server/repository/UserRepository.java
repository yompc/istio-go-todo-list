package com.yomigi.oauth2.server.repository;

import com.yomigi.oauth2.server.domain.UserEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 14:04
 */
@Repository
public interface UserRepository extends JpaRepository<UserEntity, Long> {
    /**
     * 根据用户名查询单个用户
     * @param username 用户名
     * @return 用户信息
     */
    UserEntity findOneByUsername(String username);

    /**
     * 查询所有用户
     * @param username 用户名
     * @return 所有和用户名匹配的实体
     */
    List<UserEntity> findByUsername(String username);
}
