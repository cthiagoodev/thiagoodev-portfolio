export function getApiUrl(): string | undefined {
  if (!import.meta.env.SSR) {
    throw new Error("A configuração da API só pode ser lida no servidor.");
  }

  return import.meta.env.API_URL?.trim().replace(/\/+$/, "") || undefined;
}
