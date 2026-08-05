<template>
  <div class="tw-mx-auto tw-max-w-3xl tw-px-4 tw-py-10 tw-text-very-dark-gray">
    <h1 class="tw-mb-1 tw-text-3xl tw-font-bold tw-text-black">
      Privacy Policy
    </h1>
    <div class="tw-mb-8 tw-text-sm tw-text-dark-gray">
      Last updated: 25 July 2026
    </div>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        1. Who is responsible for your data
      </h2>
      <p>
        The data controller is the
        <strong>Artificial Intelligence Student Collective Madrid (AISC)</strong
        >, a university association of Universidad Carlos III de Madrid, with
        registered office at Avenida de la Universidad, No. 30, Campus of the
        Escuela Politécnica Superior, office 2.3.C05, Leganés, 28911 (Madrid,
        Spain).
      </p>
      <p>
        Contact for any privacy matter:
        <a href="mailto:aisc.asoc@uc3m.es">aisc.asoc@uc3m.es</a>
      </p>
      <p>
        This application, Meet AISC, is a deployment, operated by AISC, of the
        open-source scheduling tool Timeful (formerly Schej.it). AISC is the
        controller of
        the personal data processed through this deployment; the original
        authors of the software do not receive your data.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        2. Signing in requires a Google account
      </h2>
      <p>
        Creating an account and using the app requires signing in with Google.
        We offer no other sign-in method, and we never receive or store your
        Google password: authentication happens entirely on Google's systems
        through OAuth 2.0.
      </p>
      <p>
        When you sign in we request the <code>openid</code>,
        <code>email</code> and <code>profile</code> scopes, and Google returns
        to us the data described in the next section.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        3. What data we collect and store
      </h2>

      <h3 class="tw-mb-1 tw-mt-4 tw-font-semibold">
        a) Google account information
      </h3>
      <ul>
        <li>First and last name</li>
        <li>Email address</li>
        <li>Profile picture (the URL provided by Google)</li>
      </ul>
      <p>
        This information is stored in our database to create and identify your
        account and to show your name and picture to the other participants of
        the events you take part in.
      </p>

      <h3 class="tw-mb-1 tw-mt-4 tw-font-semibold">b) Google OAuth tokens</h3>
      <p>
        We store the access token and refresh token issued by Google, together
        with their expiry date and the scopes you granted, so that you stay
        signed in and so that we can read your calendar if — and only if — you
        granted that permission. These tokens are stored encrypted (AES), are
        never exposed through our API and are never shown to other users.
      </p>

      <h3 class="tw-mb-1 tw-mt-4 tw-font-semibold">
        c) Google Calendar data (optional)
      </h3>
      <p>
        Connecting your calendar is optional — the app works without it and you
        can always fill in your availability by hand. If you do connect it, we
        request
        <code>calendar.calendarlist.readonly</code> (to list your calendars)
        and <code>calendar.events</code> (to read your events for
        availability pre-fill, and — only when you or someone you've
        delegated to explicitly clicks "Schedule" on an event you own — to
        create the resulting calendar event and Google Meet link and invite
        respondents to it).
      </p>
      <ul>
        <li>
          We never create, modify, or delete anything in your calendar except
          the single event created by an explicit "Schedule" action you or
          your delegate take.
        </li>
        <li>
          Calendar events are otherwise only fetched from Google while you use
          the app and shown only to you. Other participants see availability
          (free/busy time), never your event titles or details.
        </li>
        <li>
          We store the list of calendars you enable and your autofill
          preferences (working hours, buffer time) so that the app behaves as
          you configured it.
        </li>
        <li>
          You can disconnect a calendar account at any time from
          <em>Settings</em>, and revoke our access at any time from
          <a
            href="https://myaccount.google.com/permissions"
            target="_blank"
            rel="noopener noreferrer"
            >your Google account permissions page</a
          >.
        </li>
      </ul>

      <h3 class="tw-mb-1 tw-mt-4 tw-font-semibold">
        d) Data you create in the app
      </h3>
      <ul>
        <li>
          Events and groups you create: name, description, dates and times,
          settings and the people invited.
        </li>
        <li>
          Your responses: the time slots you mark as available or "if needed",
          and any sign-up form answers.
        </li>
        <li>
          If you respond to an event without signing in, the name — and the
          email address, if the organiser asks for it — that you type in.
        </li>
      </ul>

      <h3 class="tw-mb-1 tw-mt-4 tw-font-semibold">e) Technical data</h3>
      <ul>
        <li>
          A session cookie (named <code>session</code>), strictly necessary to
          keep you signed in.
        </li>
        <li>
          Server logs and basic usage records (for example, the dates on which
          an account was active), used to keep the service running and secure
          and to measure aggregate usage.
        </li>
      </ul>
      <p>
        We do not use advertising cookies, we do not build profiles about you,
        and we do not take automated decisions producing legal effects on you.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        4. Google API Services — Limited Use
      </h2>
      <p>
        Our use and transfer of information received from Google APIs adheres to
        the
        <a
          href="https://developers.google.com/terms/api-services-user-data-policy"
          target="_blank"
          rel="noopener noreferrer"
          >Google API Services User Data Policy</a
        >, including the Limited Use requirements. Data obtained from your
        Google profile and Google Calendar is used only to provide the
        scheduling features described above. It is never sold, never used for
        advertising, never transferred to others except as needed to provide the
        service or where the law requires it, and never read by a human except
        with your explicit consent, for security purposes, to comply with
        applicable law, or in aggregated and anonymised form.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        5. Why we process your data, and on what legal basis
      </h2>
      <ul>
        <li>
          <strong>To provide the service</strong> (account creation, creating
          and answering availability polls): performance of a contract with you,
          Art. 6.1.b GDPR.
        </li>
        <li>
          <strong>To read your calendar and pre-fill your availability</strong>:
          your explicit consent, Art. 6.1.a GDPR, given when you grant the
          Google permission. You may withdraw it at any time.
        </li>
        <li>
          <strong>To email you about events you take part in</strong> (reminders
          and notifications, where the organiser enables them): performance of
          the contract and our legitimate interest in the service working as
          expected, Art. 6.1.b and 6.1.f GDPR.
        </li>
        <li>
          <strong>To keep the service secure and functioning</strong> (logs,
          aggregate usage statistics): our legitimate interest, Art. 6.1.f GDPR.
        </li>
      </ul>
      <p>
        We do not use the data collected here for association newsletters or
        marketing. If you receive AISC communications, that is because you
        subscribed to them separately.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        6. How long we keep it
      </h2>
      <ul>
        <li>
          Account data and the content you create: for as long as your account
          exists.
        </li>
        <li>
          OAuth tokens: until they expire, or until you disconnect the account
          or revoke access in Google.
        </li>
        <li>
          Calendar events: not kept as a permanent record — they are retrieved
          from Google while you use the app.
        </li>
        <li>
          Technical logs: for a limited period, only as long as needed for
          security and diagnostics.
        </li>
      </ul>
      <p>
        When you ask us to delete your account we erase your personal data and
        your responses, except for anything we are legally required to keep.
        Events you created may remain visible to their participants unless you
        delete them first; tell us if you want them removed as well.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        7. Who else sees your data
      </h2>
      <ul>
        <li>
          <strong>Other participants of your events</strong>: they see your
          name, profile picture and the availability you submit — and your email
          address if the organiser enabled email collection.
        </li>
        <li>
          <strong>Google</strong>, as the provider of the sign-in and calendar
          APIs you choose to use.
        </li>
        <li>
          <strong>Our hosting and infrastructure providers</strong>, acting as
          data processors under a contract meeting the requirements of Art. 28
          GDPR, for the sole purpose of running the servers and database of this
          app.
        </li>
      </ul>
      <p>
        We do not sell your personal data and we do not share it with anyone
        else, except where required by law or by a competent authority.
      </p>
      <p>
        Where a provider processes data outside the European Economic Area, the
        transfer is covered by an adequacy decision or by the European
        Commission's Standard Contractual Clauses.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        8. Security
      </h2>
      <p>
        We apply appropriate technical and organisational measures to protect
        your data: traffic is served over HTTPS, OAuth tokens are encrypted at
        rest, session cookies are signed, and access to the database is
        restricted to the people who need it to operate the service. No system
        is completely secure, but we will notify you and the competent authority
        of any personal data breach as required by Art. 33 and 34 GDPR.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        9. Your rights
      </h2>
      <p>
        Under the GDPR (Regulation (EU) 2016/679) and Spanish Organic Law
        3/2018, you have the right to access your data, rectify it, erase it,
        restrict or object to its processing, receive it in a portable format,
        and withdraw any consent you have given (without affecting processing
        carried out before the withdrawal).
      </p>
      <p>
        To exercise these rights, write to
        <a href="mailto:aisc.asoc@uc3m.es">aisc.asoc@uc3m.es</a> from the email
        address associated with your account. We may ask for reasonable proof of
        identity if we cannot otherwise verify who you are, and we will reply
        within one month.
      </p>
      <p>
        You can also revoke this app's access to your Google account directly at
        <a
          href="https://myaccount.google.com/permissions"
          target="_blank"
          rel="noopener noreferrer"
          >myaccount.google.com/permissions</a
        >.
      </p>
      <p>
        If you believe your data has not been handled correctly, you may lodge a
        complaint with the Spanish Data Protection Agency (Agencia Española de
        Protección de Datos,
        <a href="https://www.aepd.es" target="_blank" rel="noopener noreferrer"
          >www.aepd.es</a
        >).
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        10. Minors
      </h2>
      <p>
        This service is intended for university students and other adults. It is
        not directed at children under 14, and we do not knowingly collect their
        data. If you believe a minor has registered, contact us and we will
        delete the account.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        11. Links to third-party websites
      </h2>
      <p>
        The app may contain links to third-party websites. AISC is not
        responsible for the privacy or cookie policies of those websites; your
        use of them is governed by their own terms.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        12. Changes to this policy
      </h2>
      <p>
        We may update this policy to reflect changes in the app or in the law.
        The version published on this page, with its "last updated" date, is
        always the one in force. If a change is significant, we will make it
        visible in the app.
      </p>
    </section>

    <section class="tw-mb-8">
      <h2 class="tw-mb-2 tw-text-xl tw-font-semibold tw-text-black">
        13. Acceptance
      </h2>
      <p>
        By signing in with Google and using this app, you confirm that you have
        read and understood this Privacy Policy.
      </p>
    </section>
  </div>
</template>

<script>
export default {
  name: "PrivacyPolicy",

  metaInfo: {
    title: "Privacy Policy - Schedule meetings",
  },
}
</script>

<style scoped>
p {
  @apply tw-mb-3;
}
ul {
  @apply tw-mb-3 tw-list-disc tw-pl-6;
}
li {
  @apply tw-mb-1;
}
a {
  @apply tw-text-darkest-green tw-underline;
}
code {
  @apply tw-text-xs tw-text-black;
  background-color: transparent;
}
</style>
