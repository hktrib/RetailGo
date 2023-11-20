// pages/api/send-invite.js

import nodemailer from 'nodemailer';

export default async function handler(req, res) {
  if (req.method === 'POST') {
    try {
      const { email } = req.body;

      // Set up Nodemailer transporter
      let transporter = nodemailer.createTransport({
        // Use a service like Gmail or any other SMTP server
        service: 'gmail',
        auth: {
          user: process.env.EMAIL_USER, // Your email address
          pass: process.env.EMAIL_PASSWORD, // Your email password
        },
      });

      // Set up email options
      const mailOptions = {
        from: process.env.EMAIL_USER,
        to: email,
        subject: 'Invitation to Join as an Employee',
        text: 'You have been invited to join as an employee. Please follow the link to sign up.', // Add the link to your signup page here
      };

      // Send email
      await transporter.sendMail(mailOptions);

      res.status(200).json({ message: 'Invitation sent!' });
    } catch (error) {
      res.status(500).json({ message: 'Error sending invitation' });
    }
  } else {
    res.setHeader('Allow', ['POST']);
    res.status(405).end(`Method ${req.method} Not Allowed`);
  }
}
