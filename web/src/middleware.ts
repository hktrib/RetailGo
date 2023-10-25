import { authMiddleware } from "@clerk/nextjs";

// https://clerk.com/docs/references/nextjs/auth-middleware
export default authMiddleware({
  publicRoutes: ["/"],
});

export const config = {
  matcher: ["/((?!.+\\.[\\w]+$|_next).*)", "/", "/(api|trpc)(.*)"],
};
