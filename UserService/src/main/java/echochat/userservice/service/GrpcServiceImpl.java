package echochat.userservice.service;


import echochat.userservice.grpc.DisplayNames;
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
    public void getDisplayName(GetDisplayNameRequest request, StreamObserver<DisplayNames> responseObserver) {
        DisplayNames displayName = DisplayNames.newBuilder()
                .addAllDisplayName(userService.getDisplayNames(request.getUserIdList().stream()
                        .map(UUID::fromString).toList()))
                .build();

        responseObserver.onNext(displayName);
        responseObserver.onCompleted();
    }
}