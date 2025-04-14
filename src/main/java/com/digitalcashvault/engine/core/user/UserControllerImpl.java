package com.digitalcashvault.engine.core.user;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.digitalcashvault.engine.core.user.dtos.RefreshTokenRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginResponseDTO;
import com.digitalcashvault.engine.core.user.dtos.UserRegistrationRequestDTO;

import io.jsonwebtoken.JwtException;

@RestController
@RequestMapping("/api/users")
public class UserControllerImpl implements UserController {
  private final UserService userService;

  public UserControllerImpl(UserService userService) {
    this.userService = userService;
  }

  @PostMapping("/register")
  public ResponseEntity<String> registerUser(@RequestBody UserRegistrationRequestDTO request) {
    this.userService.registerUser(request);

    return ResponseEntity.status(HttpStatus.CREATED).body("User registered successfully");
  }

  @PostMapping("/login")
  public ResponseEntity<UserLoginResponseDTO> login(@RequestBody UserLoginRequestDTO request) {
    UserLoginResponseDTO response = this.userService.login(request);

    return ResponseEntity.status(HttpStatus.OK).body(response);
  }

  @PostMapping("/refresh")
  public ResponseEntity<RefreshTokenRequestDTO> refresh(@RequestBody RefreshTokenRequestDTO request) {
    // TODO: implement me
    throw new JwtException("implement");
  }
}
