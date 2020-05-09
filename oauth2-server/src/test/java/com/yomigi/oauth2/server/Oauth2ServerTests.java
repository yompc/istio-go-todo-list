package com.yomigi.oauth2.server;

import com.yomigi.oauth2.server.domain.PermissionEntity;
import com.yomigi.oauth2.server.repository.PermissionRepository;
import org.bouncycastle.crypto.signers.ISOTrailers;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.test.context.junit4.SpringRunner;

import javax.annotation.Resource;
import java.util.List;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 18:18
 */
@RunWith(SpringRunner.class)
@SpringBootTest
public class Oauth2ServerTests {

    @Resource
    PermissionRepository permissionRepository;

    @Resource
    BCryptPasswordEncoder bCryptPasswordEncoder;

    @Test
    public void datebaseTest(){
        List<PermissionEntity> admin = permissionRepository.selectRbacPermission("admin");
        System.out.println(admin);
        String encode = bCryptPasswordEncoder.encode("123456");
        System.out.println(encode);
    }

}
