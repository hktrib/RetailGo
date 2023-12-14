import { SignUp } from "@clerk/nextjs";

const SignUpPage = () => {
  return (
    <div className="flex min-h-screen items-center justify-center bg-white">
      <SignUp
        appearance={{
          elements: {
            formButtonPrimary: "bg-sky-500 hover:bg-sky-400 focus:bg-sky-500",
            footerActionLink: "text-sky-600",
          },
        }}
        afterSignUpUrl="/register-store" // Redirect to the store registration page after sign up
      />
    </div>
  );
};

export default SignUpPage;
