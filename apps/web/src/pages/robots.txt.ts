import type { APIRoute } from "astro";
import { CONFIG } from "@/config/site";

export const prerender = true;

export const GET: APIRoute = ({ site }) => {
  const sitemapUrl = new URL(
    "sitemap-index.xml",
    site ?? CONFIG.site.url,
  ).href;

  return new Response(`User-agent: *\nAllow: /\n\nSitemap: ${sitemapUrl}\n`, {
    headers: { "Content-Type": "text/plain; charset=utf-8" },
  });
};
