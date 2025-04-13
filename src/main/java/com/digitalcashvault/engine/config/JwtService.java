package com.digitalcashvault.engine.config;

import java.nio.charset.StandardCharsets;
import java.util.Date;

import javax.crypto.SecretKey;
import javax.crypto.spec.SecretKeySpec;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import jakarta.annotation.PostConstruct;

@Component
public class JwtService {
  @Value("${jwt.secret}")
  private String secretKeyString;
  private SecretKey secretKey;

  @PostConstruct
  public void init() {
    byte[] keyBytes = this.secretKeyString.getBytes(StandardCharsets.UTF_8);
    this.secretKey = new SecretKeySpec(keyBytes, SignatureAlgorithm.HS256.getJcaName());
  }

  public String generateToken(UserDetails userDetails, long expirationMs) {
    return Jwts
        .builder()
        .setSubject(userDetails.getUsername())
        .claim("roles", userDetails.getAuthorities())
        .setIssuedAt(new Date())
        .setExpiration(new Date(System.currentTimeMillis() + expirationMs))
        .signWith(this.secretKey, SignatureAlgorithm.HS256)
        .compact();
  }

  public Claims parseToken(String token) {
    Jws<Claims> jws = Jwts.parserBuilder()
        .setSigningKey(this.secretKey)
        .build()
        .parseClaimsJws(token);

    return jws.getBody();
  }
}
