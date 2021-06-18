
package io.nats.spring;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import org.springframework.cloud.stream.annotation.EnableBinding;
import org.springframework.cloud.stream.annotation.StreamListener;
import org.springframework.cloud.stream.messaging.Sink;

@EnableBinding(Sink.class)
public class Listener {
    private static final Log logger = LogFactory.getLog(Listener.class);

    @StreamListener(Sink.INPUT)
    public void handle(Object message) {
        logger.info("received message " + message);
    }
}