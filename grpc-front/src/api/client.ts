import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { TodoService } from "@/proto/todo/v1/todo_connect";

export const serverTransport = createConnectTransport({
    baseUrl: "http://grpc-back:8080",
});

export const webTransport = createConnectTransport({
    baseUrl: "http://localhost:8080",
});

export const todoServerClient = createPromiseClient(TodoService, serverTransport);

export const todoWebClient = createPromiseClient(TodoService, webTransport);