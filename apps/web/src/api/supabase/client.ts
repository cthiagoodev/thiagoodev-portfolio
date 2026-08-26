import { createClient } from "@supabase/supabase-js";
import { getSupabaseEnv } from "@/config/env";

export function createSupabaseClient() {
  const { url: supabaseUrl, publishableKey: supabasePublishableKey } =
    getSupabaseEnv();

  if (!supabaseUrl || !supabasePublishableKey) {
    return null;
  }

  return createClient(supabaseUrl, supabasePublishableKey, {
    auth: {
      autoRefreshToken: false,
      persistSession: false,
      detectSessionInUrl: false,
    },
  });
}
