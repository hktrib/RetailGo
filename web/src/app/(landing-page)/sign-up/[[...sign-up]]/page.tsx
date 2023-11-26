import { SignUp } from "@clerk/nextjs";

const SignUpPage = () => {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <SignUp
        appearance={{
          elements: {
            formButtonPrimary:
              "bg-amber-500 hover:bg-amber-400 focus:bg-amber-500",
            footerActionLink: "text-amber-600",
          },
        }}
        afterSignUpUrl="/registrationForm" // Redirect to the store registration page after sign up
      />
    </div>
  );
};

export default SignUpPage;
