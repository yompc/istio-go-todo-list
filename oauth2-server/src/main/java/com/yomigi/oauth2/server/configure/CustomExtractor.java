package com.yomigi.oauth2.server.configure;

import org.springframework.security.core.Authentication;
import org.springframework.security.oauth2.provider.authentication.TokenExtractor;
import org.springframework.security.web.authentication.preauth.PreAuthenticatedAuthenticationToken;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import java.util.Arrays;

/**
 * @author 刘锋嘉
 * @since 2020/8/4 10:01
 */
public class CustomExtractor implements TokenExtractor {
    private static final String TOKEN_KEY_JWT = "token";

    @Override
    public Authentication extract(HttpServletRequest request) {
        return new PreAuthenticatedAuthenticationToken(getTokenFromRequest(request), "");
    }

    private String getTokenFromRequest(HttpServletRequest request) {
        final Cookie[] cookies = request.getCookies();
        if (cookies == null) {
            return null;
        }
        return Arrays.stream(cookies)
                .filter(cookie -> cookie.getName().equals(TOKEN_KEY_JWT))
                .findFirst()
                .map(Cookie::getValue).orElse(null);
    }
}