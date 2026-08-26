export function getSupabaseEnv() {
  if (!import.meta.env.SSR) {
    throw new Error("A configuração do Supabase só pode ser lida no servidor.");
  }

  return {
    url: import.meta.env.SUPABASE_URL,
    publishableKey: import.meta.env.SUPABASE_PUBLISHABLE_KEY,
  };
}
