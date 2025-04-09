package com.digitalcashvault.engine.core.module;

import java.util.Set;

import com.digitalcashvault.engine.core.permission.Permission;

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
@Table(name = "modules")
@Getter
@Setter
@NoArgsConstructor
public class Module {
  @Id
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "module_sequence")
  @SequenceGenerator(name = "module_sequence", sequenceName = "MODULE_SEQ")
  private Long id;

  @Column(name = "name", nullable = false)
  private String name;

  @Column(name = "path", nullable = false)
  private String path;

  @ManyToMany
  @JoinTable(
    name = "module_permission",
    joinColumns = @JoinColumn(name = "module_id"),
    inverseJoinColumns = @JoinColumn(name = "permission_id")
  )
  private Set<Permission> permissions;
}
