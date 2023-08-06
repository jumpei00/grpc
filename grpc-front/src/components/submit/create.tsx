"use client";

import { createTodo } from "@/api/todo";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Create() {
    const router = useRouter();
    const [value, setValue] = useState("");

    const submit = () => {
        createTodo(value)
            .then(() => {
                router.push("/todos");
                setValue("");
            })
            .catch((err) => {
                console.log(err);
            });
    };

    return (
        <div>
            <input
                type="text"
                value={value}
                onChange={(event) => setValue(event.target.value)}
            />
            <button onClick={submit}>Create</button>
        </div>
    );
}
