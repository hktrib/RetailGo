"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { useFetch } from "@/lib/utils";
import { useUser } from "@clerk/nextjs";
import { useRouter } from 'next/router'
import {config} from "@/lib/hooks/config";

// type Member = {
//   firstName: string;
//   lastName: string;
//   email: string;
//   role: string;
// }

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Router } from "next/router";

const formSchema = z.object({
  storeName: z.string().min(0),
  phoneNumber: z.string().min(7, { message: "Invalid phone number" }),
  address1: z.string().min(0),
  address2: z.string().optional(),
  businessType: z.string().min(0),
});

const businessTypes = [
  "Clothing",
  "Grocery",
  "Convenience",
  "Departmnet",
  "Restaurant",
  "Other",
];

export default function RegistrationForm() {
  const router = useRouter()
  const { user } = useUser();
  const authFetch = useFetch();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  // const [members, setMembers] = useState<Member[]>([]);
  // const [showMemberFields, setShowMemberFields] = useState(false);
  // const [memberFirstName, setMemberFirstName] = useState('');
  // const [memberLastName, setMemberLastName] = useState('');
  // const [memberEmail, setMemberEmail] = useState('');
  // const [memberRole, setMemberRole] = useState('');

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

  const handleSubmit = async (values: z.infer<typeof formSchema>) => {
    // not doing anything with `address2` yet
    const { storeName, phoneNumber, address1, address2, businessType } = values;

    // data to be sent in POST request body
    const postData = {
      store_name: storeName,
      store_phone: phoneNumber,
      store_address: address1,
      store_type: businessType,
      owner_email: user?.emailAddresses[0].emailAddress,
    };

    try {
      console.log("POST Data: ", JSON.stringify(postData));
      const response = await authFetch(config.serverURL + "create/store", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          // Any additional headers you need
        },
        body: JSON.stringify(postData),
      });

      console.log("Response:", response);

      if (response.statusCode === 200 || response.statusCode === 201) {
        router.push("/store")
      }

      // Handle the response as needed
    } catch (error) {
      console.error("Error making POST request:", error);
      // Handle errors
    }
  };

  return (
    <div className="relative isolate px-6 pt-14 lg:px-8 flex-1 flex flex-col justify-center items-center">
      <div className="mx-auto max-w-2xl w-full bg-gray-50 py-16 px-12 rounded-xl">
        <h1 className="text-center font-bold text-3xl">
          Register your business
        </h1>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="space-y-8 mt-12"
          >
            <FormField
              control={form.control}
              name="storeName"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Store name</FormLabel>
                  <FormControl>
                    <Input placeholder="Store name" {...field} />
                  </FormControl>
                  {/* <FormDescription>
                This is your public display name.
              </FormDescription> */}
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="phoneNumber"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Phone number</FormLabel>
                  <FormControl>
                    <Input placeholder="Store phone number" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="address1"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Address line 1</FormLabel>
                  <FormControl>
                    <Input placeholder="Address line 1" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="address2"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Address line 2</FormLabel>
                  <FormControl>
                    <Input placeholder="Address line 2" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="businessType"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Business type</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a business type" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {businessTypes.map((businessType) => (
                        <SelectItem key={businessType} value={businessType}>
                          {businessType}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />

            <div className="text-right">
              <Button type="submit">Submit</Button>
            </div>
          </form>
        </Form>
      </div>
    </div>
    //      add members component
    //     {showMemberFields && (
    //       <>
    //         <label className={styles.label}>Member's Name</label>
    //         <input type="text" className={styles.input} placeholder="First Name" value={memberFirstName} onChange={handleMemberFirstName} />
    //         <input type="text" className={styles.input} placeholder="Last Name" value={memberLastName} onChange={handleMemberLastName} />
    //         <input type="email" className={styles.input} placeholder="Email" value={memberEmail} onChange={handleMemberEmail} />
    //         <select className={styles.input} value={memberRole} onChange={handleMemberRole}>
    //           <option value="">Select Role</option>
    //           <option value="owner">Owner</option>
    //           <option value="manager">Manager</option>
    //           <option value="employee">Employee</option>
    //         </select>
    //         <button type="button" onClick={addMember} className={styles.addButton}>Add Member</button>
    //       </>
    //     )}
    //     <button type="button" onClick={() => setShowMemberFields(!showMemberFields)} className={styles.toggleButton}>
    //       {showMemberFields ? 'Close' : 'Add Members'}
    //     </button>

    //     {members.map((member, index) => (
    //       <div key={index} className={styles.memberInfo}>
    //         {member.firstName} {member.lastName} ({member.email}) - {member.role}
    //       </div>
    //     ))}
  );
}
