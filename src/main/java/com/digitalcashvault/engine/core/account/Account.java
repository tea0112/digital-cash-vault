package com.digitalcashvault.engine.core.account;

import java.math.BigDecimal;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;
import lombok.NoArgsConstructor;

@Entity
@Table(name = "accounts")
@NoArgsConstructor
public class Account {
  @Id
  @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "account_sequence")
  @SequenceGenerator(name = "account_sequence", sequenceName = "ACCOUNT_SEQ")
  private Long id;

  @Column(name = "account_number", nullable = false)
  private String accountNumber;

  @Column(name= "balance", nullable = false)
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

  public String getAccountNumber() {
    return this.accountNumber;
  }

  public void SetAccountNumber(String accountNumber) {
    this.accountNumber = accountNumber;
  }

  public BigDecimal getBalance() {
    return this.balance;
  }

  public void setBalance(BigDecimal balance) {
    this.balance = balance;
  }
}
