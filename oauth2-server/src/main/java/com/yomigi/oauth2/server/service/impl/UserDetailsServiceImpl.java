package com.yomigi.oauth2.server.service.impl;

import com.google.common.collect.Lists;
import com.yomigi.oauth2.server.domain.PermissionEntity;
import com.yomigi.oauth2.server.domain.UserEntity;
import com.yomigi.oauth2.server.repository.PermissionRepository;
import com.yomigi.oauth2.server.repository.UserRepository;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.List;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 13:19
 */
@Component
public class UserDetailsServiceImpl implements UserDetailsService {

    @Resource
    UserRepository userRepository;

    @Resource
    PermissionRepository permissionRepository;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        UserEntity oneByUsername = userRepository.findOneByUsername(username);
        List<GrantedAuthority> grantedAuthorities = Lists.newArrayList();

        if (oneByUsername!=null)
        {
            List<PermissionEntity> tbPermissions = permissionRepository.selectRbacPermission(oneByUsername.getUsername());
            tbPermissions.forEach(tbPermission -> {
                if (tbPermission!=null&&tbPermission.getPermissionName()!=null)
                {
                    GrantedAuthority grantedAuthority=new SimpleGrantedAuthority(tbPermission.getPermissionName());
                    grantedAuthorities.add(grantedAuthority);
                }
            });
        }

        return new User(username,oneByUsername.getPassword(),grantedAuthorities);
    }
}
