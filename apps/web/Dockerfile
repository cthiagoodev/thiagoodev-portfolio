FROM node:22-slim AS builder
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM node:22-slim AS runner
WORKDIR /app

ENV NODE_ENV=production

COPY package.json package-lock.json ./
RUN npm ci --omit=dev

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/wrangler.jsonc ./wrangler.jsonc
COPY --from=builder /app/astro.config.mjs ./astro.config.mjs
COPY --from=builder /app/src ./src

EXPOSE 4321

CMD ["npm", "run", "preview", "--", "--host", "0.0.0.0", "--port", "4321"]

