import { getTodos } from "@/api/todo";
import Create from "@/components/submit/create";
import Link from "next/link";
import React from "react";

export default async function Page() {
    const data = await getTodos();

    return (
        <>
            <h2>TODO</h2>
            <ul>
                {data.map((todo) => (
                    <li key={todo.key}>
                        <div style={{ display: "flex" }}>
                            <Link href={`todos/${todo.key}`}>
                                {todo.content}
                            </Link>
                        </div>
                    </li>
                ))}
            </ul>
            <Create />
        </>
    );
}
