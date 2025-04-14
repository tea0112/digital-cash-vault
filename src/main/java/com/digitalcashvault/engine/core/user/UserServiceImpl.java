package com.digitalcashvault.engine.core.user;

import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.digitalcashvault.engine.config.JwtConfigs;
import com.digitalcashvault.engine.config.JwtService;
import com.digitalcashvault.engine.core.user.dtos.RefreshTokenRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.RefreshTokenResponseDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginRequestDTO;
import com.digitalcashvault.engine.core.user.dtos.UserLoginResponseDTO;
import com.digitalcashvault.engine.core.user.dtos.UserRegistrationRequestDTO;

@Service
public class UserServiceImpl implements UserService {
  private final UserRepository userRepository;
  private final PasswordEncoder passwordEncoder;
  private final JwtService jwtService;

  public UserServiceImpl(UserRepository userRepository, PasswordEncoder passwordEncoder, JwtService jwtService) {
    this.userRepository = userRepository;
    this.passwordEncoder = passwordEncoder;
    this.jwtService = jwtService;
  }

  @Transactional
  public void registerUser(UserRegistrationRequestDTO request) {
    if (this.userRepository.existsByEmail(request.getEmail())) {
      throw new RuntimeException("Email already in use");
    }

    if (this.userRepository.existsByUsername(request.getUsername())) {
      throw new RuntimeException("Username already taken");
    }

    User user = new User();
    user.setUsername(request.getUsername());
    user.setEmail(request.getEmail());
    user.setPassword(this.passwordEncoder.encode(request.getPassword()));

    this.userRepository.save(user);
  }

  @Transactional
  public UserLoginResponseDTO login(UserLoginRequestDTO request) {
    String usernameOrPasswordWrongExceptionMsg = "Username, Email or Password wrong";
    boolean usernameExists = this.userRepository.existsByUsername(request.getUsername());
    boolean emailExists = this.userRepository.existsByEmail(request.getEmail());

    if (!usernameExists || !emailExists) {
      throw new BadCredentialsException(usernameOrPasswordWrongExceptionMsg);
    }

    User user = null;
    if (usernameExists) {
      user = this.userRepository.findByEmail(request.getEmail())
          .orElseThrow(() -> new BadCredentialsException(usernameOrPasswordWrongExceptionMsg));
    } else {
      user = this.userRepository.findByUsername(request.getUsername())
          .orElseThrow(() -> new BadCredentialsException(usernameOrPasswordWrongExceptionMsg));
    }

    String encodedPassword = this.passwordEncoder.encode(request.getPassword());
    if (!this.passwordEncoder.matches(encodedPassword, user.getPassword())) {
      throw new InternalError(usernameOrPasswordWrongExceptionMsg);
    }

    String accessToken = this.jwtService.generateAccessToken(new UserDetailsImpl(user),
        JwtConfigs.ACCESS_TOKEN_EXPIRATION);
    String refreshToken = this.jwtService.generateRefreshToken(new UserDetailsImpl(user),
        JwtConfigs.REFRESH_TOKEN_EXPIRATION);

    UserLoginResponseDTO response = new UserLoginResponseDTO();
    response.setAccessToken(accessToken);
    response.setAccessTokenExpirationInMs(JwtConfigs.ACCESS_TOKEN_EXPIRATION);
    response.setRefreshToken(refreshToken);
    response.setAccessTokenExpirationInMs(JwtConfigs.REFRESH_TOKEN_EXPIRATION);

    return response;
  }

  @Override
  public RefreshTokenResponseDTO refresh(RefreshTokenRequestDTO request) {
    // TODO Auto-generated method stub
    return null;
  }
}
