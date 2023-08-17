import NextAuth from "next-auth/next";
import { NextAuthOptions } from "next-auth";
import GoogleProvier from "next-auth/providers/google";

export const authOptions: NextAuthOptions = {
    providers: [
        GoogleProvier({
            clientId: process.env.GOOGLE_CLIENT_ID || "",
            clientSecret: process.env.GOOGLE_CLIENT_SECRET || "",
        }),
    ],
    debug: true,
    session: {
        strategy: "jwt",
    },
};

const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };
