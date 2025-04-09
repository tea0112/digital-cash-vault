package com.digitalcashvault.engine.core.permission;

import java.util.Set;

import com.digitalcashvault.engine.core.role.Role;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Table(name = "permissions")
@Getter
@Setter
@NoArgsConstructor
public class Permission {
  @Id
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "permission_sequence")
  @SequenceGenerator(name = "permission_sequence", sequenceName = "PERMISSION_SEQ")
  private Long id;

  @Column(name = "name", nullable = false)
  private String name;

  @ManyToMany(mappedBy = "permissions")
  private Set<Role> roles;

  @ManyToMany(mappedBy = "permissions")
  private Set<Module> modules;
}
