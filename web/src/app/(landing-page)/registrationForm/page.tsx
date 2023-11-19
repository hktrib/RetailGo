'use client';

//Comments added for add member functionality

import { useState } from 'react';
import Head from 'next/head';
import styles from './Registration.module.css';
import { useFetch } from '@/lib/utils';
import { useUser } from "@clerk/nextjs";

// type Member = {
//   firstName: string;
//   lastName: string;
//   email: string;
//   role: string;
// }

export default function RegistrationForm() {


  const {user} = useUser()
  const authFetch = useFetch()


  const [phoneNumber, setPhoneNumber] = useState('');
  const [storeName, setStoreName] = useState('');
  const [address, setAddress] = useState('');
  const [address2, setAddress2] = useState('');
  const [businessType, setBusinessType] = useState('');
  // const [members, setMembers] = useState<Member[]>([]);
  // const [showMemberFields, setShowMemberFields] = useState(false);
  // const [memberFirstName, setMemberFirstName] = useState('');
  // const [memberLastName, setMemberLastName] = useState('');
  // const [memberEmail, setMemberEmail] = useState('');
  // const [memberRole, setMemberRole] = useState('');
  const [submitted, setSubmitted] = useState(false);
  const [error, setError] = useState(false);

  const handleStorePhoneNumber = (e: React.ChangeEvent<HTMLInputElement>) => setPhoneNumber(e.target.value);
  const handleStoreName = (e: React.ChangeEvent<HTMLInputElement>) => setStoreName(e.target.value);
  const handleAddress = (e: React.ChangeEvent<HTMLInputElement>) => setAddress(e.target.value);
  const handleAddress2 = (e: React.ChangeEvent<HTMLInputElement>) => setAddress2(e.target.value);
  const handleBusinessType = (e: React.ChangeEvent<HTMLSelectElement>) => setBusinessType(e.target.value);
  // const handleAddMembersClick = () => setShowMemberFields(true);
  // const handleMemberFirstName = e => setMemberFirstName(e.target.value);
  // const handleMemberLastName = e => setMemberLastName(e.target.value);
  // const handleMemberEmail = e => setMemberEmail(e.target.value);
  // const handleMemberRole = e => setMemberRole(e.target.value);

  // const addMember = () => {
  //   setMembers(prevMembers => [
  //     ...prevMembers,
  //     { firstName: memberFirstName, lastName: memberLastName, email: memberEmail, role: memberRole }
  //   ]);
  //   setMemberFirstName('');
  //   setMemberLastName('');
  //   setMemberEmail('');
  //   setMemberRole('');
  // };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (phoneNumber === '' || storeName === '' || address === '' || businessType === '') {
      setError(true);
    } else {

      // const email = await {}

      const postData = {
        // Your data to be sent in the POST request body
        store_name: storeName,
        store_phone: phoneNumber,
        store_address: address,
        store_type: businessType,
        owner_email: user?.emailAddresses[0]
      };

      setSubmitted(true);
      try {
        const response = await authFetch("http://localhost:8080/create/store", {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            // Any additional headers you need
          },
          body: JSON.stringify(postData),
        });
    
        console.log('Response:', response);
        // Handle the response as needed
      } catch (error) {
        console.error('Error making POST request:', error);
        // Handle errors
      }
      setError(false);
    }
  };

  const successMessage = () => (
    <div className={styles.success} style={{ display: submitted ? 'block' : 'none' }}>
      <h1>{storeName} is successfully registered!</h1>
    </div>
  );

  const errorMessage = () => (
    <div className={styles.error} style={{ display: error ? 'block' : 'none' }}>
      <h1>Please enter all the fields.</h1>
    </div>
  );

  return (
    <div className={styles.formContainer}>
      <Head>
        <title>Register Your Business</title>
      </Head>
      <div className={styles.formTitle}>
        <h1>Register Your Business!</h1>
      </div>

      <div className={styles.messages}>
        {errorMessage()}
        {successMessage()}
      </div>

      <form onSubmit={handleSubmit}>

        <label className={styles.label}>Store Name<span className={styles.required}>*</span></label>
        <input onChange={handleStoreName} className={styles.input} value={storeName} type="text" />
        
        <label className={styles.label}>Store Phone Number<span className={styles.required}>*</span></label>
        <input onChange={handleStorePhoneNumber} className={styles.input} value={phoneNumber} type="text" placeholder="First Name" required />

        <label className={styles.label}>Address<span className={styles.required}>*</span></label>
        <input onChange={handleAddress} className={styles.input} value={address} type="text" placeholder = "Line 1" />
        <input onChange={handleAddress2} className={styles.input} value={address2} type="text" placeholder = "Line 2" />

        <label className={styles.label}>Business Type<span className={styles.required}>*</span></label>
        <select onChange={handleBusinessType} className={styles.input} value={businessType}>
          <option value="">Select Business Type</option>
          <option value="clothing">Clothing</option>
          <option value="grocery">Grocery</option>
          <option value="convenience">Convenience</option>
          <option value="department">Department</option>
          <option value="restaurant">Restaurant</option>
          <option value="other">Other</option>
        </select>
        
        { /* add members component  */}
        {/* {showMemberFields && (
          <>
            <label className={styles.label}>Member's Name</label>
            <input type="text" className={styles.input} placeholder="First Name" value={memberFirstName} onChange={handleMemberFirstName} />
            <input type="text" className={styles.input} placeholder="Last Name" value={memberLastName} onChange={handleMemberLastName} />
            <input type="email" className={styles.input} placeholder="Email" value={memberEmail} onChange={handleMemberEmail} />
            <select className={styles.input} value={memberRole} onChange={handleMemberRole}>
              <option value="">Select Role</option>
              <option value="owner">Owner</option>
              <option value="manager">Manager</option>
              <option value="employee">Employee</option>
            </select>
            <button type="button" onClick={addMember} className={styles.addButton}>Add Member</button>
          </>
        )}
        <button type="button" onClick={() => setShowMemberFields(!showMemberFields)} className={styles.toggleButton}>
          {showMemberFields ? 'Close' : 'Add Members'}
        </button>

        {members.map((member, index) => (
          <div key={index} className={styles.memberInfo}>
            {member.firstName} {member.lastName} ({member.email}) - {member.role}
          </div>
        ))} */}

        <button className={styles.btn} type="submit">Submit</button>
      </form>
    </div>
  );
}
