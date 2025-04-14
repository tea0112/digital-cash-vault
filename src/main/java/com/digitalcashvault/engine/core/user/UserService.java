package com.digitalcashvault.engine.core.user;

import com.digitalcashvault.engine.core.user.dtos.RefreshTokenRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.RefreshTokenResponseDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginResponseDTO;
import com.digitalcashvault.engine.core.user.dtos.UserRegistrationRequestDTO;

public interface UserService {
  public void registerUser(UserRegistrationRequestDTO request);
  public UserLoginResponseDTO login(UserLoginRequestDTO request);
  public RefreshTokenResponseDTO refresh(RefreshTokenRequestDTO request);
}
