"use client";

import { updateTodo } from "@/api/todo";
import { useParams, useRouter } from "next/navigation";
import { useState } from "react";

interface Props {
    content: string;
}

export default function Update(props: Props) {
    const params = useParams();
    const router = useRouter();
    const [value, setValue] = useState(props.content);

    const submit = () => {
        updateTodo(params.id, value)
            .then(() => {
                router.push("/todos");
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
            <button onClick={submit}>Update</button>
        </div>
    );
}
