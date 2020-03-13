package com.bulletjournal.controller;

import com.bulletjournal.clients.UserClient;
import com.bulletjournal.controller.models.CreateTransactionParams;
import com.bulletjournal.controller.models.Transaction;
import com.bulletjournal.controller.models.UpdateTransactionParams;
import com.bulletjournal.controller.utils.EtagGenerator;
import com.bulletjournal.notifications.*;
import com.bulletjournal.repository.TransactionDaoJpa;
import org.slf4j.MDC;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;
import java.util.List;

@RestController
public class TransactionController {
    protected static final String TRANSACTIONS_ROUTE = "/api/projects/{projectId}/transactions";
    protected static final String TRANSACTION_ROUTE = "/api/transactions/{transactionId}";

    @Autowired
    private TransactionDaoJpa transactionDaoJpa;

    @Autowired
    private NotificationService notificationService;

    @GetMapping(TRANSACTIONS_ROUTE)
    public ResponseEntity<List<Transaction>> getTransactions(@NotNull @PathVariable Long projectId) {

        List<Transaction> transactions = this.transactionDaoJpa.getTransactions(projectId);
        String transactionsEtag = EtagGenerator.generateEtag(EtagGenerator.HashAlgorithm.MD5,
                EtagGenerator.HashType.TO_HASHCODE, transactions);

        HttpHeaders responseHeader = new HttpHeaders();
        responseHeader.setETag(transactionsEtag);

        return ResponseEntity.ok().headers(responseHeader).body(transactions);
    }

    @PostMapping(TRANSACTIONS_ROUTE)
    @ResponseStatus(HttpStatus.CREATED)
    public Transaction createTransaction(@NotNull @PathVariable Long projectId,
                                         @Valid @RequestBody CreateTransactionParams createTransactionParams) {
        String username = MDC.get(UserClient.USER_NAME_KEY);
        return transactionDaoJpa.create(projectId, username, createTransactionParams).toPresentationModel();
    }

    @GetMapping(TRANSACTION_ROUTE)
    public Transaction getTransaction(@NotNull @PathVariable Long transactionId) {
        return this.transactionDaoJpa.getTransaction(transactionId);
    }

    @PatchMapping(TRANSACTION_ROUTE)
    public Transaction updateTransaction(@NotNull @PathVariable Long transactionId,
                                         @Valid @RequestBody UpdateTransactionParams updateTransactionParams) {
        String username = MDC.get(UserClient.USER_NAME_KEY);
        List<Event> events = transactionDaoJpa.partialUpdate(username, transactionId, updateTransactionParams);
        if (!events.isEmpty()) {
            notificationService.inform(new UpdateTransactionPayerEvent(events, username, updateTransactionParams.getPayer()));
        }
        return getTransaction(transactionId);
    }

    @DeleteMapping(TRANSACTION_ROUTE)
    public void deleteTransaction(@NotNull @PathVariable Long transactionId) {
        String username = MDC.get(UserClient.USER_NAME_KEY);
        List<Event> events = this.transactionDaoJpa.delete(username, transactionId);
        if (!events.isEmpty()) {
            this.notificationService.inform(new RemoveTransactionEvent(events, username));
        }
    }
}
