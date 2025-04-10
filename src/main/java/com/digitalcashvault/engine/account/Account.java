package com.digitalcashvault.engine.account;

import java.math.BigDecimal;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;

@Entity
public class Account {
  @Id
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "account_sequence")
  @SequenceGenerator(name = "account_sequence", sequenceName = "ACCOUNT_SEQ")
  private Long id;

  @Column(nullable = false)
  private String accountNumber;

  @Column(nullable = false)
  private BigDecimal balance;

  public Account(String accountNumber, BigDecimal balance) {
    this.accountNumber = accountNumber;
    this.balance = balance;
  }

  public Long GetId() {
    return this.id;
  }


  public void SetId(Long id) {
    this.id = id;
  }

}
