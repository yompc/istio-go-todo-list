package com.yomigi.oauth2.server.controller;
import java.time.LocalDate;

import com.yomigi.oauth2.server.controller.dto.RegParam;
import com.yomigi.oauth2.server.controller.dto.ResponseResult;
import com.yomigi.oauth2.server.domain.UserEntity;
import com.yomigi.oauth2.server.repository.UserRepository;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.List;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/5
 * @time: 19:23
 */
@CrossOrigin(origins = "*",maxAge = 3600)
@RestController
@RequestMapping("/api/v1")
public class RegController {
    @Resource
    BCryptPasswordEncoder passwordEncoder;

    @Resource
    UserRepository userRepository;

    @PostMapping(value = "/user/registered")
    public ResponseResult<Void> Reg(@RequestBody RegParam regParam){
        //检查用户名和邮箱有没有重复
        List<UserEntity> resultsUsername = userRepository.findByUsername(regParam.getUsername());
        if (resultsUsername!=null){
            return new ResponseResult<>(ResponseResult.CodeStatus.FAIL,"用户名重复",null);
        }

        //注册
        UserEntity userEntity = new UserEntity();
        userEntity.setUsername(regParam.getUsername());
        userEntity.setPassword(passwordEncoder.encode(regParam.getPassword()));
        userEntity.setCreated(LocalDate.now());
        userEntity.setUpdated(LocalDate.now());

        userRepository.save(userEntity);

        return new ResponseResult<Void>(ResponseResult.CodeStatus.OK,"注册成功",null);

    }
}
