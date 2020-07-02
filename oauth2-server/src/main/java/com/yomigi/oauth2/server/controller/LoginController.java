package com.yomigi.oauth2.server.controller;

import com.google.common.collect.Maps;
import com.yomigi.oauth2.server.controller.dto.LoginParam;
import com.yomigi.oauth2.server.controller.dto.ResponseResult;
import com.yomigi.oauth2.server.service.impl.UserDetailsServiceImpl;
import com.yomigi.oauth2.server.utils.MapperUtils;
import com.yomigi.oauth2.server.utils.OkHttpClientUtil;
import okhttp3.Response;
import org.springframework.security.core.userdetails.UserDetails;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.oauth2.common.OAuth2AccessToken;
import org.springframework.security.oauth2.provider.token.TokenStore;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletRequest;
import java.util.Map;
import java.util.Objects;

/**
 * @email yom535@outlook.com
 * @author: 有民(yom535)
 * @date: 2020/5/4
 * @time: 19:29
 */
@CrossOrigin(origins = "*",maxAge = 3600)
@RestController
@RequestMapping("/api/v1")
public class LoginController {
    private static final String URL_OAUTH_TOKEN="http://127.0.0.1:9000/oauth/token";

    @Resource
    UserDetailsServiceImpl userDetailsService;
    @Resource
    BCryptPasswordEncoder passwordEncoder;

    @Resource
    TokenStore tokenStore;

    @PostMapping(value = "/user/login")
    public ResponseResult<Map<String,Object>> login(@RequestBody LoginParam loginParam){
        Map<String,String> params= Maps.newHashMap();
        params.put("username", loginParam.getUsername());
        params.put("password", loginParam.getPassword());
        params.put("grant_type", "password");
        params.put("client_id", "client");
        params.put("client_secret", "secret");

        UserDetails userDetails=userDetailsService.loadUserByUsername(loginParam.getUsername());

        if (userDetails==null||!passwordEncoder.matches(loginParam.getPassword(),userDetails.getPassword()))
        {
            return new ResponseResult<Map<String, Object>>(ResponseResult.CodeStatus.FAIL,"用户名或密码错误",null);
        }
        try {
            Response response = OkHttpClientUtil.getInstance().postData(URL_OAUTH_TOKEN, params);
            String jsonString= Objects.requireNonNull(response.body()).string();
            Map<String, Object> jsonMap = MapperUtils.json2map(jsonString);
            String token = String.valueOf(jsonMap.get("access_token"));
            Map<String,Object> result=Maps.newHashMap();
            result.put("token",token);
            return new ResponseResult<Map<String, Object>>(ResponseResult.CodeStatus.OK,"登录成功",result);

        } catch (Exception e) {
            e.printStackTrace();
        }
        return new ResponseResult<>(ResponseResult.CodeStatus.FAIL,"登录失败",null);
    }
    //注销
    @PostMapping(value = "/users/logout")
    public ResponseResult<Void> logout(HttpServletRequest request){
        String access_token = request.getParameter("access_token");
        OAuth2AccessToken oAuth2AccessToken = tokenStore.readAccessToken(access_token);
        tokenStore.removeAccessToken(oAuth2AccessToken);
        return new ResponseResult<Void>(ResponseResult.CodeStatus.OK,"用户注销",null);
    }
}
