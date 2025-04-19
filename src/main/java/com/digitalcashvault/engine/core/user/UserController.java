package com.digitalcashvault.engine.core.user;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestBody;

import com.digitalcashvault.engine.core.user.dtos.RefreshTokenRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginResponseDTO;
import com.digitalcashvault.engine.core.user.dtos.UserRegistrationRequestDTO;

public interface UserController {
  public ResponseEntity<String> registerUser(@RequestBody UserRegistrationRequestDTO request);

  public ResponseEntity<UserLoginResponseDTO> login(@RequestBody UserLoginRequestDTO request);

  public ResponseEntity<RefreshTokenRequestDTO> refresh(@RequestBody RefreshTokenRequestDTO request);
}
