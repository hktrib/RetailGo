import { SignIn } from '@clerk/nextjs';
import styles from './SignInPage.module.css';

const SignInPage = () => {
  return (
    <div className={styles.centerContainer}>
      <SignIn />
    </div>
  );
};

export default SignInPage;
