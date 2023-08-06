import { todoServerClient, todoWebClient } from "@/client/client";
import { Todo } from "@/proto/todo/v1/todo_pb";

export async function getTodos(): Promise<Todo[]> {
    const res = await todoServerClient.index({});
    return res.todos;
}

export async function getTodo(key: string): Promise<Todo> {
    const res = await todoServerClient.show({ key: key });
    if (res.todo) {
        return res.todo;
    }
    throw new Error("Todo not found");
}

export async function createTodo(content: string): Promise<Todo> {
    const newTodo = new Todo();
    newTodo.content = content;
    const res = await todoWebClient.create({ todo: newTodo });
    if (res.todo) {
        return res.todo;
    }
    throw new Error("Todo not found");
}

export async function updateTodo(key: string, content: string): Promise<Todo> {
    const todo = new Todo();
    todo.key = key;
    todo.content = content;
    const res = await todoWebClient.update({ todo: todo });
    if (res.todo) {
        return res.todo;
    }
    throw new Error("Todo not found");
}

export async function deleteTodo(key: string): Promise<void> {
    await todoWebClient.delete({ key: key });
}