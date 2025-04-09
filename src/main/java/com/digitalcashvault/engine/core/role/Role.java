package com.digitalcashvault.engine.core.role;

import java.util.Set;

import com.digitalcashvault.engine.core.permission.Permission;
import com.digitalcashvault.engine.core.user.User;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.JoinTable;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Table(name = "roles")
@Getter
@Setter
@NoArgsConstructor
public class Role {
  @Id
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "role_sequence")
  @SequenceGenerator(name = "role_sequence", sequenceName = "ROLE_SEQ")
  private Long id;

  @Column(name = "name", nullable = false)
  private String name;

  @ManyToMany(mappedBy = "roles")
  private Set<User> users;

  @ManyToMany
  @JoinTable(
    name = "role_permission",
    joinColumns = @JoinColumn(name = "role_id"),
    inverseJoinColumns = @JoinColumn(name = "permission_id")
  )
  private Set<Permission> permissions;

  public Role(String name) {
    this.name = name;
  }
}
