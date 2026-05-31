/**
 * Site image configuration.
 *
 * Hero / cover images use Lorem Picsum (free, no attribution required).
 * Swap in your own URLs to customize.
 *
 * Decorative SVGs are inline data URIs — zero external requests, never break.
 */

// ── Photographic images (picsum.photos — replace with your own) ──────────
export const HERO_BG =
  'https://picsum.photos/id/102/1200/800'   // warm forest

export const ARTICLE_COVER_FALLBACK =
  'https://picsum.photos/id/180/1200/600'   // nature close-up

// ── Inline SVG decorative images ─────────────────────────────────────────
// Warm abstract blob composition — hero / header decoration
export const DECO_BLOB_WARM =
  "data:image/svg+xml," + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 800 400">
      <defs><filter id="b"><feGaussianBlur stdDeviation="40"/></filter></defs>
      <circle cx="150" cy="120" r="100" fill="#fef3c7" opacity="0.6" filter="url(#b)"/>
      <circle cx="600" cy="80" r="140" fill="#fde68a" opacity="0.35" filter="url(#b)"/>
      <circle cx="400" cy="280" r="120" fill="#fbbf24" opacity="0.2" filter="url(#b)"/>
      <circle cx="680" cy="320" r="90" fill="#d6d3d1" opacity="0.25" filter="url(#b)"/>
    </svg>`
  )

// Subtle grid + dot pattern — global texture overlay
export const TEXTURE_DOTS =
  "data:image/svg+xml," + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 40 40">
      <circle cx="20" cy="20" r="0.6" fill="#a8a29e" opacity="0.3"/>
    </svg>`
  )

// Soft ornamental line — section divider
export const DECO_LINE =
  "data:image/svg+xml," + encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 12" preserveAspectRatio="none">
      <line x1="0" y1="6" x2="200" y2="6" stroke="#d6d3d1" stroke-width="0.5" opacity="0.5"/>
      <circle cx="100" cy="6" r="1.5" fill="#b45309" opacity="0.5"/>
      <circle cx="80" cy="6" r="0.8" fill="#b45309" opacity="0.25"/>
      <circle cx="120" cy="6" r="0.8" fill="#b45309" opacity="0.25"/>
    </svg>`
  )
