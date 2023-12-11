import { SignIn } from "@clerk/nextjs";

const SignInPage = () => {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <SignIn
        appearance={{
          elements: {
            formButtonPrimary:
              "bg-sky-500 hover:bg-sky-400 focus:bg-sky-500",
            footerActionLink: "text-sky-600",
          },
        }}
        afterSignUpUrl="/register-store" // Redirect to the store registration page after sign up
      />
    </div>
  );
};

export default SignInPage;
