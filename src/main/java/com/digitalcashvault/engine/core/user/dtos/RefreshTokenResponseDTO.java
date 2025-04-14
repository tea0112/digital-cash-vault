package com.digitalcashvault.engine.core.user.dtos;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class RefreshTokenResponseDTO {
  private String accessToken;
  private Long accessTokenExpirationInMs;
  private String refreshToken;
  private Long refreshTokenExpirationInMs;
}
