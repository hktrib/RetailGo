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
      />
    </div>
  );
};

export default SignUpPage;
