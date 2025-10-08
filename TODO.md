# webStash TODO

## Features to Implement

### 1. Share Bookmark
- [ ] Add share button to bookmark card
- [ ] Generate shareable link for individual bookmarks
- [ ] Create public bookmark view page (non-authenticated)
- [ ] Copy link to clipboard functionality

### 2. Navigate to Link Icon
- [ ] Add external link icon to bookmark card
- [ ] Icon opens bookmark URL in new tab
- [ ] Visual indicator for clickable link

### 3. Public/Private Bookmarks
- [ ] Add `public` boolean field to bookmarks table
- [ ] Toggle switch in bookmark form (public/private)
- [ ] Filter bookmarks by visibility in UI
- [ ] Public bookmarks accessible via share link
- [ ] Private bookmarks only visible to owner

### 4. Google OAuth
- [ ] Set up Google OAuth provider in Cloudflare
- [ ] Add "Sign in with Google" button
- [ ] Handle OAuth callback
- [ ] Link Google account to existing users
- [ ] Store OAuth tokens in KV
- [ ] Update user profile with Google data

### 5. Stripe Integration
- [ ] Set up Stripe account and get API keys
- [ ] Add Stripe SDK to project
- [ ] Create pricing plans (Free, Pro, Premium)
- [ ] Implement checkout flow
- [ ] Create subscription management page
- [ ] Handle webhook events (subscription created, updated, canceled)
- [ ] Store subscription status in users table
- [ ] Add billing portal link for users to manage subscriptions

### 6. Account Type Model
- [ ] Add `account_type` field to users table (free, pro, premium)
- [ ] Add `subscription_id` field to users table
- [ ] Add `subscription_status` field (active, canceled, past_due)
- [ ] Add `subscription_expires_at` timestamp
- [ ] Create account type badge/indicator in UI
- [ ] Display current plan on user profile/settings page

### 7. Account Type Restrictions
- [ ] Free tier: max 50 bookmarks
- [ ] Pro tier: max 500 bookmarks, advanced search, tags
- [ ] Premium tier: unlimited bookmarks, all features, priority support
- [ ] Implement bookmark limit checks on creation
- [ ] Show upgrade prompts when limits reached
- [ ] Disable features based on account type
- [ ] Add feature comparison table on pricing page
