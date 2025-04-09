package com.digitalcashvault.engine.core.user;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import com.digitalcashvault.engine.core.permission.Permission;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
public class UserDetailsImpl implements UserDetails {
  private User user;

  @Override
  public String getUsername() {
    return this.user.getUsername();
  }

  @Override
  public String getPassword() {
    return this.user.getPassword();
  }

  @Override
  public Collection<? extends GrantedAuthority> getAuthorities() {
    List<GrantedAuthority> authorities = new ArrayList<>();

    authorities
        .addAll(
            this.user.getRoles().stream().map((role) -> new SimpleGrantedAuthority("ROLE_" + role.getName())).toList());

    Set<Permission> permissions = this.user.getRoles()
        .stream()
        .flatMap((role) -> role.getPermissions().stream())
        .collect(Collectors.toSet());

    authorities.addAll(permissions
        .stream()
        .map(permission -> new SimpleGrantedAuthority(permission.getName())).toList());

    return authorities;
  }

  @Override
  public boolean isAccountNonExpired() {
    return true;
  }

  @Override
  public boolean isAccountNonLocked() {
    return true;
  }

  @Override
  public boolean isCredentialsNonExpired() {
    return true;
  }

  @Override
  public boolean isEnabled() {
    return this.user.isEnabled();
  }
}
