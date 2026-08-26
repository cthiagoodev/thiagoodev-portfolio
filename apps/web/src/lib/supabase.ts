import { createClient } from "@supabase/supabase-js";

export function createSupabaseClient() {
  if (!import.meta.env.SSR) {
    throw new Error("O client do Supabase só pode ser criado no servidor.");
  }

  const supabaseUrl = import.meta.env.SUPABASE_URL;
  const supabasePublishableKey = import.meta.env.SUPABASE_PUBLISHABLE_KEY;

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
