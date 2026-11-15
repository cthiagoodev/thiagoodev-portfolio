import { getApiUrl } from "@/config/env";

export async function fetchApi<T>(path: string): Promise<T | null> {
  const apiUrl = getApiUrl();

  if (!apiUrl) {
    console.warn("API_URL não configurada; dados não serão carregados.");
    return null;
  }

  try {
    const response = await fetch(`${apiUrl}${path}`, {
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      console.error(
        `Erro ao carregar ${path} da API: status ${response.status}.`,
      );
      return null;
    }

    return (await response.json()) as T;
  } catch {
    console.error(`Não foi possível conectar à API para carregar ${path}.`);
    return null;
  }
}
