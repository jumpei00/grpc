"use client";

import Image from "next/image";
import googleSigninBtn from "../../../public/images/google_signin_btn.png";
import {
    signInWithRedirect,
    getRedirectResult,
    GoogleAuthProvider,
} from "firebase/auth";
import { firebaseAuth, googleProvider } from "@/lib/firebase";
import { useEffect } from "react";

export default function Page() {
    const auth = firebaseAuth();
    const provider = googleProvider();

    const googleLoginClick = () => {
        signInWithRedirect(auth, provider).catch((error) => {
            console.log(error);
        });
    };

    useEffect(() => {
        getRedirectResult(auth)
            .then((result) => {
                if (result !== null) {
                    const credential =
                        GoogleAuthProvider.credentialFromResult(result);
                    if (credential !== null) {
                        const token = credential.accessToken;
                        console.log(token);
                        console.log(result.user);
                    }
                }
            })
            .catch((error) => {
                console.log(error);
            });
    }, [auth]);

    return (
        <>
            <h2>Login</h2>
            <div
                onClick={() => {
                    googleLoginClick();
                }}
                style={{
                    cursor: "pointer",
                    width: "200px",
                    height: "50px",
                    position: "relative",
                }}
            >
                <Image
                    src={googleSigninBtn}
                    alt="Google Signin Button"
                    width={200}
                    height={50}
                />
            </div>
        </>
    );
}
