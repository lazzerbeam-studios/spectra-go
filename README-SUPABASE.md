# Supabase

## Google

**1. Google Branding:**

    1. Navigate to Google Auth Platform > Branding in the GCP Console.
    2. Select User Type (Internal or External).
    3. Fill in the App Information (App name, User support email, App logo).
    4. Fill in the App domain links (Homepage, Privacy Policy, Terms of Service).
    5. Add Authorized domains (e.g. [domain.com], supabase.co).
    6. Provide developer contact information and save.

**2. Google Client:**

    1. Go to Google Auth Platform > Clients.
    2. Set the Application type to Web application.
    3. Name the client.
    4. Under Authorized JavaScript origins, add:
        https://[domain.com]
        https://[staging.domain.com]
        http://localhost:[port]

    5. Under Authorized redirect URIs, add the callback URL provided by Supabase (Authentication > Providers > Google):
        https://[auth.domain.com]/auth/v1/callback

    6. Click Create to get Client ID and Client Secret.

## Supabase

**1. Supabase Providers:**

    1. Go to your Supabase Dashboard.
    2. Navigate to Authentication > Providers.
    3. Select Google and toggle "Enable Google".
    4. Enter the Client ID and Client Secret obtained from GCP.
    5. Save the configuration.

**2. Supabase URL Configuration:**

Site URL:

    https://[domain.com]

Redirect URLs:

    https://[staging.domain.com]/auth/callback
    http://localhost:[port]/auth/callback
    https://[domain.com]/auth/callback
    [app-scheme]://auth/callback

**3. Supabase Environment Variables:**

   Project Settings > API Keys:

    SUPABASE_SECRET=[Publishable_Key]
    SUPABASE_ISSUER=[Project_URL]/auth/v1
