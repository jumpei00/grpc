"use client";

import { deleteTodo } from "@/api/todo";
import { useRouter, useParams } from "next/navigation";

export default function Delete() {
    const params = useParams();
    const router = useRouter();

    const submit = () => {
        deleteTodo(params.id)
            .then(() => {
                router.push("/todos");
            })
            .catch((err) => {
                console.log(err);
            });
    };

    return <button onClick={submit} style={{marginTop: "16px"}}>Delete</button>;
}
