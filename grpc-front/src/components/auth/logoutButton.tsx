"use client";

import { signOut } from "next-auth/react";

export default function Logout() {
    return (
        <div>
            <button onClick={() => signOut()}>ログアウト</button>
        </div>
    );
}
