package com.bulletjournal.daemon;

import com.bulletjournal.calendars.google.Util;
import com.bulletjournal.clients.DaemonServiceClient;
import com.bulletjournal.clients.GoogleCalClient;
import com.bulletjournal.config.NotificationConfig;
import com.bulletjournal.protobuf.daemon.grpc.services.DaemonGrpc;
import com.bulletjournal.protobuf.daemon.grpc.types.StreamMessage;
import com.bulletjournal.protobuf.daemon.grpc.types.SubscribeNotification;
import com.bulletjournal.repository.*;
import com.bulletjournal.repository.models.GoogleCalendarProject;
import com.bulletjournal.util.CustomThreadFactory;
import com.google.api.client.json.gson.GsonFactory;
import com.google.api.services.calendar.Calendar;
import com.google.api.services.calendar.model.Channel;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.io.IOException;
import java.sql.Timestamp;
import java.util.Iterator;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

@Service
public class Cleaner {

    private static final Logger LOGGER = LoggerFactory.getLogger(Cleaner.class);

    private final ExecutorService executorService;

    private static final GsonFactory GSON = new GsonFactory();

    @Autowired
    private GoogleCalClient googleCalClient;

    @Autowired
    private DaemonServiceClient daemonServiceClient;

    @Autowired
    private GoogleCalendarProjectRepository googleCalendarProjectRepository;

    public Cleaner() {
        executorService = Executors.newFixedThreadPool(3, new CustomThreadFactory("cleaner"));
    }

    @PostConstruct
    public void postConstruct() {
        subscribeNotification();
    }

    public void subscribeNotification() {
        StreamObserver<StreamMessage> responseObserver = new StreamObserver<StreamMessage>() {
            @Override
            public void onNext(StreamMessage stream) {
                LOGGER.info("Got a daemon streaming message");
                executorService.submit(() -> {
                    try {
                        renewGoogleCalendarWatch(stream.getMessage());
                    } catch (Exception e) {
                        LOGGER.error("renewGoogleCalendarWatch client side error: ", e);
                    }
                });
                LOGGER.info("Processed a daemon streaming message");
            }
            @Override
            public void onError(Throwable t) {
                Status status = Status.fromThrowable(t);
                LOGGER.error("renewGoogleCalendarWatch server side error: ", status);
            }
            @Override
            public void onCompleted() {
                LOGGER.info("Stopped receiving GoogleCalendarProjectId");
            }
        };
        daemonServiceClient.subscribeNotification(SubscribeNotification.newBuilder().setId("cleaner").build(), responseObserver);
    }

    public void renewGoogleCalendarWatch(String googleCalendarProjectId) throws IOException {
        LOGGER.info("Got googleCalendarProjectId {}", googleCalendarProjectId);
        Calendar service = googleCalClient.getCalendarService();
        Calendar.Events.Watch watch = service.events().watch(googleCalendarProjectId, Util.getChannel());
        LOGGER.info("Created watch {}", watch);
        Channel createdChannel = watch.execute();
        GoogleCalendarProject googleCalendarProject = googleCalendarProjectRepository.findById(googleCalendarProjectId).get();
        googleCalendarProject.setChannelId(createdChannel.getId());
        googleCalendarProject.setExpiration(new Timestamp(createdChannel.getExpiration()));
        googleCalendarProject.setChannel(GSON.toString(createdChannel));
        googleCalendarProjectRepository.save(googleCalendarProject);
        LOGGER.info("Renew GoogleCalendarProject {}, googleCalendarProjectId {}", googleCalendarProject, googleCalendarProjectId);
    }

    @PreDestroy
    public void preDestroy() {
        if (executorService != null) {
            try {
                executorService.awaitTermination(5, TimeUnit.SECONDS);
            } catch (InterruptedException ex) {
                Thread.currentThread().interrupt();
            }
        }
    }
}
