package com.digitalcashvault.engine.core.user;

import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.digitalcashvault.engine.core.user.dtos.UserRegistrationRequestDTO;

@Service
public class UserService {
  private final UserRepository userRepository;
  private final PasswordEncoder passwordEncoder;

  public UserService(UserRepository userRepository, PasswordEncoder passwordEncoder) {
    this.userRepository = userRepository;
    this.passwordEncoder = passwordEncoder;
  }

  @Transactional
  public void registerUser(UserRegistrationRequestDTO request) {
    if (userRepository.existsByEmail(request.getEmail())) {
      throw new RuntimeException("Email already in use");
    }

    if (userRepository.existsByUsername(request.getUsername())) {
      throw new RuntimeException("Username already taken");
    }

    User user = new User();
    user.setUsername(request.getUsername());
    user.setEmail(request.getEmail());
    user.setPassword(this.passwordEncoder.encode(request.getPassword()));

    this.userRepository.save(user);
  }
}
