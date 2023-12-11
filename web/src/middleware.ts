import { authMiddleware, redirectToSignIn } from "@clerk/nextjs";
import { NextResponse } from "next/server";

// https://clerk.com/docs/references/nextjs/auth-middleware
export default authMiddleware({
  publicRoutes: ["/"],
  afterAuth(auth, req, evt) {
    // redirect them to organization selection page
    if (
      auth.userId &&
      !auth.orgId &&
      req.nextUrl.pathname !== "/register-store" &&
      false
    ) {
      const orgSelection = new URL("/register-store", req.url);
      return NextResponse.redirect(orgSelection);
    }
  },
});

export const config = {
  matcher: ["/((?!.+\\.[\\w]+$|_next).*)", "/", "/(api|trpc)(.*)"],
};
