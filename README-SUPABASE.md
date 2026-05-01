# Supabase

## Google

**1. Google Branding:**

    1. Navigate to Google Auth Platform > Branding in the GCP Console.
    2. Select User Type (Internal or External).
    3. Fill in the App Information (App name, User support email, App logo).
    4. Fill in the App domain links (Homepage, Privacy Policy, Terms of Service).
    5. Add Authorized domains (e.g. [domain], supabase.co).
    6. Provide developer contact information and save.

**2. Google Client:**

    1. Go to Google Auth Platform > Clients.
    2. Set the Application type to Web application.
    3. Name the client.
    4. Under Authorized JavaScript origins, add:
        https://[domain]
        https://stag.[domain]
        http://localhost:8081

    5. Under Authorized redirect URIs, add the callback URL provided by Supabase (Authentication > Providers > Google):
        https://auth.[domain]/auth/v1/callback

    6. Click Create to get Client ID and Client Secret.

## Apple

**1. Apple App ID:**

    1. Go to Apple Developer > Certificates, Identifiers & Profiles.
    2. Open or create the App ID for the iOS bundle identifier:
        [bundle_id]

    3. Enable Sign in with Apple for the App ID.
    4. Save the App ID configuration.

**2. Apple Services ID:**

    1. Go to Identifiers and create a new Services ID.
    2. Use a clear identifier, for example:
        [bundle_id].auth

    3. Enable Sign in with Apple for the Services ID.
    4. Configure the Services ID web authentication settings.
    5. Under Domains and Subdomains, add:
        auth.[domain]

    6. Under Return URLs, add the Supabase callback URL:
        https://auth.[domain]/auth/v1/callback

    7. If Supabase shows a different Apple callback URL in Authentication > Providers > Apple, use the exact URL shown by Supabase.

**3. Apple Private Key:**

    1. Go to Keys and create a new key.
    2. Enable Sign in with Apple for the key.
    3. Download the .p8 private key and store it securely. Apple only allows this download once.
    4. Record the Team ID, Services ID, Key ID, and private key content.
    5. Use the Supabase Apple auth docs generator to create the OAuth client_secret JWT from:
        Team ID
        Services ID
        Key ID
        .p8 - private key content

    6. Store the .p8 file securely. The generated OAuth client_secret expires every 6 months and must be regenerated from this file.

## Supabase

**1. Supabase Google Provider:**

    1. Go to your Supabase Dashboard.
    2. Navigate to Authentication > Providers.
    3. Select Google and toggle "Enable Google".
    4. Enter the Client ID and Client Secret obtained from GCP.
    5. Save the configuration.

**2. Supabase Apple Provider:**

    1. Go to your Supabase Dashboard.
    2. Navigate to Authentication > Providers.
    3. Select Apple and toggle "Enable Apple".
    4. Under Client IDs, enter the Apple Services ID:
        [bundle_id].auth

    5. Under Secret Key (for OAuth), enter the generated Apple OAuth client_secret JWT.
    6. Leave "Allow users without an email" disabled. The API requires an email claim from Supabase.
    7. Confirm the Callback URL (for OAuth) shown by Supabase is registered in Apple Developer:
        https://auth.[domain]/auth/v1/callback

    8. Save the configuration.

**3. Supabase URL Configuration:**

    Site URL:
    https://[domain]

    Redirect URLs:
    https://stag.[domain]/auth/callback
    http://localhost:8081/auth/callback
    https://[domain]/auth/callback
    [app]://auth/callback

**4. Supabase Environment Variables:**

   Project Settings > API Keys:

    SUPABASE_SECRET=[Publishable_Key]
    SUPABASE_ISSUER=[Project_URL]/auth/v1
