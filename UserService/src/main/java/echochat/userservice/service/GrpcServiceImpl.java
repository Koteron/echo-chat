package echochat.userservice.service;


import echochat.userservice.grpc.DisplayName;
import echochat.userservice.grpc.GetDisplayNameRequest;
import echochat.userservice.grpc.UserServiceGrpc;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.server.service.GrpcService;

import java.util.UUID;

@RequiredArgsConstructor
@GrpcService
public class GrpcServiceImpl extends UserServiceGrpc.UserServiceImplBase {

    private final UserService userService;

    @Override
    public void getDisplayName(GetDisplayNameRequest request, StreamObserver<DisplayName> responseObserver) {
        DisplayName displayName = DisplayName.newBuilder()
                .setDisplayName(userService.getDisplayName(UUID.fromString(request.getUserId())))
                .build();

        responseObserver.onNext(displayName);
        responseObserver.onCompleted();
    }
}